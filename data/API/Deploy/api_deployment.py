from flask import Flask, jsonify, request  
import os
import pyodbc

app = Flask(__name__)

#Enviroments Variables 
server = os.environ.get('AZURE_SQL_SERVER')
database = os.environ.get('AZURE_SQL_DATABASE')
username = os.environ.get('AZURE_SQL_USERNAME')
password = os.environ.get('AZURE_SQL_PASSWORD')
driver = os.environ.get('AZURE_SQL_DRIVER', '{ODBC Driver 17 for SQL Server}')


# Ruta para obtener los datos
@app.route('/getPosts', methods=['GET'])
def get_posts():
    try:
        type = request.args.get('type', default=None, type=str)
        #Conection
        conn = pyodbc.connect(f'DRIVER={driver};SERVER={server};PORT=1433;DATABASE={database};UID={username};PWD={password}')
        
        cursor2 = conn.cursor()
        tmp= """EXECUTE [dbo].[SP_GetDataAPI] """+type+""";"""
        cursor2.execute(tmp)
        cursor2.close()


        cursor = conn.cursor()
        query_temp = """SELECT * FROM Response_Matrix"""
        #SQL query to get the data
        cursor.execute(query_temp)
        rows = cursor.fetchall()

        #Convert to json
        posts = []
        for row in rows:
            posts.append({'id': row[0], 'id_model': row[1], 'Text': row[2],'Prediction':row[3]})

        return jsonify(posts)

    except Exception as e:
        return str(e)

if __name__ == '__main__':
    app.run(debug=True)
