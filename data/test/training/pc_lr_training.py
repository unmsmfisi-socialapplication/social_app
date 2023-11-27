import keras
from keras.preprocessing.sequence import pad_sequences
from sklearn.model_selection import train_test_split
from ..preprocessing import pc_lr_preprocessing as preprocessing

# Parameterized variable
DATA_PROCESSED_COLUMN ='text_sequences'
LABEL_ENCODER = 'label'
EPOCHS_NUMBER = 25
BATCH_NUMBER=32
SPLIT_NUMBER = 0.2
TEST_NUMBER = 0.2
STATE_NUMBER = 100
UNIT_NUMBER = 16
DROPOUT_NUMBER = 0.25
DENSE_NUMBER = 26

def train_model(data):
    # Values are taken from the data processed for model training
    preprocessed_data,_,Max_words = preprocessing.preprocess_data(data)

    X = pad_sequences(preprocessed_data[DATA_PROCESSED_COLUMN], maxlen=45)

    # The categories transformed to numbers by preprocessing are used to train the lstm model
    labels = preprocessed_data[LABEL_ENCODER]
    y = labels

    # Divide the dataset into test and training for the model
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=TEST_NUMBER, random_state=STATE_NUMBER)

    model = keras.models.Sequential([
        keras.layers.Embedding(Max_words, 64, input_length=45),
        keras.layers.LSTM(UNIT_NUMBER),
        keras.layers.Dropout(DROPOUT_NUMBER),
        keras.layers.Dense(DENSE_NUMBER, activation="softmax")
    ])

    model.compile(loss='sparse_categorical_crossentropy', optimizer='adam', metrics=['accuracy'])
    model.fit(X_train, y_train, epochs=EPOCHS_NUMBER, batch_size=BATCH_NUMBER, validation_split=SPLIT_NUMBER)

    return model
