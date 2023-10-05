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

# Descargar los recursos de NLTK necesarios
nltk.download('punkt')
nltk.download('stopwords')

# Define el enlace compartible de Google Drive
google_drive_url = "https://drive.google.com/uc?id=153kIWdyo8JaMoaKHnQ90qigDMW1830Mg"

# Descarga el archivo desde Google Drive
gdown.download(google_drive_url, 'Final-Dataset.csv', quiet=False)

# Cargar el archivo CSV
data = pd.read_csv('Final-Dataset.csv')

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
    # Stemming (opcional)
    # stemmer = PorterStemmer()
    # tokens = [stemmer.stem(word) for word in tokens]
    # Unir tokens nuevamente
    cleaned_text = ' '.join(tokens)
    return cleaned_text

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

print("Precisión del modelo Naive Bayes:", accuracy_naive_bayes)
print("Precisión del modelo de Regresión Logística:", accuracy_logistic_regression)
