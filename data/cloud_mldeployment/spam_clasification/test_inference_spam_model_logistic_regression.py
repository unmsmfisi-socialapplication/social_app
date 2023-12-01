import unittest
from unittest.mock import patch, Mock
from inference_spam_model_logistic_regression import clean_text, run

class TestInferenceSpamModel(unittest.TestCase):
    def test_clean_text(self):
        # Prueba para la función clean_text
        input_text = "This is an example text with HTML <tags> and special characters!"
        expected_output = "example text html tags special characters"
        cleaned_text = clean_text(input_text)
        self.assertEqual(cleaned_text, expected_output)

    @patch("inference_spam_model_logistic_regression.joblib.load")
    @patch("inference_spam_model_logistic_regression.joblib.load")
    def test_run(self, mock_load1, mock_load2):
        # Prueba para la función run

        # Configuración de las respuestas simuladas
        mock_load1.return_value = Mock()
        mock_load2.return_value = Mock()

        # Datos de entrada
        input_data = {'comment': 'This is a test comment.'}

        # Llamada a la función de inferencia
        result = run(input_data)

        # Verificación del resultado esperado
        expected_result = {"Logistic Regression Prediction": "not spam"}
        self.assertEqual(result, expected_result)

if __name__ == '__main__':
    unittest.main()
