import pandas as pd 
import numpy as np
import keras
from keras.preprocessing.sequence import pad_sequences
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import LabelEncoder
from sklearn.metrics import accuracy_score
from ..preprocessing import pc_lr_preprocessing

# Parameterized variable
DATA_PROCESSED_COLUMN ='text_sequences'
EPOCHS_NUMBER = 25
BATCH_NUMBER=32
SPLIT_NUMBER = 0.2
TEST_NUMBER = 0.2
STATE_NUMBER = 100
UNIT_NUMBER = 16
DROPOUT_NUMBER = 0.25
DENSE_NUMBER = 4

def train_LSTM(data):
    # Values are taken from the data processed for model training.
    preprocessed_data, input_shape, Max_words = pc_lr_preprocessing.preprocess_data(data)

    X = pad_sequences(preprocessed_data[DATA_PROCESSED_COLUMN], maxlen=45)
    y = LabelEncoder().fit_transform(data['category'])

    # Divide the dataset into test and training for the model.
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=TEST_NUMBER, random_state=STATE_NUMBER)

    model = keras.models.Sequential([
        keras.layers.Embedding(Max_words, 64, input_shape=[input_shape]),
        keras.layers.LSTM(UNIT_NUMBER),
        keras.layers.Dropout(DROPOUT_NUMBER),
        keras.layers.Dense(DENSE_NUMBER, activation="softmax")
    ])

    model.compile(loss='sparse_categorical_crossentropy', optimizer='adam', metrics=['accuracy'])

    model.fit(X_train, y_train, epochs=EPOCHS_NUMBER , batch_size=BATCH_NUMBER, validation_split=SPLIT_NUMBER)

    # Evaluates the model and then returns the model accuracy
    accuracy = evaluate_model(model, X_test, y_test, Max_words)
    return accuracy, model

def evaluate_model(model, X_test, y_test, Max_words):
    y_pred = model.predict(X_test)
    y_pred = [np.argmax(x) for x in y_pred]

    accuracy = accuracy_score(y_test, y_pred)
 
    categories = ['food', 'travel', 'lifestyle', 'other']
    # Classify an example post using the function classify_post 
    category_id = classify_post("Food is an art in Peru, the variety of dishes is varied.", model, Max_words, categories) 

    return accuracy,category_id

# Function to classify a post and return its category
def classify_post(post, lstm_model, max_sequence_length,categories):
    
    preprocessed_data, _, _ = pc_lr_preprocessing.preprocess_data(pd.DataFrame({'description': [post]}))
    post_sequence = pad_sequences(preprocessed_data[DATA_PROCESSED_COLUMN], maxlen=max_sequence_length)
    category_id = lstm_model.predict(post_sequence)
    predict_cgy = categories[np.argmax(category_id)]

    return predict_cgy


