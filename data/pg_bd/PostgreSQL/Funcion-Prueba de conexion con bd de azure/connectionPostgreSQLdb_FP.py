import psycopg2

def connect_azure_postgresql_fp(host, database, user, password):
    try:
        conn = psycopg2.connect(
            host=host,
            database=database,
            user=user,
            password=password
        )
        return conn
    except Exception as e:
        print(f'Error al conectar a la base de datos: {e}')
        return None