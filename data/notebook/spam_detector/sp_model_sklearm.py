import pandas as pd
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.model_selection import train_test_split
from sklearn.naive_bayes import MultinomialNB
from sklearn.metrics import accuracy_score
from sklearn.linear_model import LogisticRegression
import re
from bs4 import BeautifulSoup
import nltk
from nltk.tokenize import word_tokenize
from nltk.corpus import stopwords
import gdown
import joblib
import time
import schedule

# Función para limpiar el texto
def clean_text(text):
    # Eliminar etiquetas HTML
    text = BeautifulSoup(text, 'html.parser').get_text()
    # Eliminar caracteres especiales y números
    text = re.sub(r'[^a-zA-Z]', ' ', text)
    # Tokenización
    tokens = word_tokenize(text)
    # Convertir a minúsculas y eliminar stopwords
    stop_words = set(stopwords.words('english'))
    tokens = [word.lower() for word in tokens if word.isalpha() and word.lower() not in stop_words]
    # Unir tokens nuevamente
    cleaned_text = ' '.join(tokens)
    return cleaned_text

# Función para entrenar y evaluar el modelo
def train_and_evaluate_model():
    # Descargar los recursos de NLTK necesarios
    nltk.download('punkt')
    nltk.download('stopwords')

    # Define el enlace compartible de Google Drive
    google_drive_url = "https://drive.google.com/uc?id=153kIWdyo8JaMoaKHnQ90qigDMW1830Mg"

    # Descarga el archivo desde Google Drive
    gdown.download(google_drive_url, 'Final-Dataset.csv', quiet=False)

    # Cargar el archivo CSV
    data = pd.read_csv('Final-Dataset.csv')

    # Aplicar la función de limpieza a la columna de texto
    data['text_cleaned'] = data['FORMATTED_CONTENT'].apply(clean_text)

    # Crear un objeto TF-IDF Vectorizer con un número máximo de características
    max_features = 5000  
    tfidf_vectorizer = TfidfVectorizer(max_features=max_features)

    # Ajustar y transformar los datos de texto limpio
    tfidf_matrix = tfidf_vectorizer.fit_transform(data['text_cleaned'])

    # Obtener las características TF-IDF en un DataFrame de pandas
    tfidf_df = pd.DataFrame(tfidf_matrix.toarray(), columns=tfidf_vectorizer.get_feature_names_out())

    # Dividir los datos en características (X) y etiquetas (y)
    X = tfidf_df  # Características TF-IDF
    y = data['CLASS']  # Etiquetas (clase spam o no spam)

    # Dividir los datos en conjuntos de entrenamiento y pruebas (80% entrenamiento, 20% prueba)
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

    # Crear un clasificador Naive Bayes
    naive_bayes_classifier = MultinomialNB()

    # Entrenar el modelo en el conjunto de entrenamiento
    naive_bayes_classifier.fit(X_train, y_train)

    # Realizar predicciones en el conjunto de prueba
    y_pred_naive_bayes = naive_bayes_classifier.predict(X_test)

    # Calcular la precisión del modelo
    accuracy_naive_bayes = accuracy_score(y_test, y_pred_naive_bayes)

    # Crear un clasificador de Regresión Logística
    logistic_regression_classifier = LogisticRegression()

    # Entrenar el modelo en el conjunto de entrenamiento
    logistic_regression_classifier.fit(X_train, y_train)

    # Realizar predicciones en el conjunto de prueba
    y_pred_logistic_regression = logistic_regression_classifier.predict(X_test)

    # Calcular la precisión del modelo
    accuracy_logistic_regression = accuracy_score(y_test, y_pred_logistic_regression)

    # Guardar modelos entrenados en archivos específicos en la ruta deseada
    save_path = 'C:/Users/USUARIO/Documents/GitHub/social_app/data/notebook/spam_detector/'

    joblib.dump(naive_bayes_classifier, save_path + 'naive_bayes_model.pkl')
    joblib.dump(logistic_regression_classifier, save_path + 'logistic_regression_model.pkl')
    joblib.dump(tfidf_vectorizer, save_path + 'tfidf_vectorizer.pkl')


    # Retornar el vectorizador ajustado junto con las métricas de los modelos
    return tfidf_vectorizer, naive_bayes_classifier, logistic_regression_classifier, accuracy_naive_bayes, accuracy_logistic_regression

# Función para clasificar un comentario como spam o no spam
def classify_comment(comment, tfidf_vectorizer, naive_bayes_classifier, logistic_regression_classifier):
    # Limpia el comentario
    cleaned_comment = clean_text(comment)

    # Transforma el comentario en un vector TF-IDF
    tfidf_comment = tfidf_vectorizer.transform([cleaned_comment])

    # Clasifica el comentario usando el modelo Naive Bayes
    naive_bayes_prediction = naive_bayes_classifier.predict(tfidf_comment)

    # Clasifica el comentario usando el modelo de Regresión Logística
    logistic_regression_prediction = logistic_regression_classifier.predict(tfidf_comment)

    # Cargar modelos previamente entrenados desde la ruta deseada
    save_path = 'C:/Users/USUARIO/Documents/GitHub/social_app/data/notebook/spam_detector/'

    naive_bayes_classifier = joblib.load(save_path + 'naive_bayes_model.pkl')
    logistic_regression_classifier = joblib.load(save_path + 'logistic_regression_model.pkl')
    tfidf_vectorizer = joblib.load(save_path + 'tfidf_vectorizer.pkl')


    return {
        "Naive Bayes Prediction": "spam" if naive_bayes_prediction[0] == 1 else "not spam",
        "Logistic Regression Prediction": "spam" if logistic_regression_prediction[0] == 1 else "not spam"
    }

# Define una función para ejecutar el entrenamiento y evaluación del modelo
def run_model_training():
    print("Ejecutando entrenamiento y evaluación del modelo...")
    tfidf_vectorizer, naive_bayes_classifier, logistic_regression_classifier, accuracy_naive_bayes, accuracy_logistic_regression = train_and_evaluate_model()
    print("Precisión del modelo Naive Bayes:", accuracy_naive_bayes)
    print("Precisión del modelo de Regresión Logística:", accuracy_logistic_regression)

    # Comentario para clasificar para probar el modelo
    comment_to_classify = "Thank you for your email. I appreciate your prompt response to my inquiry."
    classification_result = classify_comment(comment_to_classify, tfidf_vectorizer, naive_bayes_classifier, logistic_regression_classifier)
    print("Clasificación del comentario:", classification_result)

if __name__ == "__main__":
    # Definir el intervalo de tiempo en segundos entre cada ejecución (24 horas)
    intervalo_segundos = 24 * 60 * 60  # 24 horas

    # Programar la tarea para que se ejecute cada 24 horas
    schedule.every(intervalo_segundos).seconds.do(run_model_training)

    # Ejecutar la tarea inicialmente para que no tengas que esperar 24 horas
    run_model_training()

    # Ejecutar el programa de forma continua
    while True:
        schedule.run_pending()
        time.sleep(1)