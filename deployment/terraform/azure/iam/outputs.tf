output "identity_id" {
  description = "ID of the User Assigned Managed Identity"
  value       = azurerm_user_assigned_identity.identity.id
}

output "principal_id" {
  description = "Principal ID of the Managed Identity"
  value       = azurerm_user_assigned_identity.identity.principal_id
}