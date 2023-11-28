import pypyodbc as odbc
from credential import username,password #Credentials are protected in local variables

#Config connection params to Azure server
server = 'sv-socialdb-sql.database.windows.net'
database = 'socialdb_sql'
connection_string = 'Driver={ODBC Driver 18 for SQL Server};Server=tcp:sv-socialdb-sql.database.windows.net,1433;Database=socialdb_sql;Uid=' + username + ';Pwd=' + password +';Encrypt=yes;TrustServerCertificate=no;Connection Timeout=30;'

#Define connection method
def connect_sql_database():
    try:
        print('-----------------------------------------------------------------------------')
        conn = odbc.connect(connection_string)
        return conn #Return connection
    except Exception as e:
        print(f'Error al conectar a la base de datos: {e}')
        return None