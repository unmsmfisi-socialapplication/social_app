import psycopg2
import pandas as pd

# Crea un dataframe de pandas para el ejemplo
df = pd.DataFrame({
    'columna1': [1, 2, 3],
    'columna2': ['dato1', 'dato2', 'dato3'],
    'columna3': [1.1, 2.2, 3.3]
})

# Datos de conexión a la base de datos PostgreSQL
dbname = 'dbname'
user = 'user'
password = 'password'
host = 'host'  # o la dirección del servidor de la base de datos
port = '5432'  # Puerto por defecto de PostgreSQL

# Conexión a la base de datos
conn = psycopg2.connect(
    dbname=dbname,
    user=user,
    password=password,
    host=host,
    port=port
)

# Crea una tabla temporal
with conn.cursor() as cur:
    cur.execute("""
        CREATE TABLE bi.temporal (
            columna1 INTEGER,
            columna2 TEXT,
            columna3 REAL
        );
    """)
    conn.commit()

    # Inserta datos del dataframe a la tabla temporal
    # Prepara una cadena de consulta SQL con la inserción de datos
    insert_query = 'INSERT INTO bi.temporal (columna1, columna2, columna3) VALUES %s;'
    
    # Utiliza execute_values que es más eficiente para insertar múltiples registros
    from psycopg2.extras import execute_values
    execute_values(cur, insert_query, df.values.tolist())
    
    conn.commit()

# Cierra la conexión
conn.close()