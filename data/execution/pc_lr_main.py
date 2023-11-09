import pandas as pd 
from ..training import pc_lr_training

if __name__ == "__main__":
    # Change by the endpoint to be given by the backend equipment
    api_url = ''
    data = pd.read_csv(api_url, sep=",")

    # The training of the model is executed with the data entered.
    model = pc_lr_training.train_LSTM(data)
