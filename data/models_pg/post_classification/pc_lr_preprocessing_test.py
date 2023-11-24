import unittest
import pandas as pd
from unittest.mock import patch
from data.preprocessing import pc_lr_preprocessing  

#PC: Post Clasification
class Test_Preprocessing_PC(unittest.TestCase):
    @patch('nltk.corpus.stopwords.words')
    @patch('tensorflow.keras.preprocessing.text.Tokenizer.texts_to_sequences')
    def test_preprocess_data(self, mock_stopwords, mock_texts_to_sequences):
        # Configure mocks to simulate the behavior of external functions
        mock_stopwords.return_value = ['common', 'stop', 'words']
        mock_texts_to_sequences.return_value = [[1, 2, 3], [4, 5, 6]]

        # Test data for preprocess_data
        test_data = pd.DataFrame({'description': ['Test post 1', 'Test post 2']})

        # Call the preprocessing function
        preprocessed_data, input_shape, Max_words = pc_lr_preprocessing.preprocess_data(test_data)

        # Verify that the stopwords function was called in the correct language
        mock_stopwords.assert_called_once_with('english')

        #  Verify that the processed data dataframe has the appropriate columns
        self.assertTrue(pc_lr_preprocessing.CLEANED_COLUMN in preprocessed_data.columns)
        self.assertTrue(pc_lr_preprocessing.SEQUENCES_COLUMN in preprocessed_data.columns)

        #  Verifies that the text-to-sequence function has been called with the appropriate data
        mock_texts_to_sequences.assert_called_once_with(preprocessed_data[pc_lr_preprocessing.CLEANED_COLUMN])

        # Verify that input_shape and Max_words have expected values
        self.assertEqual(input_shape, 3)  
        self.assertEqual(Max_words, 7)  

if __name__ == '__main__':
    unittest.main()