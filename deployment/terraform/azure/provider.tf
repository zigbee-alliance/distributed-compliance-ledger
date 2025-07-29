terraform {
  required_version = ">= 1.0"

  required_providers {
    azurerm = {
      source                = "hashicorp/azurerm"
      version               = "~> 3.0"
      configuration_aliases = [azurerm.primary, azurerm.secondary]
    }
  }
}

# Azure provider configuration
# You can use one of these authentication methods:
# 1. Service Principal (recommended for automation)
# 2. Azure CLI (requires az CLI to be installed)
# 3. Managed Identity (if running on Azure)

provider "azurerm" {
  features {}
  
  # Option 1: Service Principal Authentication (uncomment and configure)
  # subscription_id = "your-subscription-id"
  # tenant_id       = "your-tenant-id"
  # client_id       = "your-client-id"
  # client_secret   = "your-client-secret"
  
  # Option 2: Azure CLI Authentication (default, requires az CLI)
  # No additional configuration needed if using az login
  
  # Option 3: Managed Identity (if running on Azure)
  # No additional configuration needed
}
