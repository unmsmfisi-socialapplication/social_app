from fastapi import FastAPI, HTTPException
from azure.storage.blob import BlockBlobService
from models import (
    load_post_classifier,
    predict_post_classifier,
    load_sentiment_analysis_model,
    predict_sentiment_analysis,
    load_spam_detector,
    predict_spam_detector,
)

app = FastAPI()

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
    app.state.post_classifier = load_post_classifier("local_post_classifier.pkl")
    app.state.sentiment_analysis_model = load_sentiment_analysis_model("local_sentiment_analysis_model.pkl")
    app.state.spam_detector = load_spam_detector("local_spam_detector.pkl")
    
# Post Classification endpoint
@app.post("/post_classification/{version}")
def post_classification_endpoint(version: str, data: dict):
    if version == "v1":
        return predict_post_classifier(data)
    else:
        raise HTTPException(status_code=404, detail="Model version not found")

# Sentiment Analysis endpoint
@app.post("/sentiment_analysis/{version}")
def sentiment_analysis_endpoint(version: str, data: dict):
    if version == "v1":
        return predict_sentiment_analysis(data)
    else:
        raise HTTPException(status_code=404, detail="Model version not found")

# Spam Detector endpoint
@app.post("/spam_detector/{version}")
def spam_detector_endpoint(version: str, data: dict):
    if version == "v1":
        return predict_spam_detector(data)
    else:
        raise HTTPException(status_code=404, detail="Model version not found")