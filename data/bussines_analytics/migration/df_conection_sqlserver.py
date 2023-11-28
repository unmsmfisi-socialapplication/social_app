import pyodbc


# Prueba de conexión
try:
    connection = pyodbc.connect('Driver={ODBC Driver 18 for SQL Server};'
                                'Server=tcp:sv-socialdb-sql.database.windows.net,1433;'
                                'Database=socialdb_sql;Uid=useradmin;Pwd= password ;Encrypt=yes;')
    print("Conexión exitosa")


    # Crear un cursor
    cursor = connection.cursor()

    # Ejecutar una consulta de prueba
    cursor.execute("SELECT @@version;")
    row = cursor.fetchone()
    print("Versión del servidor de SQL Server: {}".format(row))

    # Ejecutar una consulta en la tabla 'bi_log'
    cursor.execute("SELECT * FROM bi_log")
    rows = cursor.fetchall()
    print(rows)
    for row in rows:
        print(row)

except Exception as ex:
    print(f"Error de conexión: {ex}")

finally:
    # Cerrar la conexión al finalizar
    if connection:
        connection.close()
        print("Conexión cerrada")
