# App Social - Data
This folder contains all data scripts and features , all made in python with initial concept testing in jupyter notebook

## File structure

  ```
├── bussines_analytics
│   ├── data_transformation
│   └── migration
├── cloud_platorm
│   ├── GoogleDrive.py
│   ├── QuickStart.py
│   └── setting.yaml
├── notebook
│   ├── post_classification
│   ├── sentiment_analysis
│   └── spam_detector
├── postgresql_bd
│   ├── connection details
├── preprocessing
│   ├── pp_model.py
├── training
│   ├── tn_pc.py
│   ├── tn_sa.py
│   └── tn_sd.py
├── execution
│   ├── ex_model.py
 ```
### bussines_analytics
This folder is about business analytics, referring to data transformation

### cloud_platorm
All the components required to configure storage will be in the Cloud_platform folder , this includes an 
initial File for Automatic credentials in Google Cloud Platform and a script with all cloud functions required
for development

### notebook
This folder is  just for concept testing made in jupyter notebook

### postgresql_bd
This folder is for the connection to the PosgreSQL database, it includes a
Initial automatic configurations file for Azure Cloud.

### preprocessing
This folder contains the script(s) for maange the preprocessing of raw data and deliver it for the respective feature model

# training
This folder contains the script(s) for manage the training of each model.

#Execution
This folder contains the script(s) for execution and results obtaining of each model