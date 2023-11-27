# main.py
import os
from dotenv import load_dotenv
import pandas as pd
from joblib import load
from preprocessing import preprocess_data
from storage import initialize_firebase_app, upload_to_firebase_storage

import logging


def setup_logging():
    logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')


def load_environment_variables():
    load_dotenv()

    return (
        os.getenv('ENDPOINT'),
        os.getenv('FIREBASE_CREDENTIAL_PATH'),
        os.getenv('FIREBASE_STORAGE_URL'),
        os.getenv('FIREBASE_STORAGE_PATH'),
        os.getenv('OUTPUT_LOCAL_FILE_PATH')
    )


def preprocess_and_predict(endpoint):
    df, label_encoder = preprocess_data(endpoint)
    loaded_pipeline = load('logistic_regression_model.joblib')
    predictions = loaded_pipeline.predict(df['clean_text'])

    df_predictions = pd.DataFrame({
        'clean_text': df['clean_text'],
        'category': label_encoder.inverse_transform(df['label']),
        'prediction': label_encoder.inverse_transform(predictions),
        'probability': loaded_pipeline.predict_proba(df['clean_text']).max(axis=1)
    })

    return df_predictions.sort_values(by='probability', ascending=False)


def main():
    setup_logging()

    try:
        endpoint, firebase_credential_path, firebase_storage_url, firebase_storage_path, output_local_file_path = load_environment_variables()

        df_predictions = preprocess_and_predict(endpoint)

        df_predictions.to_csv(output_local_file_path, index=False)

        firebase_bucket = initialize_firebase_app(firebase_credential_path, firebase_storage_url)

        upload_to_firebase_storage(firebase_bucket, output_local_file_path, firebase_storage_path)

        logging.info("Process completed successfully.")

    except FileNotFoundError as e:
        logging.error(f"File not found: {e}")

    except Exception as e:
        logging.error(f"Unexpected error: {e}")


if __name__ == "__main__":
    main()
