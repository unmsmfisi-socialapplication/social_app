from configparser import ConfigParser
  
def config(filename='database.ini', section='postgresql'):
    # Creando un parser
    parser = ConfigParser()
    # Leyendo el archivo database con las credenciales de acceso
    parser.read(filename)
  
    # Obtener sección, por defecto de postgresql
    db = {}
    if parser.has_section(section):
        params = parser.items(section)
        for param in params:
            db[param[0]] = param[1]
    else:
        raise Exception('Section {0} not found in the {1} file'.format(section, filename))
  
    return db