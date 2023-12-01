import unittest
from unittest.mock import patch, Mock
from azureml.core import Workspace
from azure.storage.blob import BlobServiceClient
from setting_up_the_connection_to_Azure_Storage import setup_azure_storage_connection

class TestAzureStorageConnection(unittest.TestCase):
    @patch("azureml.core.Workspace")
    @patch("azure.storage.blob.BlobServiceClient")
    def test_setup_azure_storage_connection(self, mock_workspace, mock_blob_service_client):
        # Configuración de las respuestas simuladas
        config_azure = {
            "subscription_id": "your_subscription_id",
            "resource_group": "your_resource_group",
            "workspace_name": "your_workspace_name",
            "storage_account_name": "your_storage_account_name",
            "storage_account_key": "your_storage_account_key"
        }

        mock_workspace.return_value = Mock(spec=Workspace)
        mock_blob_service_client.from_connection_string.return_value = Mock(spec=BlobServiceClient)

        # Llamada a la función que configura la conexión a Azure Storage
        setup_azure_storage_connection(config_azure)

        # Verificación de llamadas a métodos simulados
        mock_workspace.assert_called_once_with(
            subscription_id=config_azure["subscription_id"],
            resource_group=config_azure["resource_group"],
            workspace_name=config_azure["workspace_name"]
        )
        mock_blob_service_client.from_connection_string.assert_called_once_with(
            account_url=f"https://{config_azure['storage_account_name']}.blob.core.windows.net",
            credential=config_azure['storage_account_key']
        )

if __name__ == '__main__':
    unittest.main()
