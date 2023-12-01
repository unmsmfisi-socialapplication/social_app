# preprocessing.py
import pandas as pd
from sklearn.preprocessing import LabelEncoder


def preprocess_data(file_url):
    df = pd.read_csv(file_url)

    df = df.dropna()

    label_encoder = LabelEncoder()
    df['label'] = label_encoder.fit_transform(df['category'])

    return df, label_encoder
