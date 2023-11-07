import unittest
from pyspark.sql import SparkSession
import tempfile

from sa_lr_pyspark_preprocessing import ID_COLUMN, LABEL_COLUMN, TWEET_COLUMN, pre_process

class TestPreProcess(unittest.TestCase):
    def setUp(self):
        # Create a Spark session for each test case
        self.spark = SparkSession.builder.appName("pre_process_test").getOrCreate()

    def tearDown(self):
        # Stop the Spark session after each test case
        self.spark.stop()

    def test_pre_process(self):
        # Define the input CSV data
        data = [
            (3, 1, "@user that was fucking weird"),
            # Add more test data here if needed
        ]

        # Create a temporary CSV file with the test data
        with tempfile.NamedTemporaryFile(delete=False, suffix=".csv") as temp_csv:
            temp_csv.write(b"_c0,_c1,_c2\n")
            for row in data:
                temp_csv.write(f"{row[0]},{row[1]},{row[2]}\n".encode("utf-8"))
            temp_csv_path = temp_csv.name

        # Load the CSV data into a DataFrame
        df = self.spark.read.csv(temp_csv_path, header=True)

        # Call the pre_process function
        preprocessed_df = pre_process(df)

        # Define the expected column names after preprocessing
        expected_columns = [ID_COLUMN, LABEL_COLUMN, TWEET_COLUMN, "words", "word_clean", "rawFeatures", "features"]

        # Check if all the expected columns exist in the preprocessed DataFrame
        for col_name in expected_columns:
            self.assertIn(col_name, preprocessed_df.columns, f"Column '{col_name}' is missing in the preprocessed DataFrame")

        # Check if the preprocessing has been performed correctly
        # You can add more specific checks based on your data and preprocessing logic

        # Clean up the temporary CSV file
        temp_csv.close()

if __name__ == "__main__":
    unittest.main()
