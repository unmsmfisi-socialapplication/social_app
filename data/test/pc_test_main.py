import unittest
import pandas as pd
from unittest.mock import patch, MagicMock
from io import StringIO
from ..cloud_mldeployment.post_classification.main import classification_post

class TestMainFunctions(unittest.TestCase):

    # Creation of the test dataframe
    def setUp(self):
        self.test_data = pd.DataFrame({
            'id': [13, 21],
            'clean_text': ['4 RussianControlled Ukrainian Regions Schedule Votes This Week To Join Russia', 'James Cameron Says He Clashed With Studio Before Avatar Release'],
            'category_label': ['WORLD NEWS', 'ENTERTAINMENT']
        })

    # Decorator simulating the function reference
    @patch('main.classification_post')
    # Tests the function that sorts the posts by the value it returns or if it was called correctly
    def test_classification_post(self, mock_classification_post):
        csv_content = "id,clean_text,category_label,score_prediction\n1,4 russiancontrolled ukrainian regions schedule votes this week to join russia,WORLD NEWS,0.99\n2,james cameron says he clashed with studio before avatar release,ENTERTAINMENT,0.97\n"
        mock_classification_post.return_value = csv_content
        result = classification_post(self.test_data)
        mock_classification_post.assert_called_once_with(self.test_data)
        self.assertEqual(result, csv_content)
        self.assertIsInstance(result, str)

    # Decorators to temporarily simulate functions or references
    @patch('main.io.StringIO')
    @patch('main.uploadStorage')
    @patch('main.__import__')

    # Test the main function where a series of values are returned
    def test_main(self, mock_import, mock_uploadStorage, mock_stringIO):
        mock_uploadStorage.return_value = MagicMock()
        mock_stringIO.return_value = StringIO("id,clean_text,category_label,score_prediction\n1,4 russiancontrolled ukrainian regions schedule votes this week to join russia,WORLD NEWS,0.99\n2,james cameron says he clashed with studio before avatar release,ENTERTAINMENT,0.97\n")
        csv_content = mock_stringIO.return_value
  
        mock_import.assert_called_once_with('main')
        mock_uploadStorage.assert_called_once()
        mock_stringIO.assert_called_once_with(csv_content)
        mock_stringIO().getvalue.assert_called_once()

if __name__ == '__main__':
    unittest.main()
