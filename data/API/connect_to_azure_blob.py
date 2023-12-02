from datetime import datetime,timedelta , timezone
from azure.storage.blob import BlobServiceClient , generate_blob_sas , BlobSasPermissions
import pandas as pd

def get_data_from_containers():
    #Blob container credentials
    #Instead of one credential per container , a list of credentials is defined for each model in the following order:
    # Spam detection , Post classification , Sentiment Analysis
    accounts_names = ['spmodelsk','mlpost2286547183','mlsa5209976250']
    account_keys = ['w776BBjG7QgFTTcBwsQBxeoSY4ufYAY4G44Ozevr1SXItrU1JSrwDdH80Cfb3mCo2Lrtn8BnQEvs+AStnmwuNA==','2YIa6CgYK6+XQspj8Dzr1vL654+yVWM55GZZ95VqFyreqwBCaoUEgaNlBMQoeCenng00S1R0VVX5+AStCftE5A==','Ci4rNxrBWP5lm8qiiwkDuRNwFxVfADRvCW7sA3dpg/8K1C10miHpMVdqqK5XMxrTfGSjIVUNAegg+AStgJBh2w==']
    container_names = ['spmodelcontenedor','azureml','predictions']

    #Create clients to interact with blob storage
    conect_str_sp = 'DefaultEndpointsProtocol=https;AccountName=' + accounts_names[0] + ';AccountKey=' + account_keys[0] + ';EndpointSuffix=core.windows.net' #For spam detection
    conect_str_po = 'DefaultEndpointsProtocol=https;AccountName=' + accounts_names[1] + ';AccountKey=' + account_keys[1] + ';EndpointSuffix=core.windows.net' #For post classification
    conect_str_sa = 'DefaultEndpointsProtocol=https;AccountName=' + accounts_names[2] + ';AccountKey=' + account_keys[2] + ';EndpointSuffix=core.windows.net' #For Sentiment analysis

    #A list of all connection strings
    connect_str_list = [conect_str_sp,conect_str_po,conect_str_sa]

    blob_service_client_list = [] #List of all clients in Blob storage service
    for connect_str in connect_str_list:
        blob_service_client_list.append(BlobServiceClient.from_connection_string(connect_str)) #All connections are made

    #All retrieved containers
    container_client_list = []
    #Use the client to connect to containers
    for container , blob_service_client in zip(container_names,blob_service_client_list):
        container_client_list.append(blob_service_client.get_container_client(container))

    #Get a list of all blob files in each container
    blob_list = []
    for container_client in container_client_list:
        for blob_i in container_client.list_blobs():
            blob_list.append(blob_i.name) #We retrieve all files in container

    #Get rid of useless files that are not data results
    blob_list.remove('logistic_regression_model.pkl')
    blob_list.remove('tfidf_vectorizer.pkl')

    #Define the dataframe list to return
    df_list = []
    for blob_i , container , account , account_key in zip(blob_list,container_names,accounts_names,account_keys):
        #generate a shared access signature for each blob file (SAS)
        sas_i = generate_blob_sas(account_name = account,
                                    container_name = container,
                                    blob_name = blob_i,
                                    account_key=account_key,
                                    permission=BlobSasPermissions(read=True),
                                    expiry=datetime.utcnow() + timedelta(hours=22))

        sas_url = 'https://' + account + '.blob.core.windows.net/' + container + '/' + blob_i + '?' + sas_i #SaS is generated , it grants permissions of file reading
        
        # get every client in container
        blob_client = container_client_list[container_names.index(container)].get_blob_client(blob_i)
        last_modified = blob_client.get_blob_properties().last_modified #Get the property of LAST UPLOAD of each file
        
        # UTC Format apply
        last_modified_utc = last_modified.replace(tzinfo=timezone.utc)

        # Format to SQL SERVER database format
        formatted_last_modified = last_modified_utc.strftime("%Y-%m-%d %H:%M:%S.%f")[:-3]

        # Read csv with latin-1 decoding
        df = pd.read_csv(sas_url, encoding='latin-1')
        df['timestamp'] = formatted_last_modified

        df_list.append(df)

    return df_list
    
