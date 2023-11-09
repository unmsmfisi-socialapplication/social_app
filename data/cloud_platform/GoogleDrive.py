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

def download_file(filename,download_path):
    credential = login()
    list_files = credential.ListFile({'q': "title = '" + filename + "'"}).GetList()
    if not list_files:
        print('File not found: ' + filename)
    file = credential.CreateFile({'id': list_files[0]['id']}) 
    file.GetContentFile(download_path + filename)

def search(filename):
    result = []
    credential = login()
    list_files = credential.ListFile({'q': filename}).GetList()
    for file in list_files:
        print('ID Drive:', file['id'])
        print('File name:', file['title'])
        print('Type of file:', file['mimeType'])
        print('Size:', file['fileSize'])
        print('Date of creation:', file['createdDate'])
        print('Date of last modification:', file['modifiedDate'])
        print('Version:', file['version'])
        result.append(file)

    return result

def delete_recover(id_file):
    credential = login()
    file = credential.CreateFile({'id': id_file})
    
    file.Trash()

    file.UnTrash()

    file.Delete()

def create_folder(folder_name,id_folder):
    credential = login()
    folder = credential.CreateFile({'title': folder_name,
                                    'mimeType': 'application/vnd.google-apps.folder',
                                    'parents': [{'kind': 'drive#fileLink',
                                                'id': id_folder}]})
    folder.Upload()
 
if __name__ == "__main__":
    ruta_archivo = 'route_file'
    id_folder = 'id_folder'
    id_drive = 'id_drive'
    ruta_descarga = 'download_path'
    
    unittest.main()
    
    
