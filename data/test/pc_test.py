import unittest
import pandas as pd
from mldeployment import pc_lr_main as main
from preprocessing import pc_lr_preprocessing as preprocessing
from training import pc_lr_training as training
import keras

# Test Dataframe creation
def test_data():
    return pd.DataFrame({
        'id': [1, 2, 3, 4],
        'clean_text': ['our hearts are broken historic front pages mark the queens death', 'dwight howard rips teammates after magic loss to hornets', 'episode 87 life is a highway audio', 'the psychometer who went too far last week'],
        'category_label': ['WORLD NEWS', 'SPORTS', 'ENTERTAINMENT', 'COMEDY']
    })

class TestMain(unittest.TestCase):

    def test_creation_model(self):
        model = main.creation_model(test_data())
        # Verify if the created model is keras
        self.assertIsInstance(model, keras.models.Sequential)

    def test_classification_post(self):
        df_predictions = main.classification_post(test_data())
        # Verifies if a dataframe is created
        self.assertIsInstance(df_predictions, pd.DataFrame)

class TestPreprocessing(unittest.TestCase):

    def test_preprocess_data(self):
        processed_data, lb_encoder, max_words = preprocessing.preprocess_data(test_data())
        # Verify if the outputs are correct
        self.assertIsInstance(processed_data, pd.DataFrame)
        self.assertIsInstance(lb_encoder, preprocessing.LabelEncoder)
        self.assertIsInstance(max_words, int)

class TestTraining(unittest.TestCase):

    def test_train_model(self):
        model = training.train_model(test_data())
        # Verify if the model is trained correctly
        self.assertIsNotNone(model)

if __name__ == '__main__':
    unittest.main()