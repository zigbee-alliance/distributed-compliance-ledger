// Azure IAM Module: Create a User Assigned Managed Identity and assign roles
resource "azurerm_user_assigned_identity" "identity" {
  name                = var.identity_name
  resource_group_name = var.resource_group_name
  location            = var.location
}

resource "azurerm_role_assignment" "role" {
  for_each            = toset(var.roles)
  scope               = data.azurerm_subscription.primary.id
  role_definition_name = each.value
  principal_id        = azurerm_user_assigned_identity.identity.principal_id
}