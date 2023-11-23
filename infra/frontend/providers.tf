terraform {
  required_providers {
    google-beta = {
      source  = "hashicorp/google-beta"
      version = "5.7.0"
    }
  }
}

provider "google" {
  credentials = file(var.gcp_svc_key)
  project = var.project_id
  region  = var.client_region
}

provider "google-beta" {
  credentials = file(var.gcp_svc_key)
  project = var.project_id
  region  = var.client_region
}