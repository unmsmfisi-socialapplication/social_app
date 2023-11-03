from load_csv import consume_api_and_save_csv
 

def test_load():
    assert consume_api_and_save_csv("https://jsonplaceholder.typicode.com/posts") == True

 