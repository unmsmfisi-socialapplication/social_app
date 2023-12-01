provider "azurerm" {
  features {}
}

data "azurerm_subscription" "primary" {
}

// Call of Group Resource
data "azurerm_resource_group" "example" {
  name = "${var.environment}-${var.resource_group_name}"
}

data "azuread_group" "api_team" {
  object_id = "${var.api_team_id}" 
}

data "azuread_group" "bi_team" {
  object_id = "${var.bi_team_id}" 
}

data "azuread_group" "ml_team" {
  object_id = "${var.ml_team_id}" 
}

# add roles for groups

# Assign roles to a especific grouo in a resource api team
resource "azurerm_role_assignment" "add_api_team" {
  count               = length(var.roles_for_api)
  scope               = data.azurerm_resource_group.example.id  
  role_definition_name = var.roles_for_api[count.index]
  principal_id        = data.azuread_group.api_team.id
}

# Assign roles to a especific grouo in a resource bi team
resource "azurerm_role_assignment" "add_bi_team" {
  count               = length(var.roles_for_bi)
  scope               = data.azurerm_resource_group.example.id  
  role_definition_name = var.roles_for_bi[count.index]
  principal_id        = data.azuread_group.bi_team.id
}

# Assign roles to a especific grouo in a resource ml team
resource "azurerm_role_assignment" "add_ml_team" {
  count               = length(var.roles_for_ml)
  scope               = data.azurerm_resource_group.example.id  
  role_definition_name = var.roles_for_ml[count.index]
  principal_id        = data.azuread_group.ml_team.id
}


resource "azurerm_role_definition" "custom_role" {
  name               = "Aditional Permissions for ml team"
  scope              = data.azurerm_resource_group.example.id
  description        = "Permisos adicionales para la creacion y ejecucion de un notebook en azure ml"
  permissions {
    actions     = [
      "Microsoft.Storage/storageAccounts/write",
      "Microsoft.Resources/deployments/write",
      "Microsoft.KeyVault/vaults/write",
      "Microsoft.OperationalInsights/workspaces/write",
      "Microsoft.Insights/Components/Write",
      "Microsoft.MachineLearningServices/workspaces/write",
      "Microsoft.Resources/deployments/validate/action",
      "Microsoft.Resources/subscriptionRegistrations/read",
      "Microsoft.MachineLearningServices/workspaces/jobs/write",
      "Microsoft.Storage/storageAccounts/blobServices/read",
      "Microsoft.Storage/storageAccounts/blobServices/write",
      "Microsoft.Storage/storageAccounts/blobServices/containers/write",
      "Microsoft.Storage/storageAccounts/blobServices/containers/delete",
      "Microsoft.Storage/storageAccounts/blobServices/containers/read",
      "Microsoft.MachineLearningServices/workspaces/metadata/snapshots/write",
      "Microsoft.MachineLearningServices/workspaces/experiments/runs/submit/action",
      "Microsoft.MachineLearningServices/workspaces/experiments/write",
      "Microsoft.MachineLearningServices/workspaces/experiments/read",
      "Microsoft.MachineLearningServices/workspaces/experiments/delete",
      "Microsoft.MachineLearningServices/workspaces/experiments/runs/read",
      "Microsoft.MachineLearningServices/workspaces/experiments/runs/write",
      "Microsoft.MachineLearningServices/workspaces/experiments/runs/delete",
      "Microsoft.MachineLearningServices/workspaces/notebooks/storage/write",
      "Microsoft.MachineLearningServices/workspaces/notebooks/samples/read",
      "Microsoft.MachineLearningServices/workspaces/notebooks/vm/write"
    ]
    not_actions = []
  }
  assignable_scopes = [
    data.azurerm_resource_group.example.id,
  ]
}
//add custom role
resource "azurerm_role_assignment" "ml_team" {
  scope              = data.azurerm_resource_group.example.id
  role_definition_id = split("|",azurerm_role_definition.custom_role.id)[0]
  principal_id       = data.azuread_group.ml_team.id
}

