terraform {
  required_version = ">= 1.0"
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 3.0, < 4.0"
    }
  }
}

provider "azurerm" {
  # subscription_id="dC1be4S-6f f9-4805—bcOe-c07b08c9d68e "
  # client_id= "laldec5a-ad93-4cbe-93b4-27Scc73ba38a"
  # client_secret="QW8Q-.fZGV23V609LtosyOU9ZXLNtHC2g17eqcC8"
  # tenant_id="25a2595b-924c-4fe2-89a1-bb6ed5bc2ceb"
  features {}
}