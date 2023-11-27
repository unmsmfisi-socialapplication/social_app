resource "google_project_service" "artifact_registry_api" {
  service = "artifactregistry.googleapis.com"
}
resource "google_project_service" "cloud_sql_admin_api" {
  service = "sqladmin.googleapis.com"
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

resource "google_artifact_registry_repository" "backend-repo" {
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
      keep_count = 20
    }
  }
  depends_on = [
    google_project_service.artifact_registry_api
  ]
}

resource "google_sql_database_instance" "sql-instance" {
  database_version    = "POSTGRES_15"
  deletion_protection = false
  instance_type       = "CLOUD_SQL_INSTANCE"
  name                = var.database_instance_name
  region              = var.client_region
  settings {
    activation_policy           = "ALWAYS"
    availability_type           = "ZONAL"
    connector_enforcement       = "NOT_REQUIRED"
    deletion_protection_enabled = false
    disk_autoresize             = true
    disk_size                   = 10
    disk_type                   = "PD_SSD"
    edition                     = "ENTERPRISE"
    pricing_plan                = "PER_USE"
    tier                        = var.database_instance_tier
    backup_configuration {
      binary_log_enabled             = false
      enabled                        = true
      location                       = "us"
      point_in_time_recovery_enabled = true
      start_time                     = "07:00"
      transaction_log_retention_days = 7
      backup_retention_settings {
        retained_backups = 7
        retention_unit   = "COUNT"
      }
    }
    ip_configuration {
      enable_private_path_for_google_cloud_services = true
      ipv4_enabled                                  = true
      # private_network                               = data.google_compute_network.default.id
      private_network = google_compute_network.private_network.id
      require_ssl     = false
    }
    location_preference {
      zone = var.client_zone
    }
  }
  depends_on = [
    google_project_service.cloud_sql_admin_api,
    google_project_service.compute_engine_api,
    google_project_service.service_networking_api,
    google_compute_network.private_network,
    google_service_networking_connection.private_vpc_connection
  ]
}

resource "google_sql_database" "database" {
  name     = var.database-name
  instance = google_sql_database_instance.sql-instance.name
  depends_on = [
    google_sql_database_instance.sql-instance
  ]
}
