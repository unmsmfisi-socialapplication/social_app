import json
from azureml.core import Workspace, Model
from azure.storage.blob import BlobServiceClient

# Ruta al archivo de configuración de Azure
config_azure_path = "C:/Users/USUARIO/Documents/GitHub/social_app/data/cloud_mldeployment/spam_clasification/config_azure.json"

# Lee las credenciales desde el archivo JSON
with open(config_azure_path, "r") as config_file:
    config_azure = json.load(config_file)

# Configuración del espacio de trabajo de Azure Machine Learning
subscription_id = config_azure["subscription_id"]
resource_group = config_azure["resource_group"]
workspace_name = config_azure["workspace_name"]

# Obtén las credenciales del espacio de trabajo
ws = Workspace(subscription_id=subscription_id, resource_group=resource_group, workspace_name=workspace_name)

# Nombre del contenedor en Azure Blob Storage
container_name = "spmodelcontenedor"  # Cambia al nombre correcto

# Nombre del archivo en Azure Blob Storage
blob_name = "logistic_regression_model.pkl"  # Puedes cambiar el nombre si lo deseas

# Carga el modelo desde el contenedor en Azure Blob Storage
blob_service_client = BlobServiceClient(account_url=f"https://{config_azure['storage_account_name']}.blob.core.windows.net", credential=config_azure['storage_account_key'])
container_client = blob_service_client.get_container_client(container_name)
blob_client = container_client.get_blob_client(blob_name)

# Descarga el modelo localmente
downloaded_model_path = "logistic_regression_model_spam.pkl"
with open(downloaded_model_path, "wb") as model_file:
    model_file.write(blob_client.download_blob().readall())

# Nombre que deseas darle al modelo en Azure ML
model_name_in_azure = "logistic_regression_model"

# Registra tu modelo en Azure Machine Learning
model = Model.register(model_path=downloaded_model_path,
                       model_name=model_name_in_azure,
                       workspace=ws)


