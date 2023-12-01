provider "azurerm" {
  features {}
}

// Creation of Group Resource
resource "azurerm_resource_group" "example" {
  name     = "${var.environment}-${var.resource_group_name}"
  location = "${var.location}"
}