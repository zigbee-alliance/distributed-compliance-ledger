terraform {
  backend "azurerm" {
    # Partial configuration - these values will be provided via command line or environment variables
    # Example: terraform init -backend-config="resource_group_name=my-rg" -backend-config="storage_account_name=mystorage" -backend-config="container_name=terraform"
    
    # Required parameters (provide via command line):
    # resource_group_name  = "your-resource-group"
    # storage_account_name = "yourstorageaccount"
    # container_name       = "terraform"
    # key                  = "terraform.tfstate"
    
    # Optional: Use Azure AD authentication (recommended)
    # use_azuread_auth = true
    
    # Optional: Use access key (alternative to Azure AD)
    # access_key = "your-storage-account-access-key"
  }
}

