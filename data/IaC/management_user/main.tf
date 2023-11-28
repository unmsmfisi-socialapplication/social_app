provider "azurerm" {
  features {}
}

// Call of Group Resource
data "azurerm_resource_group" "example" {
  name = "${var.environment}-${var.resource_group_name}"
}

// Definition of users for Data Team
variable "users" {
  type = map(object({
    user_principal_name = string
    display_name        = string
    mail_nickname = string
    password        = string
  }))
  default = {
    user1 = {
		user_principal_name = "jorge.marin@eliangomez13gmail.onmicrosoft.com"
		display_name        = "Jorge Luis Marin Evangelista"
		mail_nickname = "jorge.marin"
		password = "Qaxu403135"
    }
    user2 = {
		user_principal_name = "patrick.pisana@eliangomez13gmail.onmicrosoft.com"
		display_name        = "Patrick Florian Pisaña Llamocca"
		mail_nickname = "patrick.pisana"
		password = "Bofa476481"
    }
    user3 = {
		user_principal_name = "orlando.gil@eliangomez13gmail.onmicrosoft.com"
		display_name        = "Jesus Orlando Gil Jauregui"
		mail_nickname = "orlando.gil"
		password = "Fogu658351"
    }
    user4 = {
		user_principal_name = "gustavo.guerreros@eliangomez13gmail.onmicrosoft.com"
		display_name        = "Gustavo Raul Guerreros Jacobe"
		mail_nickname = "gustavo.guerreros"
		password = "Ruwa605372"
    }
    user5 = {
		user_principal_name = "luis.quispe@eliangomez13gmail.onmicrosoft.com"
		display_name        = "Luis Quispe Inquil"
		mail_nickname = "luis.quispe"
		password = "Yoba747305"
    }
    user6 = {
		user_principal_name = "gregory.sanchez@eliangomez13gmail.onmicrosoft.com"
		display_name        = "Gregory Sanchez Rios"
		mail_nickname = "gregory.sanchez"
		password = "Bupa863383"
    }
    user7 = {
		user_principal_name = "vladimir.alvarado@eliangomez13gmail.onmicrosoft.com"
		display_name        = "Vladimir Frank Felix Alvarado Pardo"
		mail_nickname = "vladimir.alvarado"
		password = "Caja212978"
    }
    user8 = {
		user_principal_name = "nel.tocas@eliangomez13gmail.onmicrosoft.com"
		display_name        = "Nel Jesus Tocas Chipana"
		mail_nickname = "nel.tocas"
		password = "Caro254097"
    }
    user9 = {
		user_principal_name = "ademir.vasquez@eliangomez13gmail.onmicrosoft.com"
		display_name        = "Ademir Raúl Vásquez Marcani"
		mail_nickname = "ademir.vasquez"
		password = "Huzo252380"
    }
    user10 = {
		user_principal_name = "eliangomez13@eliangomez13gmail.onmicrosoft.com"
		display_name        = "Elian Dalmiro Gomez Huanca"
		mail_nickname = "eliangomez13"
		password = "Cuwa221857"
    }
    user11 = {
		user_principal_name = "alonso.zenobio@eliangomez13gmail.onmicrosoft.com"
		display_name        = "Edgar Alonso Zenobio Pariasca"
		mail_nickname = "alonso.zenobio"
		password = "Fowa718001"
    }
  }
}

// Creation of users
resource "azuread_user" "example_user" {
	for_each = var.users

	user_principal_name = each.value.user_principal_name
	display_name        = each.value.display_name
	mail_nickname = each.value.mail_nickname
	password = each.value.password
}

// Creation of Groups
resource "azuread_group" "group1" {
  display_name        = "api-team"
  description = "API Team"
  security_enabled = true
}

resource "azuread_group" "group2" {
  display_name        = "bi-team"
  description = "BI Team"
  security_enabled = true
}

resource "azuread_group" "group3" {
  display_name        = "ml-team"
  description = "ML Team"
  security_enabled = true
}

// Member Assignments to Groups
resource "azuread_group_member" "group1_members_user7" {
  for_each = { for idx, user in var.users : idx => user if idx == "user7" }

  group_object_id   = azuread_group.group1.id
  member_object_id  = azuread_user.example_user[each.key].id
}

resource "azuread_group_member" "group1_members_user8" {
  for_each = { for idx, user in var.users : idx => user if idx == "user8" }

  group_object_id   = azuread_group.group1.id
  member_object_id  = azuread_user.example_user[each.key].id
}

resource "azuread_group_member" "group1_members_user11" {
  for_each = { for idx, user in var.users : idx => user if idx == "user11" }

  group_object_id   = azuread_group.group1.id
  member_object_id  = azuread_user.example_user[each.key].id
}

resource "azuread_group_member" "group2_members_user2" {
  for_each = { for idx, user in var.users : idx => user if idx == "user2" }

  group_object_id   = azuread_group.group2.id
  member_object_id  = azuread_user.example_user[each.key].id
}

resource "azuread_group_member" "group2_members_user5" {
  for_each = { for idx, user in var.users : idx => user if idx == "user5" }

  group_object_id   = azuread_group.group2.id
  member_object_id  = azuread_user.example_user[each.key].id
}

resource "azuread_group_member" "group2_members_user6" {
  for_each = { for idx, user in var.users : idx => user if idx == "user6" }

  group_object_id   = azuread_group.group2.id
  member_object_id  = azuread_user.example_user[each.key].id
}

resource "azuread_group_member" "group2_members_user9" {
  for_each = { for idx, user in var.users : idx => user if idx == "user9" }

  group_object_id   = azuread_group.group2.id
  member_object_id  = azuread_user.example_user[each.key].id
}

resource "azuread_group_member" "group2_members_user10" {
  for_each = { for idx, user in var.users : idx => user if idx == "user10" }

  group_object_id   = azuread_group.group2.id
  member_object_id  = azuread_user.example_user[each.key].id
}

resource "azuread_group_member" "group3_members_user1" {
  for_each = { for idx, user in var.users : idx => user if idx == "user1" }

  group_object_id   = azuread_group.group3.id
  member_object_id  = azuread_user.example_user[each.key].id
}

resource "azuread_group_member" "group3_members_user3" {
  for_each = { for idx, user in var.users : idx => user if idx == "user3" }

  group_object_id   = azuread_group.group3.id
  member_object_id  = azuread_user.example_user[each.key].id
}

resource "azuread_group_member" "group3_members_user4" {
  for_each = { for idx, user in var.users : idx => user if idx == "user4" }

  group_object_id   = azuread_group.group3.id
  member_object_id  = azuread_user.example_user[each.key].id
}