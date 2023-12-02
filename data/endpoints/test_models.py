import unittest
from unittest.mock import Mock
from models import load_model, PostClassifier, SentimentAnalysisModel, SpamDetector

class TestModels(unittest.TestCase):
    def setUp(self):
        # Set up mock data for testing
        self.mock_blob_service = Mock()
        self.container_name = "test_container"
        self.model_names = ["post_classifier.pkl", "sentiment_analysis_model.pkl", "spam_detector.pkl"]
        self.local_model_paths = ["local_post_classifier.pkl", "local_sentiment_analysis_model.pkl", "local_spam_detector.pkl"]

    def test_load_post_classifier(self):
        for model_name, local_model_path in zip(self.model_names, self.local_model_paths):
            model = load_model(self.mock_blob_service, self.container_name, model_name, local_model_path, PostClassifier)
            self.assertIsInstance(model, PostClassifier)

    def test_load_sentiment_analysis_model(self):
        for model_name, local_model_path in zip(self.model_names, self.local_model_paths):
            model = load_model(self.mock_blob_service, self.container_name, model_name, local_model_path, SentimentAnalysisModel)
            self.assertIsInstance(model, SentimentAnalysisModel)

    def test_load_spam_detector(self):
        for model_name, local_model_path in zip(self.model_names, self.local_model_paths):
            model = load_model(self.mock_blob_service, self.container_name, model_name, local_model_path, SpamDetector)
            self.assertIsInstance(model, SpamDetector)

if __name__ == '__main__':
    unittest.main()
