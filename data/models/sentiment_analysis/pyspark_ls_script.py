from pyspark.sql import SparkSession
from pyspark.sql.functions import col
from pyspark.sql.types import StructType, StructField, StringType
from pyspark.ml.feature import RegexTokenizer, StopWordsRemover, CountVectorizer
from pyspark.ml.classification import LogisticRegression
from pyspark.ml import Pipeline
from pyspark.ml.feature import StringIndexer
#Classification report 
from pyspark.ml.evaluation import MulticlassClassificationEvaluator #Local basline testing only

spark = SparkSession.builder.appName("SocialApp").getOrCreate()

customSchema = StructType([
    StructField("clean_text", StringType()), 
    StructField("category", StringType())])

filename = 'Twitter_Data_ii - Twitter_Data_ii.csv'

df = spark.read.format("csv").option("header", "true").schema(customSchema).load(filename)

data = df.na.drop(how='any')

data.groupBy("category").count().orderBy(col("count").desc())

regexTokenizer = RegexTokenizer(inputCol="clean_text", outputCol="words", pattern="\\W")

add_stopwords = ["http","https","amp","rt","t","c","the"]
stopwordsRemover = StopWordsRemover(inputCol="words", outputCol="filtered").setStopWords(add_stopwords)

countVectors = CountVectorizer(inputCol="filtered", outputCol="features", vocabSize=30000, minDF=5)

label_stringIdx = StringIndexer(inputCol="category", outputCol="label")

pipeline = Pipeline(stages=[regexTokenizer, stopwordsRemover, countVectors, label_stringIdx])

pipelineFit = pipeline.fit(data)
dataset = pipelineFit.transform(data)

(trainingData, testData) = dataset.randomSplit([0.7, 0.3], seed=100)

lr = LogisticRegression(maxIter=20, regParam=0.3, elasticNetParam=0)
lrModel = lr.fit(trainingData)

predictions = lrModel.transform(testData)

# Classification report with pyspark
evaluator = MulticlassClassificationEvaluator(predictionCol="prediction", labelCol="label", metricName="accuracy")
accuracy = evaluator.evaluate(predictions)
print("Test Error = %g" % (1.0 - accuracy))
print("Accuracy = %g" % accuracy)

# Other metrics
precision = evaluator.evaluate(predictions, {evaluator.metricName: "weightedPrecision"})
recall = evaluator.evaluate(predictions, {evaluator.metricName: "weightedRecall"})
f1 = evaluator.evaluate(predictions, {evaluator.metricName: "f1"})

print("Precision = %g" % precision)
print("Recall = %g" % recall)
print("F1 Score = %g" % f1)

predictions.filter(predictions['prediction'] == 0).select("clean_text", "category", "probability", "label", "prediction")\
    .orderBy("probability", ascending=False).show(n=10, truncate=30)

spark.stop()