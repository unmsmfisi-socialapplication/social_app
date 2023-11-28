from conexion_sql import connect_sql_database
import pandas as pd

#Connection testing

my_connection = connect_sql_database() #Connections instance

if my_connection is not None:
    sql = "SELECT * FROM [dbo].[bi_log]" # SQL query for testing

    #Init cursor query
    cursor = my_connection.cursor() 
    cursor.execute(sql)

    #Create dataset by cursor
    dataset = cursor.fetchall()
    
    columns = [column[0] for column in cursor.description]

    #Create dataframe instance
    df = pd.DataFrame(dataset,columns=columns)

    print(df)

    #Close all connection elements
    cursor.close()
    my_connection.close()