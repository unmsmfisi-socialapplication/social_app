#For Api general requests
import requests
#For reddit API
import praw
#For Csv load
import csv
#For data manipulation
import pandas as pd

'''
Test variables for API (Reccomended)

client_id = '843833457491177'
client_secret = '9c1a40785518c599d032433d617a0a52'
redirect_url = 'https://www.facebook.com/v18.0/dialog/oauth?response_type=token&display=popup&client_id=843833457491177&redirect_uri=https%3A%2F%2Fdevelopers.facebook.com%2Ftools%2Fexplorer%2Fcallback&auth_type=rerequest&scope=instagram_basic%2Cinstagram_content_publish%2Cpublic_profile'
access_url = 'https://www.facebook.com/v18.0/dialog/oauth?response_type=token&display=popup&client_id=843833457491177&redirect_uri=https%3A%2F%2Fdevelopers.facebook.com%2Ftools%2Fexplorer%2Fcallback&auth_type=rerequest&scope=instagram_basic%2Cinstagram_content_publish%2Cpublic_profile'
graph_url = 'https://graph.facebook.com/v18.0/'
'''
# Functions for API collection (Testing requirements)

def func_get_url(access_url):
    print('\n access code url',access_url)
    code = input("\n enter the url")
    code = code.rsplit('access_token=')[1]
    code = code.rsplit('&data_access_expiration')[0]
    return code

def func_get_long_lived_access_token(client_id,client_secret,access_token = '', graph_url = ''):
    url = graph_url + 'oauth/access_token'
    param = dict()
    param['grant_type'] = 'fb_exchange_token'
    param['client_id'] = client_id
    param['client_secret'] = client_secret
    param['fb_exchange_token'] = access_token
    response = requests.get(url = url,params=param)
    print("\n response",response)
    response =response.json()
    print("\n response",response)
    long_lived_access_tokken = response['access_token']
    return long_lived_access_tokken

# How to recollect Instagram data
# access_code = func_get_url()
# long_lived_access_token = func_get_long_lived_access_token(access_token=access_code)

'''
Test variables for API reddit (Reccomended)
    client_id="yi_hZzmBPMkH5QFXz5WfXQ",
    client_secret="pMo2c4ydx_pLA14SZdTTJz1Nt2K0AQ",
    user_agent="app-dataRedi by u/vladnn177",
    username="vladnn177 ",
    password="B/D#Bk?Cf/6$C@U"
'''

def func_get_reddit(client_id,client_secret,user_agent,username,password):
    reddit = praw.Reddit(
        client_id= client_id,
        client_secret= client_secret,
        user_agent= user_agent,
        username= username,
        password= password 
    )

    subreddit = reddit.subreddit("python")

    top_posts = subreddit.top(limit=10)
    new_posts = subreddit.new(limit=10)

    #post = reddit.submission(id=post.id)
    #comments = post.comments
    return top_posts,new_posts

def consume_api_and_save_csv(api_url):
   # Make a GET request to the API for information
    response = requests.get(api_url)

    # Check if the application was successful
    if response.status_code == 200:
        data = response.json()

        if data:
            # Get the field names of the first object in the JSON response
            fieldnames = data[0].keys()

            # Use the hostname as the name of the CSV file
            csv_filename = api_url.split("//")[-1].replace("/", "_") + ".csv"

            # Create a CSV file and enter the data
            with open(csv_filename, "w", newline="") as csvfile:
                writer = csv.DictWriter(csvfile, fieldnames=fieldnames)

                # Write the header 
                writer.writeheader()

                # Iterate through the data and write it to the CSV file
                for item in data:
                    writer.writerow(item)

            print(f"Los datos se han guardado en el archivo '{csv_filename}'.")
            return True
        else:
            print("No se encontraron datos en la respuesta del API.")
            return False
    else:
        print("Error al obtener datos del API. Código de estado:", response.status_code)
        return False
  
# Calling the function with the API URL 
# api_url = "https://jsonplaceholder.typicode.com/posts" # Replace with the URL of the API you want to consume
# consume_api_and_save_csv(api_url)


def create_dataframe(api_url):
    try:
        # Make the request to the API
        response = requests.get(api_url)

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

