import json
from azure.identity import DefaultAzureCredential
from azure.mgmt.resource import ResourceManagementClient
from azureml.core import Workspace

# Cargar la configuración desde el archivo JSON
config_path = 'C:/Users/USUARIO/Documents/GitHub/social_app/data/notebook/spam_detector/config_azure.json'
with open(config_path) as f:
    config = json.load(f)

# Definir las credenciales de Azure
credential = DefaultAzureCredential()

# Definir las variables para el entorno de Azure Machine Learning
subscription_id = config.get("subscription_id")
resource_group_name = config.get("resource_group")
workspace_name = config.get("workspace_name")

# Crear el objeto de cliente de Azure Resource Management
resource_client = ResourceManagementClient(credential, subscription_id)

# Configurar el entorno de Azure Machine Learning
try:
    ws = Workspace.create(name=workspace_name,
                          subscription_id=subscription_id,
                          resource_group=resource_group_name,
                          create_resource_group=True,
                          location='eastus'
                          )
    print("Entorno de Azure Machine Learning configurado correctamente.")
except Exception as e:
    print(f"Error al configurar el entorno de Azure Machine Learning: {str(e)}")
