from azureml.core import Model, Workspace
from azureml.core.authentication import InteractiveLoginAuthentication
import joblib
import json

# Cargar la configuración desde el archivo JSON
config_path = 'C:/Users/USUARIO/Documents/GitHub/social_app/data/notebook/spam_detector/config_azure.json'
with open(config_path) as f:
    config = json.load(f)

# Obtener las variables del archivo de configuración
workspace_name = config.get("workspace_name")
subscription_id = config.get("subscription_id")
resource_group_name = config.get("resource_group")

# Cargar el modelo entrenado desde el archivo local
save_path = 'C:/Users/USUARIO/Documents/GitHub/social_app/data/notebook/spam_detector/'
naive_bayes_classifier = joblib.load(save_path + 'naive_bayes_model.pkl')

# Configurar la autenticación interactiva (necesaria para algunas operaciones de SDK de Azure ML)
interactive_auth = InteractiveLoginAuthentication()

# Obtener la instancia del espacio de trabajo de Azure ML
ws = Workspace.get(
    name=workspace_name,
    subscription_id=subscription_id,
    resource_group=resource_group_name,
    auth=interactive_auth
)

# Definir el nombre y la descripción del modelo
model_name = 'spam_detection_model'
model_description = 'Modelo para detectar spam en comentarios'

# Registrar el modelo en Azure ML
model = Model.register(
    workspace=ws,
    model_name=model_name,
    model_path=save_path + 'naive_bayes_model.pkl',
    description=model_description,
    tags={'type': 'classification'}
)

print(f"Modelo registrado con éxito. Versión: {model.version}")
