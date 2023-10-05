# Importar las bibliotecas necesarias
import pandas as pd
import string
import random
import numpy as np
from sklearn.feature_extraction.text import CountVectorizer
from sklearn.model_selection import train_test_split
from sklearn.naive_bayes import MultinomialNB
from sklearn.metrics import accuracy_score  
from sklearn.linear_model import LogisticRegression
from sklearn.metrics import roc_curve, roc_auc_score
import matplotlib.pyplot as plt
import pandas as pd
import re
from bs4 import BeautifulSoup
import nltk
from nltk.tokenize import word_tokenize
from nltk.corpus import stopwords
from nltk.stem import PorterStemmer
from sklearn.feature_extraction.text import TfidfVectorizer
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
