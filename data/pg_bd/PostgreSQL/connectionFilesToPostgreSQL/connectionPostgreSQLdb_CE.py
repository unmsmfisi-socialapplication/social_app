import psycopg2

from configParameters import config

def connect_azure_postgresql_ce():
    try:
        # Leyendo los parametros de la conexión
        params = config()
  
        # Conectando al servidor de postgreSQL
        print('-----------------------------------------------------------------------------------------------')
        conn = psycopg2.connect(**params)
        return conn
    except Exception as e:
        print(f'Error al conectar a la base de datos: {e}')
        return None