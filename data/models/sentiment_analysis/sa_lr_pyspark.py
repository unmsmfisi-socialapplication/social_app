from pyspark.sql import SparkSession
from pyspark.sql.functions import regexp_replace
from pyspark.ml.feature import Tokenizer, StopWordsRemover
from pyspark.ml.feature import CountVectorizer
from pyspark.ml.feature import IDF
from pyspark.ml.classification import LogisticRegression
from pyspark.ml.evaluation import BinaryClassificationEvaluator
from pyspark.ml.tuning import CrossValidator, ParamGridBuilder
from pyspark.ml.evaluation import MulticlassClassificationEvaluator

# Defines parameterized variables.
fileUrl = ""
ID_COLUMN = 'id'
LABEL_COLUMN = 'label'
TWEET_COLUMN = 'tweet'
SEED = 0
TRAIN_RATIO = 0.7
TEST_RATIO = 0.3
MAX_ITER_LR = 15
NUM_FOLDS = 8
ROC_METRIC_NAME = "areaUnderROC"
F1_METRIC_NAME = "f1"
ACCURACY_METRIC_NAME = "accuracy"

def read_file(fileUrl, spark):
    df = spark.read.csv(fileUrl, sep=",", inferSchema=True, header=False)
    return df


def pre_process(df):
    df = df.withColumnRenamed('_c0', ID_COLUMN).withColumnRenamed('_c1', LABEL_COLUMN).withColumnRenamed('_c2', TWEET_COLUMN)
    
    df = df.withColumn(TWEET_COLUMN, regexp_replace(TWEET_COLUMN, '[^a-z0-9A-Z`~!@#$%&<>?., ]', ''))
    df = df.withColumn(TWEET_COLUMN, regexp_replace(TWEET_COLUMN, '[0-9`~!@#$%&<>?,\']', ''))
    df = df.withColumn(TWEET_COLUMN, regexp_replace(TWEET_COLUMN, 'http://*.*.com', ''))
    df = df.withColumn(TWEET_COLUMN, regexp_replace(TWEET_COLUMN, 'www.*.com', ''))
    df = df.withColumn(TWEET_COLUMN, regexp_replace(TWEET_COLUMN, '\.', ''))
    
    tokenizer = Tokenizer(inputCol=TWEET_COLUMN, outputCol="words")
    wordData = tokenizer.transform(df)
    
    remover = StopWordsRemover(inputCol="words", outputCol="word_clean")
    word_clean_data = remover.transform(wordData)
    
    count = CountVectorizer(inputCol="word_clean", outputCol="rawFeatures")
    model = count.fit(word_clean_data)
    
    featurizedData = model.transform(word_clean_data)
    idf = IDF(inputCol="rawFeatures", outputCol="features")
    idfModel = idf.fit(featurizedData)
    rescaledData = idfModel.transform(featurizedData)
    
    return rescaledData


def train_test_split(df):
    trainDf, testDf = df.randomSplit([TRAIN_RATIO, TEST_RATIO], SEED)
    return trainDf, testDf


def details_table(train_predictions, test_predictions):
    train_predictions.groupBy(LABEL_COLUMN, 'prediction').count().show()
    test_predictions.groupBy(LABEL_COLUMN, 'prediction').count().show()


def evaluate_model(predictions, labelCol=LABEL_COLUMN, predictionCol="prediction"):
    evaluator = BinaryClassificationEvaluator(rawPredictionCol=predictionCol, labelCol=labelCol, metricName=ROC_METRIC_NAME)
    roc = evaluator.evaluate(predictions)
    
    evaluator = MulticlassClassificationEvaluator(predictionCol=predictionCol, labelCol=labelCol, metricName=F1_METRIC_NAME)
    f1 = evaluator.evaluate(predictions)
    
    evaluator = MulticlassClassificationEvaluator(predictionCol=predictionCol, labelCol=labelCol, metricName=ACCURACY_METRIC_NAME)
    accuracy = evaluator.evaluate(predictions)
    
    return {"ROC": roc, "F1": f1, "Accuracy": accuracy}


def logistic_regression(train_data, test_data):
    lr = LogisticRegression(maxIter=MAX_ITER_LR)
    paramGrid_lr = ParamGridBuilder().build()
    
    crossval_lr = CrossValidator(estimator=lr, estimatorParamMaps=paramGrid_lr, evaluator=BinaryClassificationEvaluator(), numFolds=NUM_FOLDS)
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
    # Use the file Url variable to upload the CSV file
    df = read_file(fileUrl, spark)
    df = pre_process(df)
    train_data, test_data = train_test_split(df)
    train_summary, test_summary = logistic_regression(train_data, test_data)
    spark.stop()
