import pandas as pd
import conexion_sql
from connect_to_azure_blob import get_data_from_containers

my_engine = conexion_sql.connect_with_alchemy()

def load_data_to_db(dataframe_list,table_names , batch_size=1000):

    #Iterate for each dataframe and table_name
    for df ,table_name in zip(dataframe_list,table_names):
        try:
            #Divide dataframe in batches of 1000 registers each
            total_rows = len(df)
            num_batches = total_rows // batch_size + 1
            num_batches = min(num_batches, 10)

            for i in range(num_batches):
                start_idx = i * batch_size
                end_idx = (i + 1) * batch_size
                batch = df.iloc[start_idx:end_idx]

                #For each bacth the upload is made
                if not batch.empty:
                    batch.to_sql(table_name, con=my_engine, index=False, if_exists='append') #Upload to Azure DB
                    print(f"Batch {i + 1} of {num_batches} succesfully uploaded.")
            print("Successfull data uploading")
        except Exception as e:
            print(f"Error in data uploading: {e}")
    
def format_dataframe(dataframe_list):
    #For spam detecion
    spam_columns = ['v1','v2','timestamp']
    dataframe_list[0] = dataframe_list[0][spam_columns]
    dataframe_list[0].columns = ['prediction','text','timestamp']
    dataframe_list[0]['prediction'] = dataframe_list[0]['prediction'].map({'spam': 1, 'ham': 0})

    #For spam detection
    post_columns = ['clean_text','score_prediction','timestamp']
    dataframe_list[1] = dataframe_list[1][post_columns]
    dataframe_list[1].columns = ['text','prediction','timestamp']

    #For sentiment analysis
    sentiment_columns = ['clean_text','prediction','timestamp']
    dataframe_list[2] = dataframe_list[2][sentiment_columns]
    dataframe_list[2].columns = ['text','prediction','timestamp']

    # Iterate and add id_model for each dataframe
    for model_id, df in enumerate(dataframe_list, start=1):
        df['model_id'] = model_id
    return dataframe_list

if __name__ == "__main__":
    try:
        #Get all dataframes
        model_results = get_data_from_containers()
        #Format dataframes
        model_results = format_dataframe(model_results)
        #Load All dataframes to Database
        table_names = ['Master_spam','Master_post','Master_sentiment']
        load_data_to_db(model_results,table_names,batch_size=1000)
    except Exception as e:
        print(f"Error al cargar datos desde el archivo: {e}")
    