import pypyodbc as odbc
from credential import username,password #Credentials are protected in local variables
from sqlalchemy import create_engine
import urllib


#Config connection params to Azure server
database = 'socialdb_sql'
connection_string = 'Driver={ODBC Driver 18 for SQL Server};Server=tcp:sv-socialdb-sql.database.windows.net,1433;Database=socialdb_sql;Uid=' + username + ';Pwd=' + password +';Encrypt=yes;TrustServerCertificate=no;Connection Timeout=30;'

#Define connection method
def connect_sql_database():
    try:
        print('-----------------------------------------------------------------------------')
        conn = odbc.connect(connection_string)
        return conn #Return connection
    except Exception as e:
        print(f'Connection error: {e}')
        return None 

#Define alchemy connection
def connect_with_alchemy():
    try:
        params = urllib.parse.quote_plus(connection_string)
        conn_string = 'mssql+pyodbc:///?odbc_connect={}'.format(params) #Develop the connection string in url format
        engine_azure = create_engine(conn_string,echo=True) #Define engine

        return engine_azure
    except Exception as e:
        print(f'Connection error: {e}')
        return None 