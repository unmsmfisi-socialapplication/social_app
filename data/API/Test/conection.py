import os
import pyodbc
import pandas as pd

#Variables enviroments
server = os.environ.get('AZURE_SQL_SERVER')
database = os.environ.get('AZURE_SQL_DATABASE')
username = os.environ.get('AZURE_SQL_USERNAME')
password = os.environ.get('AZURE_SQL_PASSWORD')
driver = os.environ.get('AZURE_SQL_DRIVER', '{ODBC Driver 17 for SQL Server}')

#Conection
conn = pyodbc.connect(f'DRIVER={driver};SERVER={server};PORT=1433;DATABASE={database};UID={username};PWD={password}')

#querytes
query_temp = """
SELECT * FROM API_prueba
"""
#Read table to Dataframe
df_act= pd.read_sql_query(query_temp,conn)

#print DataTable
print(df_act)