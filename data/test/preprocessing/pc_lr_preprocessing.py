import string
from nltk.tokenize import word_tokenize
from nltk.corpus import stopwords
from tensorflow import keras

# Parameterized variable
CLEANED_COLUMN ='text_cleaned'
SEQUENCES_COLUMN ='text_sequences'
POST_COLUMN = 'description' # It will be modified by the post column brought from backend.

def preprocess_data(data):
    # Collects all types of punctuation and obtains common Spanish words
    punct = string.punctuation
    common_words=stopwords.words('spanish')

    data[CLEANED_COLUMN] = data[POST_COLUMN].str.lower().str.replace('[' + ''.join(punct) + ']', '', regex=True)
    data[CLEANED_COLUMN] = data[CLEANED_COLUMN].apply(word_tokenize)
    data[CLEANED_COLUMN] = data[CLEANED_COLUMN].apply(lambda x: [word for word in x if word not in common_words])

    # Text Tokenization using Keras Tokenizer
    tokenizer = keras.preprocessing.text.Tokenizer()
    tokenizer.fit_on_texts(data[CLEANED_COLUMN])
    data[SEQUENCES_COLUMN] = tokenizer.texts_to_sequences(data[CLEANED_COLUMN])

    # The input_shape to be used in the LSTM model and the maximum number of words per post are determined.
    input_shape = int(sum(data[SEQUENCES_COLUMN].apply(len)) / len(data[SEQUENCES_COLUMN]))
    Max_words = max(map(max, data[SEQUENCES_COLUMN])) + 1

    return data, input_shape, Max_words
