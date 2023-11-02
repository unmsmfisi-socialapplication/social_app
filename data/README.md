# App Social - Data
This folder contains all data scripts and features , all made in python with initial concept testing in jupyter notebook

## File structure

  ```
├── cloud_platorm
│   ├── GoogleDrive.py
│   └── setting.yaml
├── models_pg
│   ├── post_classification
│   ├── sentiment_analysis
│   └── spam_detector
├── pg_bd
│   ├── PostgreSQL
│   ├── test_SQLServer
├── preprocessing
│   ├── pp_model.py
├── training
│   ├── tn_pc.py
│   ├── tn_sa.py
│   └── tn_sd.py
├── execution
│   ├── ex_model.py
 ```

### Cloud_platorm
All the components required to configure storage will be in the Cloud_platform folder , this includes an 
initial File for Automatic credentials in Google Cloud Platform and a script with all cloud functions required
for development

### models (Provisional)
This folder contains all oficial features of Data development , each subfolder contains the respective
feature with an official python script of logical machine learning model. (Provisional)

### models_pg
This folder is  just for concept testing made in jupyter notebook

### pg_bd
This folder is for the connection to the database in PosgreSQL, this includes a
Initial automatic configurations file for Azure Cloud. There is a subfolder /test_SQLServer that is provisionally present and will soon be deleted

### preprocessing
This folder contains the script(s) for maange the preprocessing of raw data and deliver it for the respective feature model

# training
This folder contains the script(s) for manage the training of each model.

#Execution
This folder contains the script(s) for execution and results obtaining of each model