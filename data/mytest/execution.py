from pyspark.sql import SparkSession
from preprocessing import preprocess_data
from training import train_model

if __name__ == "__main__":
    # Crear una instancia de SparkSession
    spark = SparkSession.builder.appName("SocialApp").getOrCreate()

    # Nombres de los archivos CSV que contienen los datos de tweets
    filename = './Twitter_Data.csv'

    # Preprocesamiento de datos
    dataset = preprocess_data(spark, filename)

    # Entrenamiento del modelo
    lrModel, testData = train_model(dataset)

    # Realización de predicciones en el conjunto de prueba
    predictions = lrModel.transform(testData)

    # Mostrar las 10 primeras predicciones ordenadas por probabilidad
    predictions.filter(predictions['prediction'] == 0).select("clean_text", "category", "probability", "label", "prediction")\
        .orderBy("probability", ascending=False).show(n=10, truncate=30)

    # Guardar todas las predicciones en un archivo CSV
    predictions.select("clean_text", "category", "probability", "label", "prediction")\
        .toPandas().to_csv('all_predicciones.csv', index=False)

    # Detener la sesión de PySpark
    spark.stop()
