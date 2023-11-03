import pytest
from load_csv import consume_api_and_save_csv
 

 

@pytest.mark.parametrize(
    "input_url, expected",
    [
        ("https://jsonplaceholder.typicode.com/posts",True),
        ("https://jsonplaceholder.typicode.com/users", True),
        ("https://jsonplaceholder.typicode.com/photos", True)
       
    ]
)

def test_consume_api(input_url,expected):
    assert consume_api_and_save_csv(input_url) == True