import pickle

class BasePredictor:
    def __init__(self, version, description):
        self.version = version
        self.description = description

    def predict(self, data):
        raise NotImplementedError("Subclasses must implement the predict method.")

class PostClassifier(BasePredictor):
    def predict(self, data):
        # Replace this with your actual post classification prediction logic
        return predict_post(data, self.version)

class SentimentAnalysisModel(BasePredictor):
    def predict(self, data):
        # Replace this with your actual sentiment analysis prediction logic
        return predict_sentiment(data, self.version)

class SpamDetector(BasePredictor):
    def predict(self, data):
        # Replace this with your actual spam detection prediction logic
        return predict_spam(data, self.version)

def load_model(model_path, model_class):
    try:
        with open(model_path, 'rb') as model_file:
            model = pickle.load(model_file)
            if not isinstance(model, model_class):
                raise ValueError(f"The loaded model is not an instance of {model_class.__name__}.")
        return model
    except FileNotFoundError:
        raise FileNotFoundError(f"Model file not found: {model_path}")
    except pickle.PickleError as e:
        raise RuntimeError(f"Error loading model: {e}")

def predict_post(model, data):
    # Replace this with your actual prediction logic for post classification
    prediction = model.predict(data)
    return {"post_classification": prediction}

def predict_sentiment(model, data):
    # Replace this with your actual prediction logic for sentiment analysis
    prediction = model.predict(data)
    return {"sentiment_analysis": prediction}

def predict_spam(model, data):
    # Replace this with your actual prediction logic for spam detection
    prediction = model.predict(data)
    return {"spam_detection": prediction}