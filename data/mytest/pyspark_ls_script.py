from pyspark.sql import SparkSession
from pyspark.sql.functions import col
from pyspark.sql.types import StructType, StructField, StringType
# Configuración de transformaciones para procesar el texto
from pyspark.ml.feature import RegexTokenizer, StopWordsRemover, CountVectorizer
from pyspark.ml.classification import LogisticRegression
# Configuración de una canalización (pipeline) para aplicar las transformaciones
from pyspark.ml import Pipeline
from pyspark.ml.feature import StringIndexer

# Crear una instancia de SparkSession
spark = SparkSession.builder.appName("SocialApp").getOrCreate()

# Definición del esquema para los datos
customSchema = StructType([
    StructField("clean_text", StringType()), 
    StructField("category", StringType())])

# Nombres de los archivos CSV que contienen los datos de tweets
filename = './Twitter_Data.csv'

# Lectura de datos desde los archivos CSV
df = spark.read.format("csv").option("header", "true").schema(customSchema).load(filename)

# Eliminación de filas con valores nulos
data = df.na.drop(how='any')

# Agrupación de los datos por la columna "category" y conteo de las categorías
data.groupBy("category").count().orderBy(col("count").desc())

# Tokenización del texto utilizando una expresión regular
regexTokenizer = RegexTokenizer(inputCol="clean_text", outputCol="words", pattern="\\W")

# Definición de palabras vacías (stop words)
add_stopwords = ["http","https","amp","rt","t","c","the"]
stopwordsRemover = StopWordsRemover(inputCol="words", outputCol="filtered").setStopWords(add_stopwords)

# Creación de una representación de "bag of words" a partir de las palabras tokenizadas
countVectors = CountVectorizer(inputCol="filtered", outputCol="features", vocabSize=30000, minDF=5)

# Conversión de la columna "category" a etiquetas numéricas
label_stringIdx = StringIndexer(inputCol="category", outputCol="label")

# Definición de la canalización
pipeline = Pipeline(stages=[regexTokenizer, stopwordsRemover, countVectors, label_stringIdx])

# Ajuste de la canalización a los datos
pipelineFit = pipeline.fit(data)
dataset = pipelineFit.transform(data)

# División de los datos en conjuntos de entrenamiento y prueba
(trainingData, testData) = dataset.randomSplit([0.7, 0.3], seed=100)

# Entrenamiento de un modelo de regresión logística
lr = LogisticRegression(maxIter=20, regParam=0.3, elasticNetParam=0)
lrModel = lr.fit(trainingData)

# Realización de predicciones en el conjunto de prueba
predictions = lrModel.transform(testData)

# Mostrar las 10 primeras predicciones ordenadas por probabilidad
predictions.filter(predictions['prediction'] == 0).select("clean_text", "category", "probability", "label", "prediction")\
    .orderBy("probability", ascending=False).show(n=10, truncate=30)

# Detener la sesión de pyspark
spark.stop()