import unittest
from unittest.mock import patch, MagicMock
from io import StringIO
from azure.storage.blob import BlobServiceClient
from ..cloud_mldeployment.post_classification.azure_storage import uploadStorage

# Define a function to obtain the credentials, replace it with your credentials
def get_storage_credentials():
    return {
        "storage_account_key": "your account_key",
        "storage_account_name": "your account_name",
        "connection_string": "your connection_string",
        "container_name": "your container_name"
    }

class TestAzureStorageFunctions(unittest.TestCase):

    # Decorators to simulate the behavior of certain functions
    @patch('main.BlobServiceClient.from_connection_string')
    @patch('main.BlobClient.upload_blob')
    def test_uploadStorage(self, mock_upload_blob, mock_from_connection_string):
        # Simulates the behavior of BlobServiceClient and BlobClient
        mock_from_connection_string.return_value = MagicMock(spec=BlobServiceClient)
        mock_blob_client = mock_from_connection_string.return_value.get_blob_client.return_value
        mock_blob_client.upload_blob.return_value = None

        # Obtain credentials
        azure_credentials = get_storage_credentials()
        file_content = StringIO("id,clean_text,category_label,score_prediction\n1,4 russiancontrolled ukrainian regions schedule votes this week to join russia,WORLD NEWS,0.99\n2,james cameron says he clashed with studio before avatar release,ENTERTAINMENT,0.97\n")
    
        uploadStorage(file_content, 'post_classification.csv', azure_credentials)

        # Check that the BlobServiceClient and BlobClient were called correctly.
        mock_from_connection_string.assert_called_once_with(azure_credentials["connection_string"])
        mock_from_connection_string.return_value.get_blob_client.assert_called_once_with(azure_credentials["container_name"], blob='post_classification.csv')
        mock_blob_client.upload_blob.assert_called_once_with(file_content.getvalue().encode('utf-8'), overwrite=True)

if __name__ == '__main__':
    unittest.main()