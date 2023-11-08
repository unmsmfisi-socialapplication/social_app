import unittest
import pandas as pd
from pyspark.sql import SparkSession
from preprocessing import preprocess_data

class PreprocessingTest(unittest.TestCase):
    def setUp(self):
        # Crear datos de muestra en tiempo de ejecución
        data = {
            'clean_text': [
                "This is a positive tweet.",
                "This is a neutral tweet.",
                "This is a negative tweet."
            ],
            'category': [1, 0, -1]
        }
        self.test_data_filename = 'test_data.csv'

        # Crear un DataFrame de pandas con los datos de muestra
        df = pd.DataFrame(data)

        # Guardar el DataFrame como un archivo CSV
        df.to_csv(self.test_data_filename, index=False)

    def test_preprocess_data(self):
        spark = SparkSession.builder.appName("TestApp").getOrCreate()
        filename1 = self.test_data_filename  # Utiliza el archivo de muestra generado en tiempo de ejecución
        filename2 = self.test_data_filename  # Utiliza el archivo de muestra generado en tiempo de ejecución
        dataset = preprocess_data(spark, filename1, filename2)
        self.assertIsNotNone(dataset)
        self.assertGreaterEqual(dataset.count(), 0)

    def tearDown(self):
        # Elimina el archivo de prueba generado en tiempo de ejecución después de las pruebas
        import os
        os.remove(self.test_data_filename)

if __name__ == '__main__':
    unittest.main()
