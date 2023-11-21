resource "google_project_service" "artifact_registry_api" {
  service = "artifactregistry.googleapis.com"
}
resource "google_project_service" "compute_engine_api" {
  service = "compute.googleapis.com"
}
resource "google_project_service" "service_networking_api" {
  service    = "servicenetworking.googleapis.com"
  depends_on = [google_project_service.compute_engine_api]
}
resource "google_project_service" "cloud_run_admin_api" {
  service = "run.googleapis.com"
}
resource "google_project_service" "service_usage_api" {
  service = "serviceusage.googleapis.com"
}

resource "google_compute_network" "private_network" {
  provider = google-beta
  project  = var.project_id
  name     = var.private-network-name
  depends_on = [
    google_project_service.compute_engine_api,
    google_project_service.service_networking_api
  ]
}

resource "google_compute_global_address" "private_ip_address" {
  provider      = google-beta
  project       = var.project_id
  name          = var.private-ip-address-name
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.private_network.id
  depends_on    = [google_project_service.compute_engine_api]
}

resource "google_service_networking_connection" "private_vpc_connection" {
  provider                = google-beta
  network                 = google_compute_network.private_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_address.name]
  depends_on              = [google_project_service.compute_engine_api]
}

resource "random_id" "db_name_suffix" {
  byte_length = 4
}

resource "google_artifact_registry_repository" "frontend-repo" {
  provider               = google-beta
  project                = var.project_id
  location               = var.client_region
  repository_id          = var.artifact_registry_repository_id
  description            = "Docker repository with cleanup policies for backend images"
  format                 = "DOCKER"
  cleanup_policy_dry_run = false
  cleanup_policies {
    id     = "keep-minimum-versions"
    action = "KEEP"
    most_recent_versions {
      keep_count = 10
    }
  }
  depends_on = [
    google_project_service.artifact_registry_api
  ]
}
