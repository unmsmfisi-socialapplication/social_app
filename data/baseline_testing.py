from sklearn.dummy import DummyClassifier
from sklearn.metrics import classification_report
from training.sp_lr_sklearn_training import load_and_analyze_data
from sklearn.model_selection import train_test_split
import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt

#Define the dummy model
def dummy_model(X,y,strategy):
    #Divide respective data into training sets at 25% of distribution
    X_train , X_test , y_train , y_test = train_test_split(X,y,test_size=0.25,random_state=42)

    #Setting the dummy model (Baseline model)
    dummy_clf = DummyClassifier(strategy=strategy,random_state=42)

    #Fitting with training sets
    dummy_clf.fit(X_train,y_train)

    #Predicted baseline values
    y_pred = dummy_clf.predict(X_test)

    report = classification_report(y_test,y_pred,output_dict=True)
    return report

#Function for knowing data class distribution
def get_data_distribution(data):
    distribution = data.value_counts()
    print(distribution)

def visualize_report(report):
    sns.heatmap(pd.DataFrame(report).iloc[:-1, :].T, annot=True)
    plt.show()

if __name__ == "__main__":
    # Models data loading and training sets
    sp_data = load_and_analyze_data("https://drive.google.com/uc?id=153kIWdyo8JaMoaKHnQ90qigDMW1830Mg")
    sp_x = sp_data['FORMATTED_CONTENT'] #X spam detector data values
    sp_y = sp_data['CLASS'] #Y spam detector data values
    
    #Read saentiment analysis data [Local consume by now]
    sa_data = pd.read_csv("Twitter_Data_ii.csv",sep=',')
    sa_data.dropna(inplace=True)
    sa_x = sa_data['clean_text'] #X sentiment analysis data values
    sa_y = sa_data['category'] #Y sentiment analysis data values

    #Getting bot distributions
    #get_data_distribution(sp_y)
    #get_data_distribution(sa_y)

    '''
    * DummyClassifier strategies
    uniform: In case the Y data distribution is balanced (Class values not far from each other) [Represents Random Rate classifier]
    most_frequnt: In case the Y data distribution is balanced (Class values not far from each other)  [Represents Zero Rate classifier]
    '''
    #Obtaining baselines metrics 
    bl_sp_report = dummy_model(sp_x,sp_y,'uniform') #baseline metrics for spam detector
    bl_sa_report = dummy_model(sa_x,sa_y,'uniform') #baseline metrics for sentiment analysis model

    visualize_report(bl_sp_report)
    visualize_report(bl_sa_report)
    