terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 3.0"
    }
  }
}

provider "azurerm" {
  features {}
}

# Resource Group
resource "azurerm_resource_group" "dcl" {
  name     = var.resource_group_name
  location = var.location
}

# IAM analogue (Managed Identity)
module "iam" {
  source = "./iam"
  resource_group_name = azurerm_resource_group.dcl.name
  location            = var.location
}

# Validator
module "validator" {
  source              = "./validator"
  resource_group_name = azurerm_resource_group.dcl.name
  location            = var.location
  ssh_public_key_path = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path
  instance_type        = var.validator_config.instance_type
  managed_identity_id  = module.iam.managed_identity_id
}

# Private Sentries
module "private_sentries" {
  count                 = var.private_sentries_config.enable ? 1 : 0
  source                = "./private-sentries"
  resource_group_name   = azurerm_resource_group.dcl.name
  location              = var.location
  nodes_count           = var.private_sentries_config.nodes_count
  instance_type         = var.private_sentries_config.instance_type
  managed_identity_id   = module.iam.managed_identity_id
  ssh_public_key_path   = var.ssh_public_key_path
  ssh_private_key_path  = var.ssh_private_key_path
  peer_vnet             = module.validator.vnet
}

# Public Sentries
module "public_sentries" {
  count                 = (var.private_sentries_config.enable &&
                           var.public_sentries_config.enable) ? 1 : 0
  source                = "./public-sentries"
  resource_group_name   = azurerm_resource_group.dcl.name
  location              = var.location
  nodes_count           = var.public_sentries_config.nodes_count
  instance_type         = var.public_sentries_config.instance_type
  managed_identity_id   = module.iam.managed_identity_id
  enable_ipv6           = var.public_sentries_config.enable_ipv6
  ssh_public_key_path   = var.ssh_public_key_path
  ssh_private_key_path  = var.ssh_private_key_path
  peer_vnet             = module.private_sentries[0].vnet
  region_index          = 1 # можно расширить на несколько регионов через переменные
}

# Observers
module "observers" {
  count                 = (var.private_sentries_config.enable &&
                           var.observers_config.enable) ? 1 : 0
  source                = "./observers"
  resource_group_name   = azurerm_resource_group.dcl.name
  location              = var.location
  nodes_count           = var.observers_config.nodes_count
  instance_type         = var.observers_config.instance_type
  managed_identity_id   = module.iam.managed_identity_id
  root_domain_name      = var.observers_config.root_domain_name
  enable_tls            = var.observers_config.enable_tls
  ssh_public_key_path   = var.ssh_public_key_path
  ssh_private_key_path  = var.ssh_private_key_path
  peer_vnet             = module.private_sentries[0].vnet
  region_index          = 1 # можно расширить на несколько регионов через переменные
}

# Prometheus
module "prometheus" {
  count                 = (var.private_sentries_config.enable &&
                           var.prometheus_config.enable) ? 1 : 0
  source                = "./prometheus"
  resource_group_name   = azurerm_resource_group.dcl.name
  location              = var.location
  instance_type         = var.prometheus_config.instance_type
  endpoints             = local.prometheus_endpoints
  managed_identity_id   = module.iam.managed_identity_id
  ssh_public_key_path   = var.ssh_public_key_path
  ssh_private_key_path  = var.ssh_private_key_path
  vnet                  = module.private_sentries[0].vnet
}