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
