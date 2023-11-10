from azure.identity import DefaultAzureCredential
from azure.mgmt.resource import ResourceManagementClient
from azureml.core import Workspace

# Definir las credenciales de Azure
credential = DefaultAzureCredential()

# Definir las variables para el entorno de Azure Machine Learning
subscription_id = '4f283a8e-be85-4cdd-938c-3c6b05ed9b49'
resource_group_name = 'rg_myfirstaction'
workspace_name = 'ML_SPAM_RL'

# Crear el objeto de cliente de Azure Resource Management
resource_client = ResourceManagementClient(credential, subscription_id)

# Crear un grupo de recursos en Azure
resource_client.resource_groups.create_or_update(resource_group_name, {"location": "eastus"})

# Configurar el entorno de Azure Machine Learning
ws = Workspace.create(name=workspace_name,
                      subscription_id=subscription_id,
                      resource_group=resource_group_name,
                      create_resource_group=True,
                      location='eastus'
                      )
