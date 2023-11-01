from pyspark.sql import SparkSession
from sa_lr_pyspark_preprocessing import pre_process
from sa_lr_pyspark_training import train_test_split, logistic_regression

# Defines parameterized variables.
fileUrl = ""

def read_file(fileUrl, spark):
    df = spark.read.csv(fileUrl, sep=",", inferSchema=True, header=False)
    return df

if __name__ == "__main__":
    spark = SparkSession.builder.appName("SocialApp").getOrCreate()
    df = read_file(fileUrl, spark)
    df = pre_process(df)
    train_data, test_data = train_test_split(df)
    train_summary, test_summary = logistic_regression(train_data, test_data)
    spark.stop()