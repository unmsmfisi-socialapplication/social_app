import pickle

def load_post_classifier(model_path):
    with open(model_path, 'rb') as model_file:
        model = pickle.load(model_file)
    return model

def predict_post_classifier(model, data):
    # Replace this with your actual prediction logic for post classification
    prediction = model.predict(data)
    return {"post_classification": prediction}

def load_sentiment_analysis_model(model_path):
    with open(model_path, 'rb') as model_file:
        model = pickle.load(model_file)
    return model

def predict_sentiment_analysis(model, data):
    # Replace this with your actual prediction logic for sentiment analysis
    prediction = model.predict(data)
    return {"sentiment_analysis": prediction}

def load_spam_detector(model_path):
    with open(model_path, 'rb') as model_file:
        model = pickle.load(model_file)
    return model

def predict_spam_detector(model, data):
    # Replace this with your actual prediction logic for spam detection
    prediction = model.predict(data)
    return {"spam_detection": prediction}