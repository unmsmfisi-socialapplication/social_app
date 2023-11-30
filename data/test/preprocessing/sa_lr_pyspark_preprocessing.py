from pyspark.sql.types import StructType, StructField, StringType
from pyspark.ml.feature import RegexTokenizer, StopWordsRemover, CountVectorizer
from pyspark.ml import Pipeline
from pyspark.ml.feature import StringIndexer
from pyspark.sql.functions import col

def preprocess_data(spark, filename):
    customSchema = StructType([
        StructField("clean_text", StringType()), 
        StructField("category", StringType())])

    df = spark.read.format("csv").option("header", "true").schema(customSchema).load(filename)
    data = df.na.drop(how='any')
    data.groupBy("category").count().orderBy(col("count").desc())

    regexTokenizer = RegexTokenizer(inputCol="clean_text", outputCol="words", pattern="\\W")
    add_stopwords = ["http", "https", "amp", "rt", "t", "c", "the"]
    stopwordsRemover = StopWordsRemover(inputCol="words", outputCol="filtered").setStopWords(add_stopwords)
    countVectors = CountVectorizer(inputCol="filtered", outputCol="features", vocabSize=30000, minDF=5)

    label_stringIdx = StringIndexer(inputCol="category", outputCol="label")
    pipeline = Pipeline(stages=[regexTokenizer, stopwordsRemover, countVectors, label_stringIdx])
    pipelineFit = pipeline.fit(data)
    dataset = pipelineFit.transform(data)

    return dataset
