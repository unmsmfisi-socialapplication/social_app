locals {
  region = "us-central1"
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "google" {
  project     = "my-project-id"
  region      = local.region
}

resource "google_sql_database_instance" "main" {
  name             = "backend"
  database_version = "POSTGRES_16"
  region           = local.region

  settings {
    # Second-generation instance tiers are based on the machine
    # type. See argument reference below.
    tier = "db-f1-micro"
    ip_configuration {
      ipv4_enabled                                  = true
      private_network                               = vpc.network_id
      enable_private_path_for_google_cloud_services = true
    }
  }
}

resource "google_cloud_run_service" "default" {
  name     = "cloudrun-srv"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  metadata {
    annotations = {
      "autoscaling.knative.dev/maxScale"        = 4
      "autoscaling.knative.dev/minScale"        = 2
      "run.googleapis.com/vpc-access-connector" = element(tolist(module.serverless_connector.connector_ids), 1)
      "run.googleapis.com/vpc-access-egress"    = "all-traffic"
    }
  }
}

module "vpc" {
  source  = "terraform-google-modules/network/google"
  version = "~> 4.0"

  project_id   = var.project_id
  network_name = "cloud-run-vpc"
  routing_mode = "GLOBAL"

  subnets = [
    {
      subnet_name           = "cloud-run-subnet"
      subnet_ip             = "10.10.0.0/28"
      subnet_region         = local.region
      subnet_private_access = "true"
      subnet_flow_logs      = "false"
      description           = "Cloud Run VPC Connector Subnet"
    }
  ]
}

module "serverless-connector" {
  source     = "terraform-google-modules/network/google//modules/vpc-serverless-connector-beta"
  project_id = <PROJECT ID>
  vpc_connectors = [{
    name            = "central-serverless"
    region          = local.region
    subnet_name     = "<SUBNET NAME>"
    host_project_id = "<HOST PROJECT ID>"
    machine_type    = "e2-standard-4"
    min_instances   = 2
    max_instances   = 3
    max_throughput  = 300
  }]
}
