import azure.storage.blob

def initialize_azure_storage_client(connection_string):
    return azure.storage.blob.BlobServiceClient.from_connection_string(connection_string)

def upload_to_azure_storage(blob_service_client, local_file_path, container_name, blob_name):
    container_client = blob_service_client.get_container_client(container_name)
    blob_client = container_client.get_blob_client(blob_name)
    
    with open(local_file_path, "rb") as data:
        blob_client.upload_blob(data)
