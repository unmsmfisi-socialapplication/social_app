# test_storage.py
import unittest
from unittest.mock import Mock, patch
from cloud_mldeployment.sentiment_analysis.storage import initialize_azure_storage_client, upload_to_azure_storage

class TestAzureStorageClientInitialization(unittest.TestCase):

    def setUp(self):
        # Configure test data
        self.connection_string = "fake_connection_string"
        self.local_file_path = "fake_local_file_path.txt"
        self.container_name = "fake_container"
        self.blob_name = "fake_blob_name"

        # Configures the file simulation
        with patch("builtins.open", create=True) as mock_file:
            # Defines the content of the simulated file
            file_content = "fake_file_content"
            mock_file.return_value.read.return_value = file_content
            # Assigns the mock object to an instance variable so that it is accessible in other methods
            self.mock_file = mock_file()

    def test_initialize_azure_storage_client(self):
        # Act
        with unittest.mock.patch("azure.storage.blob.BlobServiceClient") as mock_blob_service_client:
            initialize_azure_storage_client(self.connection_string)

            # Assert
            mock_blob_service_client.from_connection_string.assert_called_once_with(self.connection_string)

    def test_upload_to_azure_storage(self):
        # Creates a Mock object for BlobServiceClient, ContainerClient, and BlobClient
        blob_service_client_mock = Mock()
        container_client_mock = Mock()
        blob_client_mock = Mock()

        # Configures the methods and attributes required for blob loading
        blob_service_client_mock.get_container_client.return_value = container_client_mock
        container_client_mock.get_blob_client.return_value = blob_client_mock

        # Call the function with the Mock objects
        with patch("builtins.open", self.mock_file):
            upload_to_azure_storage(blob_service_client_mock, self.local_file_path, self.container_name, self.blob_name)

        # Make sure that the necessary methods are called
        container_client_mock.get_blob_client.assert_called_once_with(self.blob_name)
        blob_client_mock.upload_blob.assert_called_once()

if __name__ == '__main__':
    unittest.main()
