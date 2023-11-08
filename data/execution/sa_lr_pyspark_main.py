from pyspark.sql import SparkSession
from preprocessing.sa_lr_pyspark_preprocessing import preprocess_data
from training.sa_lr_pyspark_training import train_model

if __name__ == "__main__":
    spark = SparkSession.builder.appName("SocialApp").getOrCreate()

    filename = 'https://docs.google.com/spreadsheets/d/e/2PACX-1vQrsbFzUZtypCv80I7lGN4qs1m56Qss5X54FzTH-gb0lx569sjkRKCtSRemMhF1tca38rVu-mQFhbez/pubhtml?gid=817597830&single=true'

    dataset = preprocess_data(spark, filename)

    lrModel, testData = train_model(dataset)

    predictions = lrModel.transform(testData)

    predictions.filter(predictions['prediction'] == 0).select("clean_text", "category", "probability", "label", "prediction")\
        .orderBy("probability", ascending=False).show(n=10, truncate=30)

    predictions.select("clean_text", "category", "probability", "label", "prediction")\
        .toPandas().to_csv('all_predicciones.csv', index=False)

    spark.stop()
