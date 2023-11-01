from pyspark.ml.classification import LogisticRegression
from pyspark.ml.evaluation import BinaryClassificationEvaluator
from pyspark.ml.tuning import CrossValidator, ParamGridBuilder
from pyspark.ml.evaluation import MulticlassClassificationEvaluator

# Defines parameterized variables.
SEED = 0
TRAIN_RATIO = 0.7
TEST_RATIO = 0.3
MAX_ITER_LR = 15
NUM_FOLDS = 8
LABEL_COLUMN = 'label'
ROC_METRIC_NAME = "areaUnderROC"
F1_METRIC_NAME = "f1"
ACCURACY_METRIC_NAME = "accuracy"

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
