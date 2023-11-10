from connectionPostgreSQLdb_FP import connect_azure_postgresql_fp
from connectionPostgreSQLdb_CE import connect_azure_postgresql_ce


#TEST DE LA CONEXIÓN USANDO UNA FUNCION A LA QUE SE LE PASAN LAS CREDENCIALES COMO PARAMETROS

# Llama a la función con tus credenciales
my_connection = connect_azure_postgresql_fp(
    host='svbd-data-pg.postgres.database.azure.com',
    database='socialdb',
    user='useradmin',
    password=' '
)

# Verifica si la conexión se estableció correctamente
if my_connection is not None:
    cursor = my_connection.cursor()

    # Operaciones en la base de datos aquí
    print('TEST DE LA CONEXION MEDIANTE CREDENCIALES COMO PARAMETROS')
    print('Versión de la base de datos PostgreSQL:')
    cursor.execute('SELECT version()')
  
    # Mostrar la version de la base de datos postgreSQL instalada en el servidor
    db_version = cursor.fetchone()
    print(db_version)

    # Cerrar la conexión cuando se haya terminado
    cursor.close()
    my_connection.close()


#################################################################################################



#TEST DE LA CONEXIÓN USANDO UNA FUNCION QUE UTILIZA LAS CREDENCIALES EN EL ARCHIVO DATABASE.INI

# Llama a la función con tus credenciales
my_connection = connect_azure_postgresql_ce()

# Verifica si la conexión se estableció correctamente
if my_connection is not None:
    cursor = my_connection.cursor()

    # Haz operaciones en la base de datos aquí
    print('TEST DE LA CONEXION MEDIANTE CREENCIALES DEFINIDAS EN DATABASE.INI')
    print('Versión de la base de datos PostgreSQL:')
    cursor.execute('SELECT version()')
  
    # Mostrar la version de la base de datos postgreSQL instalada en el servidor
    db_version = cursor.fetchone()
    print(db_version)

    # Cerrar la conexión cuando se haya terminado
    cursor.close()
    my_connection.close()
