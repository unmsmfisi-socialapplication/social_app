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