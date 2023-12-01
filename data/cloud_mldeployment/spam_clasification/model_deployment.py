import os
from azureml.core import Workspace, Environment, Model
from azureml.core.model import InferenceConfig
from azureml.core.webservice import AciWebservice
os.environ['AZUREML_LOG_DEPRECATION_WARNING_ENABLED'] = 'False'


# Configura tu espacio de trabajo de Azure Machine Learning
ws = Workspace.from_config(path=r"C:\Users\USUARIO\Documents\GitHub\social_app\data\cloud_mldeployment\spam_clasification\config_azure.json")

# Configura el entorno de inferencia
env = Environment.from_conda_specification(name="spam_detection_environment", file_path="C:/Users/USUARIO/Documents/GitHub/social_app/data/cloud_mldeployment/spam_clasification/spam_detection_environment.yml")

# Configura el archivo de configuración de inferencia
inference_config = InferenceConfig(entry_script=r"C:\Users\USUARIO\Documents\GitHub\social_app\data\cloud_mldeployment\spam_clasification\inference_spam_model_logistic_regression.py", environment=env)

# Nombre del modelo registrado en Azure ML
model_name_in_azure = "logistic_regression_model"

# Obtiene el modelo registrado
model = Model(ws, name=model_name_in_azure)

# Despliega el servicio web ACI (Azure Container Instance)
deployment_config = AciWebservice.deploy_configuration(cpu_cores=1, memory_gb=1)
service_name = "sp-log-regres-service"  
service = Model.deploy(workspace=ws,
                       name=service_name,
                       models=[model],
                       inference_config=inference_config,
                       deployment_config=deployment_config)

service.wait_for_deployment(show_output=True)

