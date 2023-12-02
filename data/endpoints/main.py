from fastapi import FastAPI
from azure.storage.blob import BlockBlobService
from models import load_model, PostClassifier, SentimentAnalysisModel, SpamDetector

app = FastAPI()

def download_model_from_azure(blob_service, container_name, model_name, local_model_path):
    blob_service.get_blob_to_path(container_name, model_name, local_model_path)

@app.on_event("startup")
async def startup_event():
    # Download models from Azure Blob Storage
    azure_storage_account_name = "your_storage_account_name"
    azure_storage_account_key = "your_storage_account_key"
    container_name = "your_container_name"
    model_names = ["post_classifier.pkl", "sentiment_analysis_model.pkl", "spam_detector.pkl"]
    local_model_paths = ["local_post_classifier.pkl", "local_sentiment_analysis_model.pkl", "local_spam_detector.pkl"]

    blob_service = BlockBlobService(account_name=azure_storage_account_name, account_key=azure_storage_account_key)

    for model_name, local_model_path in zip(model_names, local_model_paths):
        download_model_from_azure(blob_service, container_name, model_name, local_model_path)

    # Load models on application startup
    app.state.post_classifier = load_model("local_post_classifier.pkl", PostClassifier)
    app.state.sentiment_analysis_model = load_model("local_sentiment_analysis_model.pkl", SentimentAnalysisModel)
    app.state.spam_detector = load_model("local_spam_detector.pkl", SpamDetector)
    
# Post Classification endpoint Define endpoints for predictions
@app.post("/post_classification")
def post_classification_endpoint(data: dict):
    prediction = app.state.post_classifier.predict(data)
    return {"post_classification": prediction}

@app.post("/sentiment_analysis")
def sentiment_analysis_endpoint(data: dict):
    prediction = app.state.sentiment_analysis_model.predict(data)
    return {"sentiment_analysis": prediction}

@app.post("/spam_detector")
def spam_detector_endpoint(data: dict):
    prediction = app.state.spam_detector.predict(data)
    return {"spam_detection": prediction}