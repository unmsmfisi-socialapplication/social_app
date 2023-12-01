provider "azurerm" {
  features {}
}

data "azurerm_subscription" "primary" {
}

data "azurerm_client_config" "current" {}

// Call of Group Resource
data "azurerm_resource_group" "example" {
  name = "${var.environment}-${var.resource_group_name}"
}

resource "azurerm_storage_account" "container" {
  name                     = "containerforml"
  resource_group_name      = data.azurerm_resource_group.example.name
  location                 = data.azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

# Crear el Azure Key Vault
resource "azurerm_key_vault" "gekey" {
  name                        = "keyvaultforml"
  resource_group_name         = data.azurerm_resource_group.example.name
  location                    = data.azurerm_resource_group.example.location
  sku_name                    = "standard"
  tenant_id                   = data.azurerm_client_config.current.tenant_id
  soft_delete_retention_days  = 7
  purge_protection_enabled    = false
  enable_rbac_authorization   = true
}

# Crear el Application Insights
resource "azurerm_application_insights" "ge_apl_ins" {
  name                 = "general_application_insights"
  resource_group_name  = data.azurerm_resource_group.example.name
  location             = data.azurerm_resource_group.example.location
  application_type     = "web"
}

# Recurso de Azure Machine Learning Workspace
resource "azurerm_machine_learning_workspace" "example" {
  name                    = "workspacemachinelearn"
  resource_group_name     = data.azurerm_resource_group.example.name
  location                = data.azurerm_resource_group.example.location
  storage_account_id      = azurerm_storage_account.container.id
  key_vault_id            = azurerm_key_vault.gekey.id
  application_insights_id = azurerm_application_insights.ge_apl_ins.id

    identity {
    type = "SystemAssigned"
  }
  
}

output "client_config" {
  value = data.azurerm_client_config.current
}

resource "azurerm_key_vault_access_policy" "example" {
  key_vault_id = azurerm_key_vault.gekey.id
  tenant_id    = data.azurerm_client_config.current.tenant_id
  object_id    = data.azurerm_client_config.current.client_id

  key_permissions = [
    "Create",
    "Get",
    "Delete",
    "Purge",
    "Update",
  ]
}