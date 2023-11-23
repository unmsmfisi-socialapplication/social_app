terraform {
  required_providers {
    google-beta = {
      source  = "hashicorp/google-beta"
      version = "5.7.0"
    }
  }
}

provider "google" {
  credentials = var.gcp_svc_key != "" ? file(var.gcp_svc_key) : null
  project = var.project_id
  region  = var.client_region
}

provider "google-beta" {
  credentials = var.gcp_svc_key != "" ? file(var.gcp_svc_key) : null
  project = var.project_id
  region  = var.client_region
}