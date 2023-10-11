from pydrive2.auth import GoogleAuth
from pydrive2.drive import GoogleDrive 

import unittest

directorio_credenciales = 'credentials_module.json'

# Login
def login():
    GoogleAuth.DEFAULT_SETTINGS['client_config_file'] = directorio_credenciales
    gauth = GoogleAuth()
    gauth.LoadCredentialsFile(directorio_credenciales)
    
    if gauth.credentials is None:
        gauth.LocalWebserverAuth(port_numbers=[8092])
    elif gauth.access_token_expired:
        gauth.Refresh()
    else:
        gauth.Authorize()
        
    gauth.SaveCredentialsFile(directorio_credenciales)
    credenciales = GoogleDrive(gauth)
    return credenciales

#create file
def crear_archivo_texto(nombre_archivo,contenido,id_folder):
    credenciales = login()
    archivo = credenciales.CreateFile({'title': nombre_archivo,\
                                        'parents': [{"kind": "drive#fileLink",\
                                        "id": id_folder}]})
    archivo.SetContentString('comment')
    archivo.Upload()


# UPLOAD A FILE TO DRIVE
def subir_archivo(ruta_archivo,id_folder):
    credenciales = login()
    archivo = credenciales.CreateFile({'parents': [{"kind": "drive#fileLink",\
                                                    "id": id_folder}]})
    archivo['title'] = ruta_archivo.split("/")[-1]
    archivo.SetContentFile(ruta_archivo)
    archivo.Upload()


def test_subir_archivo(self):
    # Define a local file path for the test
    ruta_archivo = 'route_file'   
     # Define a test folder ID in Google Drive
    id_folder = 'id_folder'   
    # Execute the subir_archivo function with the test parameters
    subir_archivo(ruta_archivo, id_folder)

         
 
if __name__ == "__main__":
    ruta_archivo = 'route_file'
    id_folder = 'id_folder'
    id_drive = 'id_drive'
    ruta_descarga = 'download_path'
    #crear_archivo_texto('file','commet',id_folder)
    subir_archivo(ruta_archivo,id_folder)
    unittest.main()
    
    
