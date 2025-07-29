# Multi‐region Azure provider configuration
provider "azurerm" {
  alias    = "region_1"
  features {}
  # NOTE: Azure provider doesn’t accept a location here;
  # location is passed into each resource/module.
}

provider "azurerm" {
  alias    = "region_2"
  features {}
}

# IAM (Azure AD / Role Assignments)
module "iam" {
  source = "./iam"
  providers = {
    azurerm = azurerm.region_1
  }
  tags = local.tags

  # e.g. 
  resource_group_name = var.azure_resource_group_name
}

# Validator VM(s)
module "validator" {
  source = "./validator"
  providers = {
    azurerm = azurerm.region_1
  }

  tags                    = local.tags
  resource_group_name     = var.azure_resource_group_name
  location                = var.azure_locations[0]
  ssh_public_key_path     = var.ssh_public_key_path
  ssh_private_key_path    = var.ssh_private_key_path
  instance_type           = var.validator_config.instance_type
  disable_protection      = local.disable_validator_protection
  identity_name           = module.iam.identity_name
}

# Private Sentries (single region)
#module "private_sentries" {
#  count = var.private_sentries_config.enable ? 1 : 0

#  source = "./private-sentries"
#  providers = {
#    azurerm = azurerm.region_1
#  }
#
#  tags                  = local.tags
#  resource_group_name   = var.azure_resource_group_name
#  location              = var.azure_locations[0]
#  nodes_count           = var.private_sentries_config.nodes_count
#  instance_type         = var.private_sentries_config.instance_type
#  subnet_id             = module.vnet_azure.private_subnet_ids[0]

#  ssh_public_key_path   = var.ssh_public_key_path
#  ssh_private_key_path  = var.ssh_private_key_path
#}

# Public Sentries – region 1
module "public_sentries_1" {
  count = var.public_sentries_config.enable && contains(var.public_sentries_config.locations, var.azure_locations[0]) ? 1 : 0

  source = "./public-sentries"
  providers = {
    azurerm = azurerm.region_1
  }

  tags                  = local.tags
  resource_group_name   = var.azure_resource_group_name
  location              = var.azure_locations[0]
  nodes_count           = var.public_sentries_config.nodes_count
  instance_type         = var.public_sentries_config.instance_type
  enable_ipv6           = var.public_sentries_config.enable_ipv6
  subnet_id             = module.vnet_azure.public_subnet_ids[0]

  ssh_public_key_path   = var.ssh_public_key_path
  ssh_private_key_path  = var.ssh_private_key_path
}

# Public Sentries – region 2
module "public_sentries_2" {
  count = var.public_sentries_config.enable && contains(var.public_sentries_config.locations, var.azure_locations[1]) ? 1 : 0

  source = "./public-sentries"
  providers = {
    azurerm = azurerm.region_2
  }

  tags                  = local.tags
  resource_group_name   = var.azure_resource_group_name
  location              = var.azure_locations[1]
  nodes_count           = var.public_sentries_config.nodes_count
  instance_type         = var.public_sentries_config.instance_type
  enable_ipv6           = var.public_sentries_config.enable_ipv6
  subnet_id             = module.vnet_azure.public_subnet_ids[1]

  ssh_public_key_path   = var.ssh_public_key_path
  ssh_private_key_path  = var.ssh_private_key_path
}

# Observers – region 1
module "observers_1" {
  count = var.observers_config.enable && contains(var.observers_config.locations, var.azure_locations[0]) ? 1 : 0

  source = "./observers"
  providers = {
    azurerm = azurerm.region_1
  }

  tags                  = local.tags
  resource_group_name   = var.azure_resource_group_name
  location              = var.azure_locations[0]
  nodes_count           = var.observers_config.nodes_count
  instance_type         = var.observers_config.instance_type
  root_domain_name      = var.observers_config.root_domain_name
  enable_tls            = var.observers_config.enable_tls
  subnet_id             = module.vnet_azure.public_subnet_ids[0]

  ssh_public_key_path   = var.ssh_public_key_path
  ssh_private_key_path  = var.ssh_private_key_path
}

# Observers – region 2
module "observers_2" {
  count = var.observers_config.enable && contains(var.observers_config.locations, var.azure_locations[1]) ? 1 : 0

  source = "./observers"
  providers = {
    azurerm = azurerm.region_2
  }

  tags                  = local.tags
  resource_group_name   = var.azure_resource_group_name
  location              = var.azure_locations[1]
  nodes_count           = var.observers_config.nodes_count
  instance_type         = var.observers_config.instance_type
  root_domain_name      = var.observers_config.root_domain_name
  enable_tls            = var.observers_config.enable_tls
  subnet_id             = module.vnet_azure.public_subnet_ids[1]

  ssh_public_key_path   = var.ssh_public_key_path
  ssh_private_key_path  = var.ssh_private_key_path
}

# Prometheus
module "prometheus" {
  count = local.prometheus_enabled ? 1 : 0

  source = "./prometheus"
  providers = {
    azurerm = azurerm.region_1
  }

  tags                  = local.tags
  resource_group_name   = var.azure_resource_group_name
  location              = var.azure_locations[0]
  instance_type         = var.prometheus_config.instance_type
  endpoints             = local.prometheus_endpoints
  ssh_public_key_path   = var.ssh_public_key_path
  ssh_private_key_path  = var.ssh_private_key_path
  subnet_id             = module.vnet_azure.public_subnet_ids[0]
}
