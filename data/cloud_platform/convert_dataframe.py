import requests as rq
import pandas as pd

def create_dataframe(api_url):
    try:
        # Make the request to the API
        response = rq.get(api_url)

        # If the request is successful, it saves the data in a dataframe.
        if response.status_code == 200:
            data = response.json()
            df = pd.DataFrame(data)
            return df
        else:
            print(f"No data was obtained from API, state: {response.status_code}")
            return None
        
    except Exception as e:
        print(f"Error: {e}")
        return None

# API call
if __name__ == "__main__":
    
    # Replace it with the URL of your Api
    api = "https://jsonplaceholder.typicode.com/comments"
    dataframe = create_dataframe(api)

    # Display the first 5 records of the dataframe
    if dataframe is not None:
        print("Data:")
        print(dataframe.head())