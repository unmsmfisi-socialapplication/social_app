import pandas as pd
import requests
import io
from sklearn.feature_extraction.text import CountVectorizer
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LogisticRegression
from sklearn.metrics import classification_report
import re
from bs4 import BeautifulSoup
import nltk
from nltk.tokenize import word_tokenize
from nltk.corpus import stopwords
import time
import schedule

# Definición de variables parametrizadas
file_url = "https://drive.google.com/uc?id=153kIWdyo8JaMoaKHnQ90qigDMW1830Mg"
TEXT_COLUMN = "FORMATTED_CONTENT"
LABEL_COLUMN = "CLASS"
max_features = 5000
test_size = 0.2
random_state = 42
INTERVAL_SECONDS = 24 * 60 * 60  # 24 horas


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

# Función para cargar y analizar el archivo desde la URL
def load_and_analyze_data(file_url):
    try:
        response = requests.get(file_url)
        response.raise_for_status()  # Verificar si la solicitud fue exitosa
        data = pd.read_csv(io.StringIO(response.content.decode('utf-8')))
        return data
    except requests.exceptions.RequestException as e:
        print(f"Error en la solicitud: {e}")
        return None

# Función para entrenar y evaluar el modelo
def train_and_evaluate_model():
    # Descargar los recursos de NLTK necesarios
    nltk.download('punkt')
    nltk.download('stopwords')

    # Cargar el archivo CSV
    data = load_and_analyze_data(file_url)

    # Aplicar la función de limpieza a la columna de texto
    data['text_cleaned'] = data['FORMATTED_CONTENT'].apply(clean_text)

    # Crear un objeto TF-IDF Vectorizer con un número máximo de características
    max_features = 5000  
    count_vectorizer = CountVectorizer(max_features=max_features)

    # Ajustar y transformar los datos de texto limpio
    count_matrix = count_vectorizer.fit_transform(data['text_cleaned'])

    # Obtener las características CountVectorizer en un DataFrame de pandas
    count_df = pd.DataFrame(count_matrix.toarray(), columns=count_vectorizer.get_feature_names_out())
    
    # Dividir los datos en características (X) y etiquetas (y)
    X = count_df  # Características TF-IDF
    y = data['CLASS']  # Etiquetas (clase spam o no spam)

    # Dividir los datos en conjuntos de entrenamiento y pruebas (80% entrenamiento, 20% prueba)
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.25, random_state=42)

    # Crear un clasificador de Regresión Logística
    logistic_regression_classifier = LogisticRegression(max_iter=1000,C=100,class_weight='balanced')

    # Entrenar el modelo en el conjunto de entrenamiento
    logistic_regression_classifier.fit(X_train, y_train)

    # Realizar predicciones en el conjunto de prueba
    y_pred_logistic_regression = logistic_regression_classifier.predict(X_test)

    # Obtaion report
    lr_report = classification_report(y_test, y_pred_logistic_regression)
    
    return  lr_report

# Define una función para ejecutar el entrenamiento y evaluación del modelo
def run_model_training():
    print("Ejecutando entrenamiento y evaluación del modelo...")
    report_lr = train_and_evaluate_model()

    print("Reporte de Regresion Logistica:",report_lr)

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