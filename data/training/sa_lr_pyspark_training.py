from pyspark.ml.classification import LogisticRegression

def train_model(dataset):
    (trainingData, testData) = dataset.randomSplit([0.7, 0.3], seed=100)
    lr = LogisticRegression(maxIter=20, regParam=0.3, elasticNetParam=0)
    lrModel = lr.fit(trainingData)
    return lrModel, testData
