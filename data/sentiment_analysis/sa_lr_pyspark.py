from pyspark.sql import SparkSession
from pyspark.sql.functions import *
from pyspark.ml.feature import Tokenizer, StopWordsRemover
from pyspark.ml.feature import CountVectorizer
from pyspark.ml.feature import IDF
from pyspark.ml.classification import LogisticRegression
from pyspark.ml.evaluation import BinaryClassificationEvaluator
from pyspark.ml.tuning import CrossValidator, ParamGridBuilder
from pyspark.ml.evaluation import MulticlassClassificationEvaluator
from pyspark.sql import SQLContext
import pyspark as ps
import warnings
from nltk.stem.snowball import SnowballStemmer

def read_file(fileUrl, spark):
    df = spark.read.csv(fileUrl, sep=",", inferSchema=True, header=False)
    return df


def pre_process(df):
    df = df.withColumnRenamed('_c0', "id").withColumnRenamed('_c1', 'label').withColumnRenamed('_c2', 'tweet')
    
    df = df.withColumn('tweet', regexp_replace('tweet', '[^a-z0-9A-Z`~!@#$%&<>?., ]', ''))
    df = df.withColumn('tweet', regexp_replace('tweet', '[0-9`~!@#$%&<>?,\']', ''))
    df = df.withColumn('tweet', regexp_replace('tweet', 'http://*.*.com', ''))
    df = df.withColumn('tweet', regexp_replace('tweet', 'www.*.com', ''))
    df = df.withColumn('tweet', regexp_replace('tweet', '\.', ''))
    
    tokenizer = Tokenizer(inputCol="tweet", outputCol="words")
    wordData = tokenizer.transform(df)
    
    remover = StopWordsRemover(inputCol="words", outputCol="word_clean")
    word_clean_data = remover.transform(wordData)
    
    stemmer = SnowballStemmer(language='english')
    stemmer_udf = udf(lambda tokens: [stemmer.stem(token) for token in tokens])
    
    count = CountVectorizer(inputCol="word_clean", outputCol="rawFeatures")
    model = count.fit(word_clean_data)
    
    featurizedData = model.transform(word_clean_data)
    idf = IDF(inputCol="rawFeatures", outputCol="features")
    idfModel = idf.fit(featurizedData)
    rescaledData = idfModel.transform(featurizedData)
    
    return rescaledData


def train_test_split(df):
    seed = 0
    trainDf, testDf = df.randomSplit([0.7, 0.3], seed)
    return trainDf, testDf


def details_table(train_predictions, test_predictions):
    train_predictions.groupBy('label', 'prediction').count().show()
    test_predictions.groupBy('label', 'prediction').count().show()


def evaluate_model(predictions, labelCol="label", predictionCol="prediction"):
    evaluator = BinaryClassificationEvaluator(rawPredictionCol=predictionCol, labelCol=labelCol, metricName="areaUnderROC")
    roc = evaluator.evaluate(predictions)
    
    evaluator = MulticlassClassificationEvaluator(predictionCol=predictionCol, labelCol=labelCol, metricName="f1")
    f1 = evaluator.evaluate(predictions)
    
    evaluator = MulticlassClassificationEvaluator(predictionCol=predictionCol, labelCol=labelCol, metricName="accuracy")
    accuracy = evaluator.evaluate(predictions)
    
    return {"ROC": roc, "F1": f1, "Accuracy": accuracy}


def logistic_regression(train_data, test_data):
    lr = LogisticRegression(maxIter=15)
    paramGrid_lr = ParamGridBuilder().build()
    
    crossval_lr = CrossValidator(estimator=lr, estimatorParamMaps=paramGrid_lr, evaluator=BinaryClassificationEvaluator(), numFolds=8)
    cvModel_lr = crossval_lr.fit(train_data)
    
    best_model_lr = cvModel_lr.bestModel
    train_fit_lr = best_model_lr.transform(train_data)
    train_summary = evaluate_model(train_fit_lr)
    
    predictions_lr = cvModel_lr.transform(test_data)
    test_summary = evaluate_model(predictions_lr)

    details_table(train_fit_lr, predictions_lr)
    
    return train_summary, test_summary


if __name__ == "__main__":
    spark = SparkSession.builder.appName("SocialApp").getOrCreate()
    url = ""  # Provide the path to file
    df = read_file(url, spark)
    df = pre_process(df)
    train_data, test_data = train_test_split(df)
    train_summary, test_summary = logistic_regression(train_data, test_data)
    spark.stop()
