
import pandas as pd
import requests
import io
import conexion_sql
from datetime import datetime

my_engine = conexion_sql.connect_with_alchemy()

file_url = "https://drive.google.com/uc?id=1FZxte6zEgjdGPrNLRlNLKZzln-w72M55"

def load_data_to_db(dataframe,batch_size=1000):
    table_name = 'Master_Sentiment' #Define the table name

    try:
        #Divide dataframe in batches of 1000 registers each
        total_rows = len(dataframe)
        num_batches = total_rows // batch_size + 1

        for i in range(num_batches):
            start_idx = i * batch_size
            end_idx = (i + 1) * batch_size
            batch = dataframe.iloc[start_idx:end_idx]

            #For each bacth the upload is made
            if not batch.empty:
                batch.to_sql(table_name, con=my_engine, index=False, if_exists='append') #Upload to Azure DB
                print(f"Lote {i + 1} de {num_batches} cargado correctamente.")
        print("Datos insertados correctamente.")
    except Exception as e:
        print(f"Error al insertar datos: {e}")
    

if __name__ == "__main__":
    try:
        #Get data in Cloud storage (local by now)
        response = requests.get(file_url)
        response.raise_for_status()  # Verify successful response
        temp_data = pd.read_csv(io.StringIO(response.content.decode('utf-8')))

        #Add id and timestamp
        temp_data['id_model'] = 1
        temp_data['timestamp'] = datetime.now().strftime("%Y-%m-%d %H:%M:%S.%f")[:-3]

        load_data_to_db(temp_data,batch_size=1000)
    except Exception as e:
        print(f"Error al cargar datos desde el archivo: {e}")
    