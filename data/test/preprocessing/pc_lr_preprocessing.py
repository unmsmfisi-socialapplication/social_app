from sklearn.preprocessing import LabelEncoder
from nltk.tokenize import word_tokenize
from tensorflow import keras
import re

# Variables that parameterize the columns of the final dataframe
SEQUENCES_COLUMN ='text_sequences'
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
    Max_words = 0

    if sequences:
        Max_words = len(tokenizer.word_index) + 1

    data[SEQUENCES_COLUMN] = sequences

    return data, lb_encoder, Max_words


