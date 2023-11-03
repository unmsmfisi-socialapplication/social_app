import requests
client_id = '843833457491177'
client_secret = '9c1a40785518c599d032433d617a0a52'
redirect_url = 'https://www.facebook.com/v18.0/dialog/oauth?response_type=token&display=popup&client_id=843833457491177&redirect_uri=https%3A%2F%2Fdevelopers.facebook.com%2Ftools%2Fexplorer%2Fcallback&auth_type=rerequest&scope=instagram_basic%2Cinstagram_content_publish%2Cpublic_profile'
access_url = 'https://www.facebook.com/v18.0/dialog/oauth?response_type=token&display=popup&client_id=843833457491177&redirect_uri=https%3A%2F%2Fdevelopers.facebook.com%2Ftools%2Fexplorer%2Fcallback&auth_type=rerequest&scope=instagram_basic%2Cinstagram_content_publish%2Cpublic_profile'
graph_url = 'https://graph.facebook.com/v18.0/'

def func_get_url():
    print('\n access code url',access_url)
    code = input("\n enter the url")
    code = code.rsplit('access_token=')[1]
    code = code.rsplit('&data_access_expiration')[0]
    return code

def func_get_long_lived_access_token(access_token = ''):
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
  
access_code = func_get_url()
long_lived_access_token = func_get_long_lived_access_token(access_token=access_code)



# def func_get_page_id(access_token = 'EAALZCdkl7vOkBO0AzuKvxlZB4ELjpWNNlOgNxxsrUCaN23j9dL6VPHgcuBgdokMiiLPGZB4NkGCUB0zWZBtNRVJHYVoT0JySEIOfmlZC8IQUk686cypafi1AaZBuO6dJX5oGt9CVzZCMjjyvMhLjuR1ZCWjXQ6qPC0l8nbnDwFisAR78uCZA502NKc1UG5wgZClT6sm6mjppjMX2MJlBmiPAZDZD'):
#     url = graph_url + 'me/accounts'
#     param = dict()
#     param['access_token'] = access_token
#     response = requests.get(url = url,params=param)
#     print("\n response", response)
#     response = response.json()
#     print("\n response", response)
#     page_id = ['data'][0]['id']
#     print("\n page_id",page_id)
#     return page_id

# def func_get_instagram_business_account(page_id = '',access_token = ''):
#     url = graph_url + page_id
#     param = dict()
#     param['fields'] = 'instagram_business_account'
#     param['access_token'] = access_token
#     response = requests.get(url = url,params=param)
#     print("\n response",response)
#     response = response.json()
#     print("\n response", response)
#     try:
#         instagram_account_id = response['instagram_business_account']['id']
#     except:
#         return {'error':'Instagram account not linked'}
#     return instagram_account_id
    
# def get_post_data(media_id='', access_token=''):
#     url = graph_url + media_id
#     param = dict()
#     param['fields'] = 'caption,like_count,media_url,owner,permalink'
#     param['access_token'] = access_token
#     response = requests.get(url=url, params=param)
#     response = response.json()
#     return response


# def func_get_media_id(instagram_account_id = '',access_token = ''):
#     url = graph_url + instagram_account_id +'/media'
#     param = dict()
#     param['access_token'] = access_token
#     response = requests.get(url =url,params = param)
#     response = response.json()
#     media = []
#     for i in response['data']:
#         media_data = get_post_data(media_id =i['id'],access_token=access_token)
#         media.append(media_data)
#     return media
    
# page_id =func_get_page_id(access_token=long_lived_access_token)
# insta_id = func_get_instagram_business_account(page_id=page_id,access_token=long_lived_access_token)
# post_data = func_get_media_id(instagram_account_id= insta_id,access_token=long_lived_access_token)
