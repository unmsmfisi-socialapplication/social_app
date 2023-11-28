# Preview is the instalation of psycopg2
# pip install psycopg2
import psycopg2

# Parámetros de conexión
db_config = {
    "host": "svbd-data-pg.postgres.database.azure.com",
    "database": "socialdb",
    "user": "useradmin",
    "password": "", # hidden key
    "sslmode": "require",  # Utiliza SSL para conexiones seguras
}

try:
    # Establecer la conexión a la base de datos
    conn = psycopg2.connect(**db_config)

    # Crear un cursor
    cursor = conn.cursor()

    # Ejecutar una consulta
    cursor.execute("SELECT * FROM comments.comment")

    # Recuperar los resultados
    rows = cursor.fetchall()
    for row in rows:
        print(row)

    # Cerrar el cursor y la conexión
    cursor.close()
    conn.close()

except Exception as e:
    print(f"Error de conexión: {e}")
