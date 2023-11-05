from pyspark.sql.functions import regexp_replace
from pyspark.ml.feature import Tokenizer, StopWordsRemover
from pyspark.ml.feature import CountVectorizer
from pyspark.ml.feature import IDF

# Defines parameterized variables.
ID_COLUMN = 'id'
LABEL_COLUMN = 'label'
TWEET_COLUMN = 'tweet'

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
