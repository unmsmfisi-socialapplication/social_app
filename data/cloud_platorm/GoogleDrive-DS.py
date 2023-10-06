from pydrive2.auth import GoogleAuth
from pydrive2.drive import GoogleDrive
import unittest

credentials_directory = 'credentials_module.json'

def login():
    GoogleAuth.DEFAULT_SETTINGS['client_config_file'] = credentials_directory
    gauth = GoogleAuth()
    gauth.LoadCredentialsFile(credentials_directory)
    
    if gauth.credentials is None:
        gauth.LocalWebserverAuth(port_numbers=[8092])
    elif gauth.access_token_expired:
        gauth.Refresh()
    else:
        gauth.Authorize()
        
    gauth.SaveCredentialsFile(credentials_directory)
    credenciales = GoogleDrive(gauth)
    return credenciales

# File search function
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

# Function to download a file according to its name
def download_file(filename,download_path):
    credential = "GoogleAuth credentials"
    list_files = credential.ListFile({'q': "title = '" + filename + "'"}).GetList()
    if not list_files:
        print('File not found: ' + filename)
    file = credential.CreateFile({'id': list_files[0]['id']}) 
    file.GetContentFile(download_path + filename)

def test(self):
        # Assign file name for the test
        fileName = 'test_name.extension'   
         # Define a test folder ID in Google Drive
        download_path = 'test_path/file.extension'   
        download_file(fileName, download_path)

if __name__ == "__main__":
    route_file = 'route_file'
    id_folder = 'id_folder'
    id_drive = 'id_drive'
    download_path = 'download_path'
    search("title = 'name_file'")
    download_file('name_file',download_path)
    unittest.main()

