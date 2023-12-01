from azure.storage.blob import BlobServiceClient

storage_account_key = "your account_key"
storage_account_name = "your account_name"
connection_string = "your connection_string"
container_name = "your container_name"

# Function that receives the .csv file and the name of the file to be saved in Azure Storage.
def uploadStorage(file_content,file_name):
    blob_service_client = BlobServiceClient.from_connection_string(connection_string)
    blob_client = blob_service_client.get_blob_client(container=container_name, blob=file_name)
    # Upload the file to Azure Storage
    blob_client.upload_blob(file_content.getvalue().encode('utf-8'), overwrite=True)