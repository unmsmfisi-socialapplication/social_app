import pyodbc

import time


user =''
password =''
server = ''

def conexion():
    try:
        connection = pyodbc.connect('Driver={ODBC Driver 18 for SQL Server};'
                                    'Server='+server+',1433;'
                                    'Database=socialdb_sql;Uid='+user+';Pwd='+password+';Encrypt=yes;')
        print("Conexión exitosa")
        time.sleep(50000)

        # Create a cursor
        cursor = connection.cursor()

    except Exception as ex:
        print(f"Error de conexión: {ex}")
    
    finally:
        # Close the connection at completion
        if connection:
            connection.close()
            print("Conexión cerrada")
=======


# Connection Testing
try:
    connection = pyodbc.connect('Driver={ODBC Driver 18 for SQL Server};'
                                'Server=tcp:sv-socialdb-sql.database.windows.net,1433;'
                                'Database=socialdb_sql;Uid=useradmin;Pwd= password ;Encrypt=yes;')
    print("Conexión exitosa")


    # Create a cursor
    cursor = connection.cursor()

    # Run a test query
    cursor.execute("SELECT @@version;")
    row = cursor.fetchone()
    print("Versión del servidor de SQL Server: {}".format(row))

    # Run a query on the 'bi_log' table
    cursor.execute("SELECT * FROM bi_log")
    rows = cursor.fetchall()
    print(rows)
    for row in rows:
        print(row)

except Exception as ex:
    print(f"Error de conexión: {ex}")

finally:
    # Close the connection at completion
    if connection:
        connection.close()
        print("Conexión cerrada")
>>>>>>> 7be6b8235b5b8216bff234995dd3bcdd31742614
