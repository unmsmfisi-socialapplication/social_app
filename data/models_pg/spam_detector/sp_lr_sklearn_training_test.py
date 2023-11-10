import unittest
import sys
from training.sp_lr_sklearn_training import load_and_analyze_data, train_and_evaluate_model, classify_comment

sys.path.append('C:/Users/USUARIO/Documents/GitHub/social_app/data')

class TestTraining(unittest.TestCase):

    def test_load_and_analyze_data(self):
        # Probar la función de carga de datos y análisis
        file_url = "https://drive.google.com/uc?id=153kIWdyo8JaMoaKHnQ90qigDMW1830Mg"
        data = load_and_analyze_data(file_url)

        # Verificar si los datos se han cargado correctamente
        self.assertIsNotNone(data)
        self.assertEqual(len(data), 5572)

    def test_train_and_evaluate_model(self):
        # Probar la función de entrenamiento y evaluación del modelo
        tfidf_vectorizer, naive_bayes_classifier, logistic_regression_classifier, accuracy_naive_bayes, accuracy_logistic_regression = train_and_evaluate_model()

        # Verificar si los modelos se han ajustado correctamente
        self.assertIsNotNone(tfidf_vectorizer)
        self.assertIsNotNone(naive_bayes_classifier)
        self.assertIsNotNone(logistic_regression_classifier)
        self.assertGreater(accuracy_naive_bayes, 0)
        self.assertGreater(accuracy_logistic_regression, 0)

    def test_classify_comment(self):
        # Probar la función de clasificación de comentarios
        comment = "Thank you for your email. I appreciate your prompt response to my inquiry."
        tfidf_vectorizer, naive_bayes_classifier, logistic_regression_classifier, _, _ = train_and_evaluate_model()
        result = classify_comment(comment, tfidf_vectorizer, naive_bayes_classifier, logistic_regression_classifier)

        # Verificar si la clasificación es la esperada
        self.assertIn("Naive Bayes Prediction", result)
        self.assertIn("Logistic Regression Prediction", result)

if __name__ == '__main__':
    unittest.main()
