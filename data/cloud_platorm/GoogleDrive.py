#from pydrive2.auth import GoogleAuth
#from pydrive2.drive import GoogleDrive
#from pydrive2.files import FileNotUploadedError
#import unittest

directorio_credenciales = 'credentials_module.json'

# log in
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
    archivo.SetContentString('Hey MoonCoders!')
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
    #Define a local file path for the test
    ruta_archivo = 'ruta_de_prueba/fondo_prueba.jpg'   

    #Define a test folder ID in Google Drive
    id_folder = 'ID_DE_LA_CARPETA_DE_PRUEBA'   

    #Execute the subir_archivo function with the test parameters
    subir_archivo(ruta_archivo, id_folder)

         


if __name__ == "__main__":
    #ruta_archivo = '/home/falv/Escritorio/fondo.jpg'
    #id_folder = '0AI_9cD6f9EEZUk9PVA'
    #id_drive = '1LVdc-DUwr30kfrA30cVO3K92RVh56pmw'
    #ruta_descarga = '/home/falv/Descargas/'
    #crear_archivo_texto('HolaDrive.txt','Hey MoonCoders',id_folder)
    #subir_archivo(ruta_archivo,id_folder)
    #unittest.main()
    
