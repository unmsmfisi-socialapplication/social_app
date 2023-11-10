variable "project_id" {
  description = "Social App DEV Project ID"
  type        = string
}
variable "client_region" {
  description = "Default region to gcp services"
  type        = string
}
variable "client_zone" {
  description = "Default zone to gcp services"
  type        = string
}
variable "database_instance_name" {
  description = "Database instance name"
  type        = string
}
variable "database_instance_tier" {
  description = "Tier configuration to database instance"
  type        = string
}
variable "artifact_registry_repository_id" {
  description = "Artifact Registry repository name"
  type        = string
}
variable "private-network-name" {
  description = "Private Network name"
  type        = string
}
variable "private-ip-address-name" {
  description = "Private IP adress name"
  type        = string
}
variable "database-name" {
  description = "Database name"
  type        = string
}
