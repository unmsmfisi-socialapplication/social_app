from sklearn.preprocessing import LabelEncoder
from nltk.tokenize import word_tokenize
from tensorflow import keras
import re
import pandas as pd
import numpy as np
from joblib import load
from keras.preprocessing.sequence import pad_sequences
from azure_storage import uploadStorage
import io

# Parameterized dataframe columns
SEQUENCES_COLUMN = 'text_sequences'
POST_COLUMN = 'clean_text'
LABEL_ENCODER = 'label'
TEXT_TOKENIZED = 'text_tokenized'
CATEGORY_LABEL = 'category_label'

def preprocess_data(data):
    # Post cleaning
    data[POST_COLUMN] = data[POST_COLUMN].apply(lambda x: re.sub(r'[^\w\s]', '', x))
    data[POST_COLUMN] = data[POST_COLUMN].apply(lambda x: x.lower())
    data[TEXT_TOKENIZED] = data[POST_COLUMN].apply(word_tokenize)

    lb_encoder = LabelEncoder()
    data[LABEL_ENCODER] = lb_encoder.fit_transform(data[CATEGORY_LABEL])

    # Text tokenization using Keras
    tokenizer = keras.preprocessing.text.Tokenizer()
    tokenizer.fit_on_texts(data[TEXT_TOKENIZED])

    sequences = tokenizer.texts_to_sequences(data[TEXT_TOKENIZED])

    data[SEQUENCES_COLUMN] = sequences

    return data, lb_encoder

def classification_post(df):
    # The data are cleaned to give them the proper format.
    df,lb_encoder = preprocess_data(df)

    # Loading and prediction of the model with preprocessed text
    load_model = load('lstm_model.joblib')
    sequences = pad_sequences(df['text_sequences'], maxlen=45)
    predictions = load_model.predict(sequences)

    # Selects the percentage of classification accuracy in a range 0-1
    max_values = np.max(predictions, axis=1)
    # The most probable category identified by the model is identified.
    max_indices = np.argmax(predictions, axis=1)

    df_predictions = pd.DataFrame({
        'id': df['id'],
        'clean_text': df['clean_text'],
        'category_label': [lb_encoder.inverse_transform([idx])[0] for idx in max_indices],
        'score_prediction': max_values
        })

    # The classified posts are saved in a .csv file according to the label assigned by the model.
    return df_predictions.to_csv(index=False)

if __name__ == "__main__":
    # Route file with the posts to be sorted
    df = pd.read_csv('post_classification_initial.csv')
    csv_content = classification_post(df)
    # A temporary file is created with the content of the classified posts.
    with io.StringIO(csv_content) as tmp_buffer:
        # Send content to Azure Storage
        uploadStorage(tmp_buffer, 'post_classification')