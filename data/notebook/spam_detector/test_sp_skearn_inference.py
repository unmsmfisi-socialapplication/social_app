import unittest
from unittest.mock import patch, MagicMock
from sp_sklearn_inference import init, run, clean_text, naive_bayes_classifier, tfidf_vectorizer

class TestSpSklearnInference(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        # Configurar cualquier configuración común para todas las pruebas
        init()  # Llama a la función init antes de ejecutar las pruebas

    def setUp(self):
        # Configurar cualquier configuración específica de la prueba
        pass

    def test_init(self):
        # Simular la carga de modelos y vectorizadores durante la inicialización
        with patch('sp_sklearn_inference.Model.get_model_path') as mock_get_model_path:
            mock_get_model_path.return_value = 'dummy_path'

            init()

            # Verificar que las variables globales se hayan configurado correctamente
            self.assertIsNotNone(naive_bayes_classifier)
            self.assertIsNotNone(tfidf_vectorizer)
            
    def test_clean_text(self):
        # Probar la limpieza de texto
        text = "This is <html> some text 123 with symbols &amp; numbers!"
        cleaned_text = clean_text(text)

        # Verificar que el texto se haya limpiado correctamente
        self.assertEqual(cleaned_text, "text symbols numbers")

    def _simulate_run(self, raw_data, return_value):
        # Simular la carga de modelos y vectorizadores durante la ejecución
        with patch('sp_sklearn_inference.Model.get_model_path') as mock_get_model_path:
            mock_get_model_path.return_value = 'dummy_path'

            # Simular la predicción sin lanzar excepciones
            with patch('sp_sklearn_inference.joblib.load') as mock_load:
                mock_load.return_value = MagicMock()  # O mock_load.return_value = el modelo real

                # Ejecutar la función run
                result = run(raw_data)

                # Verificar que el resultado sea el esperado
                self.assertEqual(result, return_value)

    def test_run_not_spam(self):
        # Simular datos de entrada
        raw_data = '{"text": "Sample text for prediction"}'
        self._simulate_run(raw_data, {"result": "not spam"})

    def test_run_spam(self):
        # Simular datos de entrada para un caso de "spam"
        raw_data = '{"text": "Check out this amazing offer!"}'
        self._simulate_run(raw_data, {"result": "spam"})

    def test_run_exception(self):
        # Simular datos de entrada para un caso que lanza una excepción
        raw_data = '{"text": "Invalid data"}'

        # Simular la predicción que lanza una excepción
        with patch('sp_sklearn_inference.naive_bayes_classifier.predict') as mock_predict:
            mock_predict.side_effect = Exception("Error during prediction")

            # Ejecutar la función run
            result = run(raw_data)

            # Verificar que el resultado indique el error
            self.assertIn("error", result)

    def tearDown(self):
        
        pass

    @classmethod
    def tearDownClass(cls):
        
        pass

if __name__ == '__main__':
    unittest.main()
