variable "location" {
  description = "The Azure region for resources."
  type        = string
  default     = "East US"
}

variable "resource_group_name" {
  description = "The name of the Azure resource group."
  type        = string
  default     = "SocialData"
}

variable "environment" {
  description = "The environment for this resource."
  type        = string
  default     = "RG"
}
# Lista de grupos
variable "bi_team_id" {
  description = "The id of the group of bi-team."
  type        = string
  default     = "0c20751e-f711-4f38-ab28-201eed19f83c"
}

variable "api_team_id" {
  description = "The id of the group of api-team."
  type        = string
  default     = "2595ecae-3023-43bd-899c-df8a6cf0f111"
}

variable "ml_team_id" {
  description = "The id of the group of ml-team."
  type        = string
  default     = "33423782-3f11-4b0e-8b37-f92cd0b3c4e9"
}

# Lista de roles que deseas asignar
variable "roles_for_bi" {
  type    = list(string)
  default = ["Data Factory Contributor", "Storage Blob Data Contributor", "SQL Server Contributor"]  
}

variable "roles_for_api" {
  type    = list(string)
  default = ["Data Factory Contributor", "API Management Service Contributor", "Storage Blob Data Contributor","SQL Server Contributor","AzureML Compute Operator",]  
}

variable "roles_for_ml" {
  type    = list(string)
  default = ["AzureML Compute Operator", "Storage Blob Data Contributor", "Reader"]  
}

