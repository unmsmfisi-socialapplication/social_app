variable "subscription_id" {
  description = "Azure subscription ID"
  type        = string
}

variable "client_id" {
  description = "Azure AD application client ID"
  type        = string
}

variable "client_secret" {
  description = "Azure AD application client secret"
  type        = string
}

variable "tenant_id" {
  description = "Azure AD tenant ID"
  type        = string
}

variable "location" {
  description = "The Azure region for the resources."
  type        = string
  default     = "Brazil South"
}

variable "resource_group_name" {
  description = "The name of the Azure resource group."
  type        = string
  default     = "example-resources"
}

variable "environment" {
  description = "The environment for the resources (e.g., production, test, dev)."
  type        = string
  default     = "dev"
}

variable "storage_account_name" {
  description = "The name of the Azure Storage Account."
  type        = string
  default     = "examplesa"
}

variable "sql_server_name" {
  description = "The name of the Azure SQL Server."
  type        = string
  default     = "example-sqlserver"
}

variable "sql_server_version" {
  description = "The version of the Azure SQL Server."
  type        = string
  default     = "12.0"
}

variable "sql_admin_login" {
  description = "The administrator login for the SQL Server."
  type        = string
  default     = "4dm1n157r470r"
}

variable "sql_admin_password" {
  description = "The administrator password for the SQL Server."
  type        = string
  default     = "4-v3ry-53cr37-p455w0rd"
}

variable "database_name" {
  description = "The name of the Azure SQL Database."
  type        = string
  default     = "acctest-db-d"
}

variable "collation" {
  description = "The collation for the SQL Database."
  type        = string
  default     = "SQL_Latin1_General_CP1_CI_AS"
}

variable "max_size_gb" {
  description = "The maximum size of the SQL Database in gigabytes."
  type        = number
  default     = 4
}

variable "read_scale" {
  description = "Enable read scale for the SQL Database."
  type        = bool
  default     = true
}

variable "sku_name" {
  description = "The SKU name for the SQL Database."
  type        = string
  default     = "S0"
}

variable "zone_redundant" {
  description = "Enable zone redundancy for the SQL Database."
  type        = bool
  default     = true
}

variable "firewall_rule_start_ip" {
  description = "Start IP address for the SQL Database firewall rule."
  type        = string
  default     = "YourStartIPAddress"
}

variable "firewall_rule_end_ip" {
  description = "End IP address for the SQL Database firewall rule."
  type        = string
  default     = "YourEndIPAddress"
}

variable "ssl_enforcement" {
  description = "Enable SSL enforcement for the SQL Database server."
  type        = bool
  default     = true
}