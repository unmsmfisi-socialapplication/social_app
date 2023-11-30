# main.py
import os
from dotenv import load_dotenv
import pandas as pd
from joblib import load
from preprocessing import preprocess_data
from storage import initialize_azure_storage_client, upload_to_azure_storage

import logging


def setup_logging():
    logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')


def load_environment_variables():
    load_dotenv()

    return (
        os.getenv('DATA_ENDPOINT'),
        os.getenv('AZURE_STORAGE_CONNECTION_STRING'),
        os.getenv('OUTPUT_LOCAL_FILE_PATH'),
        os.getenv('AZURE_STORAGE_CONTAINER_NAME'),
        os.getenv('AZURE_STORAGE_BLOB_NAME')
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
        data_endpoint, azure_storage_connection_string, output_local_file_path, container_name, blob_name = load_environment_variables()

        df_predictions = preprocess_and_predict(data_endpoint)

        df_predictions.to_csv(output_local_file_path, index=False)

        blob_service_client = initialize_azure_storage_client(azure_storage_connection_string)

        upload_to_azure_storage(blob_service_client, output_local_file_path, container_name, blob_name)

        logging.info("Process completed successfully.")

    except FileNotFoundError as e:
        logging.error(f"File not found: {e}")

    except Exception as e:
        logging.error(f"Unexpected error: {e}")


if __name__ == "__main__":
    main()
