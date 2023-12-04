provider "azurerm" {
  features {}
}

data "azurerm_resource_group" "example" {
  name = "${var.environment}-${var.resource_group_name}"
}

resource "azurerm_sql_server" "example" {
  name                         = "socialdb-sv"
  resource_group_name          = data.azurerm_resource_group.example.name
  location                     = data.azurerm_resource_group.example.location
  version                      = "12.0"
  administrator_login          = "useradmin"
  administrator_login_password = "_k8A801_"
}

resource "azurerm_sql_database" "example" {
  name                = "socialdb"
  resource_group_name = data.azurerm_resource_group.example.name
  location            = data.azurerm_resource_group.example.location
  server_name         = azurerm_sql_server.example.name
  edition             = "Standard"
  collation           = "SQL_Latin1_General_CP1_CI_AS"
  max_size_gb         = 1
}