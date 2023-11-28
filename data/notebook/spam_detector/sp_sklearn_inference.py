import joblib
import json
from azureml.core.model import Model
from bs4 import BeautifulSoup
import re
from nltk.tokenize import word_tokenize
from nltk.corpus import stopwords

# Función de inicialización
def init():
    # Se ejecuta durante la inicialización del servicio de inferencia
    global naive_bayes_classifier, tfidf_vectorizer

    # Cargar el modelo Naive Bayes y el vectorizador TF-IDF
    model_path = Model.get_model_path('spam_detection_model')
    naive_bayes_classifier = joblib.load(model_path)
    tfidf_vectorizer_path = Model.get_model_path('tfidf_vectorizer')
    tfidf_vectorizer = joblib.load(tfidf_vectorizer_path)

# Función para realizar predicciones
def run(raw_data):
    # Se ejecuta durante la inferencia
    try:
        # Cargar los modelos y el vectorizador durante cada inferencia para garantizar la actualización
        model_path = Model.get_model_path('spam_detection_model')
        naive_bayes_classifier = joblib.load(model_path)

        tfidf_vectorizer_path = Model.get_model_path('tfidf_vectorizer')
        tfidf_vectorizer = joblib.load(tfidf_vectorizer_path)

        # Limpiar y procesar el texto de entrada
        data = json.loads(raw_data)
        text = data['text']
        cleaned_text = clean_text(text)

        # Transformar el texto en un vector TF-IDF
        tfidf_comment = tfidf_vectorizer.transform([cleaned_text])

        # Realizar la predicción utilizando el modelo Naive Bayes
        prediction = naive_bayes_classifier.predict(tfidf_comment)

        # Devolver la predicción
        return {"result": "spam" if prediction[0] == 1 else "not spam"}

    except Exception as e:
        error = str(e)
        return {"error": error}

# Función para limpiar el texto
def clean_text(text):
    text = BeautifulSoup(text, 'html.parser').get_text()
    text = re.sub(r'[^a-zA-Z]', ' ', text)
    tokens = word_tokenize(text)
    stop_words = set(stopwords.words('english'))
    tokens = [word.lower() for word in tokens if word.isalpha() and word.lower() not in stop_words]
    cleaned_text = ' '.join(tokens)
    return cleaned_text
