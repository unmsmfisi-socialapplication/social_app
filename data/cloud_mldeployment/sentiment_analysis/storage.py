import firebase_admin
from firebase_admin import credentials, storage


def initialize_firebase_app(firebase_credential_path, firebase_storage_url):
    cred = credentials.Certificate(firebase_credential_path)
    firebase_admin.initialize_app(cred, {'storageBucket': firebase_storage_url})

    firebase_bucket = storage.bucket(app=firebase_admin.get_app())

    return firebase_bucket


def upload_to_firebase_storage(firebase_bucket, output_local_file_path, firebase_storage_path):
    blob = firebase_bucket.blob(firebase_storage_path)
    blob.upload_from_filename(output_local_file_path)
