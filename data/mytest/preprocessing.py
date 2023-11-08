from pyspark.sql import SparkSession
from pyspark.sql import functions as F
from pyspark.sql.types import *
from pyspark.ml.feature import RegexTokenizer, StopWordsRemover, CountVectorizer
from pyspark.ml import Pipeline
from pyspark.ml.feature import StringIndexer

def preprocess_data(spark, filename1, filename2):
    customSchema = StructType([
        StructField("clean_text", StringType()), 
        StructField("category", StringType())])

    df1 = spark.read.format("csv").option("header", "true").schema(customSchema).load(filename1)
    df2 = spark.read.format("csv").option("header", "true").schema(customSchema).load(filename2)
    df = df1.union(df2)
    data = df.na.drop(how='any')

    regexTokenizer = RegexTokenizer(inputCol="clean_text", outputCol="words", pattern="\\W")
    add_stopwords = ["http", "https", "amp", "rt", "t", "c", "the"]
    stopwordsRemover = StopWordsRemover(inputCol="words", outputCol="filtered").setStopWords(add_stopwords)
    countVectors = CountVectorizer(inputCol="filtered", outputCol="features", vocabSize=30000, minDF=5)

    label_stringIdx = StringIndexer(inputCol="category", outputCol="label")
    pipeline = Pipeline(stages=[regexTokenizer, stopwordsRemover, countVectors, label_stringIdx])
    pipelineFit = pipeline.fit(data)
    dataset = pipelineFit.transform(data)

    return dataset
