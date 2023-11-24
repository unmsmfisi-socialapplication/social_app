import unittest
import pandas as pd
from unittest.mock import patch
from ...execution import pc_lr_main
from ...training import pc_lr_training

#PC: Post Clasification
class Test_Model_PC(unittest.TestCase):  
    @patch('pandas.read_csv')
    def test_main_function_with_mocked_data(self, mock_read_csv, mock_train_lstm):
        # Configure mocks to simulate CSV reading and model training
        mock_read_csv.return_value = pd.DataFrame({'column1': [1, 2, 3], 'column2': [4, 5, 6]})
        mock_train_lstm.return_value = 'mocked_model'

        # Configure the dummy URL for the API
        pc_lr_main.api_url = 'fake_api_url'

        # Call the main function of the script
        pc_lr_main.main()

        # Verify that the CSV read function was called with the correct URL
        mock_read_csv.assert_called_once_with('fake_api_url', sep=',')

        # Verify that the model's training function has been called with the appropriate data
        mock_train_lstm.assert_called_once_with(pd.DataFrame({'column1': [1, 2, 3], 'column2': [4, 5, 6]}))

if __name__ == '__main__':
    unittest.main()