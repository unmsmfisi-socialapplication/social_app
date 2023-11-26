# App Social - Data
This folder contains all data scripts and features , all made in python with initial concept testing in jupyter notebook

## File structure

  ```
├── bussines_analytics
│   ├── data_transformation
│   └── migration
├── cloud_mldeployment
│   ├── sentiment_analysis
├── database_management
│   ├── connection_details
│   ├── store_procedure
├── notebook
│   ├── post_classification
│   ├── sentiment_analysis
│   └── spam_detector
├── IaC
│   ├── management_user
│   ├── management_resource
├── test
│   ├── mldeployment
│   ├── preprocessing
│   ├── training
 ```
### bussines_analytics
This folder is about business analytics, referring to data transformation

### cloud_mldeployment
All the components required to configure storage will be in the Cloud_platform folder , this includes an 
initial File for Automatic credentials in Google Cloud Platform and a script with all cloud functions required
for development

### database_management
This folder contains the connection and creation of the objects.

### notebook
This folder is  just for concept testing made in jupyter notebook

### IaC
This folder contains all the Infraestructure as Code for the Azure Service.
Inside of this folder there are 2 sub-folders :
Management_user contains all the IaC that we need to create groups,asign roles and create roles with specific permissions,and add member to tha groups.
Management_resource contains all the IaC that we need to create the principal resources that is: resource groups, machine learning workspace,data factory,etc.
Managemenet_resource also contains services that must be previously created in order to create some principal services.

### test
This folder contains the local tests