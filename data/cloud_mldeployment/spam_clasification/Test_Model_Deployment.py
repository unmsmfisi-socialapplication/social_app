import unittest
from unittest.mock import patch, Mock
from azureml.core import Workspace, Environment
from azureml.core.model import InferenceConfig
from azureml.core.webservice import AciWebservice
from model_deployment import deploy_model

class TestModelDeployment(unittest.TestCase):
    @patch("Model_Deployment.Workspace.from_config")
    @patch("Model_Deployment.Environment.from_conda_specification")
    @patch("Model_Deployment.InferenceConfig")
    @patch("Model_Deployment.Model.deploy")
    def test_deploy_model(self, mock_workspace, mock_environment, mock_inference_config, mock_deploy_model):
        # Configuración de las respuestas simuladas
        mock_workspace.return_value = Mock(spec=Workspace)
        mock_environment.return_value = Mock(spec=Environment)
        mock_inference_config.return_value = Mock(spec=InferenceConfig)
        mock_deploy_model.return_value = Mock(spec=AciWebservice)

        # Llamada a la función que despliega el modelo
        deploy_model()

        # Verificación de llamadas a métodos simulados
        mock_workspace.from_config.assert_called_once()
        mock_environment.from_conda_specification.assert_called_once()
        mock_inference_config.assert_called_once()
        mock_deploy_model.assert_called_once()

if __name__ == '__main__':
    unittest.main()
