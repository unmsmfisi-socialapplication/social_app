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
        cursor = conn.cursor()
        query_temp = """SELECT * FROM API_prueba WHERE TIPO = '"""+type+"""' """
        #SQL query to get the data
        cursor.execute(query_temp)
        rows = cursor.fetchall()

        #Convert to json
        posts = []
        for row in rows:
            posts.append({'id': row[0], 'Nombre_Usuraio': row[1], 'Comentario': row[2],'Tipo':row[3],'Clasificacion':row[4],'Hora_publicacion':row[5]})

        return jsonify(posts)

    except Exception as e:
        return str(e)

if __name__ == '__main__':
    app.run(debug=True)
