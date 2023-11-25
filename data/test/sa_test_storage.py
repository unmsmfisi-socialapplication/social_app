import os
import json
import unittest
from unittest.mock import MagicMock
from cloud_platform.sentiment_analysis.storage import initialize_firebase_app, upload_to_firebase_storage


class TestFirebaseFunctions(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        firebase_credentials = {
            "type": "service_account",
            "project_id": "usuarios-80f40",
            "private_key_id": "4168b3e0d04ea7d54c3d7a38dc604cdb4af63041",
            "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCbDmOVK5hrCtDp\nXHmh5YcLgbQ37vN0KXfRfA6/X5TraFT2hLQzTtcd9ORYBYxKU+V/eMNTWFw6OcSI\nSmyRk1BtLKPU5mwXiBxqzmZBy20SjiCkBp8NUvfMlfh0lboUI2v9WbtLUXD0Ohih\nl86hqjIPg5RAAkUYhmU+4ui++AqaFqKgp0Sv8G+6QgpXKrGMy7/Gt5ucrZgm3g3H\nvSBIqKDWbqT1/+4rVLsELlYKfClX3guBAcshL/nXpNMx44fq1daIsicpM0vd+Uni\nBciAv42FnYwdf1CxXxc9t4bfwbMl4CjMdRYlwJH8qly1Lwb+qqSCev7TaN3YO5u7\nGv6xcjiPAgMBAAECggEABWSlOxGGniaEqLFCO8rXfSN4SL/DcQVmOP+XJbnl3KBG\n4EIyYFx8J/sWHCepPwaXA22IYbQ/lPjgaJsqFpAoqnt0gVm4ZcMiteZS8TPpqc0O\nvpXKX3TE0LVgrHFuxPeMcIX3+UEubBg/NhOpKL9aOWbqt6U8kQNPrzXGQaov+DcG\ndvNfnW5qeyFQdiZ9qOmPIW8FmTXQzkdIbRSdIBOqPVkRuAA7dG0zkzUhYSqLeRcU\ndRKRnKd8DUA76MUULjPjaQgIH5lQ0w8VnO2U7u0MiS0ryUge0u0Ej/f1wXTJ7dRY\ntgr/5xyN+f/OgF/TNbjTwptlyUfantnd8uL/qP+D9QKBgQDPoOOPZd225g7FhOV+\nnXd4zMMKZR41z3K1xK7VWuhho6LPWJHLJUIvZvmuGB+voQegICdXPa4Pv5qIFfih\ndJX0HkicSWjpl220LryClOvA6OWqqPjYOhhHB2g39cPtRV2xVLCpbWQSwZYClMYW\n6aYRTwPphzp6RNRh2U9yYq3cuwKBgQC/Lg1t4pSOinBL5ov4/wcBoO+tqQ3k8SfE\n6EQ7dx+6HXrMdEacPU8iSCrgLaq8k64YtFlzBXgQiDN90IGB3GIuV9UOtLpx9MwV\nuh1u98ozkiNqz1Wpe2S+eCyjCkYqqGHWYNVVLgGFEIpFx21rqzja6lYwkVh1Wmjg\ng77L0RPgPQKBgQCQiRmgm3MKfYBO88Kmr4uZpkEoHVg4cV8TNgXXNF+MiNMvsCJd\n3mQCh5bMILL+GgUq7Pc1/hQLloBGt7TLAFPbCrHVSi0kX2vZAKft5yluszId6V7B\ndz+rQnbEbeukYqNkbOZGK4Z8oBKDXXnAPYqDhhPdH2wP7ZZDP3nU0OjjPQKBgQCt\n9obrzbjmd9y1hEUsP0wt1Dckq+djw3kvuTaVcdFh9v9N2iSCeVtv04GIWqW5URyj\niEC45GobiQRp4PhnpZWFnBDYf+YEWf/3aVLtLfdFtB4cTCwju4+xrgLYHTo6CNx2\nScj16JYjciiQVGciCT0EHyZsrWVz69/xJJDSQB849QKBgB1gGw6cAkNFpr4EK+3z\nBPtxLN3p9pR6+ZOn3AhKvJaC356pVQs8kadyMDptCYrIhKV69o7wnxN63lr4iKVR\nGLYfW6fnA2U+WuO89RfOBPbSipYcMwC/QgmG4YV14ms0ou/FydReG82HqToODRaE\ntB5Leqd+DS25N/we+eM3+M54\n-----END PRIVATE KEY-----\n",
            "client_email": "firebase-adminsdk-ggzny@usuarios-80f40.iam.gserviceaccount.com",
            "client_id": "112083269638755958463",
            "auth_uri": "https://accounts.google.com/o/oauth2/auth",
            "token_uri": "https://oauth2.googleapis.com/token",
            "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
            "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-ggzny"
                                    "%40usuarios-80f40.iam.gserviceaccount.com",
            "universe_domain": "googleapis.com"
        }

        with open('./mock-key-firebase.json', 'w') as f:
            json.dump(firebase_credentials, f)

        cls.firebase_credential_path = './mock-key-firebase.json'
        cls.firebase_storage_url = 'mock-usuarios-80f40.appspot.com'
        cls.local_file_path = "mock/local-file.txt"
        cls.storage_path = "mock/storage-path.txt"
        cls.firebase_bucket = MagicMock()

    def test_initialize_firebase_app(self):
        self.firebase_bucket = initialize_firebase_app(self.firebase_credential_path, self.firebase_storage_url)

        self.assertIsNotNone(self.firebase_bucket)

    def test_upload_to_firebase_storage(self):
        upload_to_firebase_storage(self.firebase_bucket, self.local_file_path, self.storage_path)

        self.firebase_bucket.blob.assert_called_once_with(self.storage_path)
        self.firebase_bucket.blob().upload_from_filename.assert_called_once_with(self.local_file_path)

    @classmethod
    def tearDownClass(cls):
        os.remove('./mock-key-firebase.json')


if __name__ == '__main__':
    unittest.main()
