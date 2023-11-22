import json
from azure.identity import DefaultAzureCredential
from azure.mgmt.resource import ResourceManagementClient
from azureml.core import Workspace

# Cargar la configuración desde el archivo JSON
with open('C:/Users/USUARIO/Documents/GitHub/social_app/data/notebook/spam_detector/config_azure.json') as f:
    config = json.load(f)

# Definir las credenciales de Azure
credential = DefaultAzureCredential()

# Definir las variables para el entorno de Azure Machine Learning
subscription_id = config["subscription_id"]
resource_group_name = config["resource_group"]
workspace_name = config["workspace_name"]

# Crear el objeto de cliente de Azure Resource Management
resource_client = ResourceManagementClient(credential, subscription_id)

# Crear un grupo de recursos en Azure (esto solo es necesario si aún no existe)
resource_client.resource_groups.create_or_update(resource_group_name, {"location": "eastus"})

# Configurar el entorno de Azure Machine Learning
ws = Workspace.create(name=workspace_name,
                      subscription_id=subscription_id,
                      resource_group=resource_group_name,
                      create_resource_group=True,
                      location='eastus'
                      )
