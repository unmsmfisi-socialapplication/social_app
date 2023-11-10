import nltk
nltk.download('punkt')
import unittest
import pandas as pd
import numpy as np
from unittest.mock import patch 
from data.training.pc_lr_training import train_LSTM, evaluate_model, classify_post   

class TestPCLRTraining(unittest.TestCase):
    @patch('data.preprocessing.pc_lr_preprocessing.preprocess_data')
    @patch('keras.models.Sequential.fit')
    @patch('data.training.pc_lr_training.evaluate_model')
    @patch('data.training.pc_lr_training.classify_post')
    def test_train_LSTM(self, mock_preprocess_data, mock_fit, mock_evaluate_model, mock_classify_post):
        # Configure mocks to simulate model preprocessing and training
        mock_preprocess_data.return_value = (pd.DataFrame({'text_sequences': [[1, 2, 3]]}), 45, 100)
        mock_fit.return_value = None
        mock_evaluate_model.return_value = (0.85, 'food')
        mock_classify_post.return_value = 'food'

        # Test data for train_LSTM
        test_data = pd.DataFrame({'text_sequences': [[1, 2, 3]], 'category': ['food']})

        # Calls the model's training function
        accuracy, model = train_LSTM(test_data)

        # Verify that the preprocessing function has been called with the appropriate data
        mock_preprocess_data.assert_called_once_with(test_data)

        # Verify that the model's training function has been called with the appropriate data
        mock_fit.assert_called_once()

        # Verify that the model's evaluation function has been called with the appropriate data
        mock_evaluate_model.assert_called_once_with(model, mock_fit.return_value, mock_fit.return_value, 100)

        # Verifies that the correct accuracy and model are returned
        self.assertEqual(accuracy, 0.85)
        self.assertIsNotNone(model)

    @patch('numpy.argmax')
    def test_evaluate_model(self, mock_argmax):
        # Configure the mock to simulate the model's prediction 
        mock_argmax.return_value = 1

        # Test data for evaluate_model
        test_model = 'mock_model'
        test_X_test = np.array([[1, 2, 3]])
        test_y_test = np.array([1])
        test_Max_words = 100

        # Calls the model's evaluation function
        accuracy, category_id = evaluate_model(test_model, test_X_test, test_y_test, test_Max_words)
 
        # Verify that the argmax function has been called with the correct data
        mock_argmax.assert_called_once_with(test_model.predict(test_X_test)[0])

        #  Verifies that the correct accuracy and correct category are returned
        self.assertEqual(accuracy, 1.0)
        self.assertEqual(category_id, 1)

    def test_classify_post(self):
        # Test data for classify_post
        test_post = "Food is an art in Peru, the variety of dishes is varied."
        test_lstm_model = 'mock_model'
        test_max_sequence_length = 100
        test_categories = ['food', 'travel', 'lifestyle', 'other']

        # Call the post sorting feature
        category = classify_post(test_post, test_lstm_model, test_max_sequence_length, test_categories)

        # Verify that the returned category is correct
        self.assertEqual(category, 'food')

if __name__ == '__main__':
    unittest.main()