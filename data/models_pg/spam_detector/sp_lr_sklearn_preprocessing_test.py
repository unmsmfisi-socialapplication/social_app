import unittest
import sys
from preprocessing.sp_lr_sklearn_preprocessing import clean_text

sys.path.append('C:/Users/USUARIO/Documents/GitHub/social_app/data')
class TestPreprocessing(unittest.TestCase):

    def test_clean_text(self):
        # Definir texto de prueba y resultado esperado
        sample_text = "<p>Hello, this is a sample text with some <strong>HTML</strong> tags!</p>"
        expected_result = "hello sample text html tags"

        # Llamar a la función clean_text y verificar si el resultado coincide con lo esperado
        self.assertEqual(clean_text(sample_text), expected_result)

    def test_clean_text_no_html(self):
        # Definir texto de prueba sin etiquetas HTML y resultado esperado
        sample_text = "This is a sample text without any HTML tags."
        expected_result = "sample text without html tags"

        # Llamar a la función clean_text y verificar si el resultado coincide con lo esperado
        self.assertEqual(clean_text(sample_text), expected_result)

    def test_clean_text_with_numbers(self):
        # Definir texto de prueba con números y resultado esperado
        sample_text = "There are 123 numbers in this text."
        expected_result = "numbers text"

        # Llamar a la función clean_text y verificar si el resultado coincide con lo esperado
        self.assertEqual(clean_text(sample_text), expected_result)

    def test_clean_text_empty_string(self):
        # Definir una cadena vacía y resultado esperado
        sample_text = ""
        expected_result = ""

        # Llamar a la función clean_text y verificar si el resultado coincide con lo esperado
        self.assertEqual(clean_text(sample_text), expected_result)

if __name__ == '__main__':
    unittest.main()
