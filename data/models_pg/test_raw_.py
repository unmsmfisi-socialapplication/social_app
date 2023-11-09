import pytest
from raw.raw_collection import consume_api_and_save_csv, create_dataframe

API_URL_SUCCESSFUL  = "https://jsonplaceholder.typicode.com/comments"
API_URL_FAILED = "https://example.com/api"

@pytest.mark.parametrize(
    "input_url, expected",
    [
        ("https://jsonplaceholder.typicode.com/posts", True),
        ("https://jsonplaceholder.typicode.com/users", True),
        ("https://jsonplaceholder.typicode.com/photos", True)
    ]
)
def test_consume_api(input_url, expected):
    assert consume_api_and_save_csv(input_url) 

# Test for a successful application
def test_create_dataframe_successful():
    df = create_dataframe(API_URL_SUCCESSFUL)
    assert df is not None
    #Change by headers of your api's dataframe
    expected_columns = ["postId", "id", "name", "email", "body"] 
    assert list(df.columns) == expected_columns

# Test for unsuccessful application
def test_create_dataframe_failed():
    df = create_dataframe(API_URL_FAILED)
    assert df is None

if __name__ == '__main':
    pytest.main()
