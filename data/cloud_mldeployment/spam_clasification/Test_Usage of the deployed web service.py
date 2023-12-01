import unittest
from unittest.mock import patch, Mock
import json
from azureml.core import Workspace
from azureml.core.webservice import AciWebservice
from usage_of_the_deployed_web_service import get_azure_ml_service_response

class TestAzureMLService(unittest.TestCase):
    @patch("azureml.core.Workspace")
    @patch("azureml.core.webservice.AciWebservice")
    def test_get_azure_ml_service_response(self, mock_workspace, mock_aci_webservice):
        # Configuración de las respuestas simuladas
        mock_workspace.return_value = Mock(spec=Workspace)
        mock_aci_webservice.return_value = Mock(spec=AciWebservice)
        mock_aci_webservice.run.return_value = {"result": "mocked_result"}

        # Llamada a la función que interactúa con el servicio web de Azure ML
        result = get_azure_ml_service_response()

        # Verificación de llamadas a métodos simulados
        mock_workspace.assert_called_once_with.from_config(path="path/to/config.json")
        mock_aci_webservice.assert_called_once_with(workspace=mock_workspace.return_value, name="sp-log-regres-service")
        mock_aci_webservice.run.assert_called_once_with(json.dumps({"comment": "Ejemplo de texto para la predicción", "formatted_content": "Some formatted content"}))

        # Verificación del resultado
        self.assertEqual(result, {"result": "mocked_result"})

if __name__ == '__main__':
    unittest.main()
