import unittest
import pandas as pd
from pyspark.sql import SparkSession
from preprocessing.sa_lr_pyspark_preprocessing import preprocess_data

class PreprocessingTest(unittest.TestCase):
    def setUp(self):
        data = {
            'clean_text': [
                "This is a positive tweet.",
                "This is a neutral tweet.",
                "This is a negative tweet."
            ],
            'category': [1, 0, -1]
        }
        self.test_data_filename = 'test_data.csv'

        df = pd.DataFrame(data)

        df.to_csv(self.test_data_filename, index=False)

    def test_preprocess_data(self):
        spark = SparkSession.builder.appName("TestApp").getOrCreate()
        filename = self.test_data_filename
        dataset = preprocess_data(spark, filename)
        self.assertIsNotNone(dataset)
        self.assertGreaterEqual(dataset.count(), 0)

    def tearDown(self):
        import os
        os.remove(self.test_data_filename)

if __name__ == '__main__':
    unittest.main()
