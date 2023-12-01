import pyodbc
import unittest
#from unittest.mock import patch, MagicMock
#import io
#import sys
from df_conection_sqlserver import conexion

class Test_AzureSqlConnection(unittest.TestCase):
    # Connection Testing
    def test_connection(self):
        connection = None  # Inicializar la variable connection fuera del bloque try

        try:
            connection = pyodbc.connect('Driver={ODBC Driver 18 for SQL Server};'
                                        'Server=' + conexion.server + ',1433;'
                                        'Database=socialdb_sql;Uid=' + conexion.user + ';Pwd=' + conexion.password + ';Encrypt=yes;')
            print("Conexión exitosa")

            cursor = connection.cursor()

            cursor.execute("SELECT * FROM bi_log")
            rows = cursor.fetchall()
            print(rows)
            for row in rows:
                print(row)

        except Exception as ex:
            print(f"Error de conexión: {ex}")

        finally:
            if connection:
                connection.close()
                print("Conexión cerrada")

if __name__ == '__main__':
    unittest.main()
