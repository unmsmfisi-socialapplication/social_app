# test_preprocessing.py
import os
import pandas as pd
import unittest
from cloud_platform.sentiment_analysis.preprocessing import preprocess_data


class TestPreprocessing(unittest.TestCase):
    csv_path = 'test_mock.csv'
    empty_file_path = 'empty_file_mock.csv'

    @classmethod
    def setUpClass(cls):
        data = {
            'clean_text': [
                "This is a positive tweet.",
                "This is a neutral tweet.",
                "This is a negative tweet."
            ],
            'category': [1, 0, -1]
        }

        df = pd.DataFrame(data)

        df.to_csv(cls.csv_path, index=False)

        open(cls.empty_file_path, 'w').close()

    def test_preprocess_data(self):
        processed_df, label_encoder = preprocess_data(self.csv_path)

        self.assertTrue(processed_df.isnull().sum().sum() == 0)
        self.assertTrue('label' in processed_df.columns)
        self.assertTrue(len(label_encoder.classes_) > 0)

    def test_preprocess_data_with_empty_file(self):
        with self.assertRaises(Exception):
            preprocess_data(self.empty_file_path)

    @classmethod
    def tearDownClass(cls):
        os.remove(cls.empty_file_path)
        os.remove(cls.csv_path)


if __name__ == '__main__':
    unittest.main()
