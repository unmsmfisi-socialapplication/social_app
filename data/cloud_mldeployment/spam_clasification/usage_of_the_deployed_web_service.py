import json
from azureml.core import Workspace
from azureml.core.webservice import AciWebservice

# Configura tu espacio de trabajo de Azure Machine Learning
ws = Workspace.from_config(path=r"C:\Users\USUARIO\Documents\GitHub\social_app\data\cloud_mldeployment\spam_clasification\config_azure.json")

# Nombre del servicio web desplegado
service_name = "sp-log-regres-service"

# Obtiene el servicio web desplegado
service = AciWebservice(workspace=ws, name=service_name)

# Datos de entrada para la inferencia
input_data = {
    "comment": "Ejemplo de texto para la predicción",
    "formatted_content": "Some formatted content"
}

# Realiza una solicitud HTTP para obtener la predicción
response = service.run(json.dumps(input_data))

# Imprime el resultado de la predicción
print(response)
