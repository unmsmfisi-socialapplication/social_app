import joblib
from sklearn.feature_extraction.text import TfidfVectorizer
from bs4 import BeautifulSoup
import re
from nltk.tokenize import word_tokenize
from nltk.corpus import stopwords

# Función para limpiar el texto
def clean_text(text):
    text = BeautifulSoup(text, 'html.parser').get_text()
    text = re.sub(r'[^a-zA-Z]', ' ', text)
    tokens = word_tokenize(text)
    stop_words = set(stopwords.words('english'))
    tokens = [word.lower() for word in tokens if word.isalpha() and word.lower() not in stop_words]
    cleaned_text = ' '.join(tokens)
    return cleaned_text

# Cargar modelos previamente entrenados desde la ruta deseada
save_path = 'C:/Users/USUARIO/Documents/GitHub/social_app/data/notebook/spam_detector/'
logistic_regression_classifier = joblib.load(save_path + 'logistic_regression_model.pkl')
tfidf_vectorizer = joblib.load(save_path + 'tfidf_vectorizer.pkl')

def init():
    pass

def run(input_data):
    try:
        comment = input_data['comment']
        cleaned_comment = clean_text(comment)
        tfidf_comment = tfidf_vectorizer.transform([cleaned_comment])
        logistic_regression_prediction = logistic_regression_classifier.predict(tfidf_comment)

        return {
            "Logistic Regression Prediction": "spam" if logistic_regression_prediction[0] == 1 else "not spam"
        }
    except Exception as e:
        error = str(e)
        return error
