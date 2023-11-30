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
        #cursor = connection.cursor()

    except Exception as ex:
        print(f"Error de conexión: {ex}")
    
    finally:
        # Close the connection at completion
        if connection:
            connection.close()
            print("Conexión cerrada")
