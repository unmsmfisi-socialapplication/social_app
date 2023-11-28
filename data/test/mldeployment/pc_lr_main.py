import pandas as pd
import numpy as np
import os
from ..preprocessing import pc_lr_preprocessing as preprocessing
from ..training import pc_lr_training as training
import joblib
from joblib import load
from keras.preprocessing.sequence import pad_sequences

# Endpoint to new posts for categorizing them
df = '' #Place endpoint backend
df = pd.read_csv(df)
# Endpoint to the post-base for training the model
df_initial = ''  #Place endpoint backend
df_initial = pd.read_csv(df_initial) 

# Function to create a post classification template
def creation_model(df_initial):
    model = training.train_model(df_initial)
    # The trained model is saved
    joblib.dump(model, 'lstm_model.joblib')
    return model

# Function for tagging posts with the previously trained model
def classification_post(df):

    df,lb_encoder,_= preprocessing.preprocess_data(df)
    # Loading and prediction of the model with preprocessed text
    load_model = load('lstm_model.joblib')
    sequences = pad_sequences(df['text_sequences'], maxlen=45)
    predictions = load_model.predict(sequences)

    # Assigns the most probable category identified by the model
    max_indices = np.argmax(predictions, axis=1)

    df_predictions = pd.DataFrame({
        'id': df['id'],
        'clean_text': df['clean_text'],
        'category_label': [lb_encoder.inverse_transform([idx])[0] for idx in max_indices],
        })

    # The posts classified according to the label assigned by the model are saved in a '.csv' file.
    df_predictions.to_csv('post_classification.csv', index=False)
    return df_predictions

if __name__ == "__main__":
    
    if os.path.exists('lstm_model.joblib'):
        classification_post(df)
    else:
        creation_model(df_initial)
    
    