# FIXME docs:
# - authentication
# - pre-requisites:
#   - resource group should exist (not managed by the logic)
#   - subscription resource providers should be registered (or the client should have permissions to do that)
#   - ensure you have enough quota for the selected virtual machine sizes in the required locations and request the new quotes if needed
#   - (optional) if disk encryption is on need to enable that for the subscription
#     https://learn.microsoft.com/en-us/azure/virtual-machines/linux/disks-enable-host-based-encryption-cli

provider "azurerm" {
  features {}
  resource_provider_registrations = "core"
}

resource "azurerm_resource_group" "dcl" {
  count    = var.resource_group_name == null ? length(local.locations) : 0

  name     = "${var.resource_group_name_prefix}-${local.locations[count.index]}"
  location = local.locations[count.index]
}

# Validator
module "validator" {
  source              = "./validator"

  providers = {
    azurerm = azurerm
  }

  resource_group_name = local.resource_group_names[0]
  location = local.locations[0]

  tags = local.tags

  ssh_public_key_path = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path
  instance_size        = var.validator_config.instance_size

  disable_instance_protection = local.disable_validator_protection
  enable_encryption_at_host = tobool(var.validator_config.enable_encryption_at_host) == true
  # iam_instance_profile        = module.iam.iam_instance_profile # FIXME
}

# Private Sentries
module "private_sentries" {
  count = var.private_sentries_config.enable ? 1 : 0

  source = "./private-sentries"
  providers = {
    azurerm = azurerm
  }

  resource_group_name = local.resource_group_names[0]
  location = local.locations[0]

  tags = local.tags

  nodes_count          = var.private_sentries_config.nodes_count
  instance_size        = var.private_sentries_config.instance_size
  # iam_instance_profile = module.iam.iam_instance_profile FIXME

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  peer_vnet_name = module.validator.vnet.name
  peer_vnet_resource_group_name = local.resource_group_names[0]

  depends_on = [module.validator]
}

# Public Sentries location 1
module "public_sentries_1" {
  count = (var.private_sentries_config.enable &&
    var.public_sentries_config.enable &&
  contains(var.public_sentries_config.locations, 1)) ? 1 : 0

  source = "./public-sentries"
  providers = {
    azurerm = azurerm
  }

  resource_group_name = local.resource_group_names[0]
  location = local.locations[0]

  tags = local.tags

  nodes_count          = var.public_sentries_config.nodes_count
  instance_size        = var.public_sentries_config.instance_size
  # iam_instance_profile = module.iam.iam_instance_profile # FIXME

  enable_ipv6 = var.public_sentries_config.enable_ipv6

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  location_index = 1

  peer_vnet_name = module.private_sentries[0].vnet.name
  peer_vnet_resource_group_name = local.resource_group_names[0]

  depends_on = [module.private_sentries[0]]
}

# Public Sentries location 2
module "public_sentries_2" {
  count = (var.private_sentries_config.enable &&
    var.public_sentries_config.enable &&
  contains(var.public_sentries_config.locations, 2)) ? 1 : 0

  source = "./public-sentries"
  providers = {
    azurerm = azurerm
  }

  resource_group_name = local.resource_group_names[1]
  location = local.locations[1]
  resource_suffix = local.locations[0]

  tags = local.tags

  nodes_count          = var.public_sentries_config.nodes_count
  instance_size        = var.public_sentries_config.instance_size
  # iam_instance_profile = module.iam.iam_instance_profile # FIXME

  enable_ipv6 = var.public_sentries_config.enable_ipv6

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  location_index = 2

  peer_vnet_name = module.private_sentries[0].vnet.name
  peer_vnet_resource_group_name = local.resource_group_names[0]

  depends_on = [module.private_sentries[0]]
}

#   # Observers location 1
#   module "observers_1" {
#     count = (var.private_sentries_config.enable &&
#       var.observers_config.enable &&
#     contains(var.observers_config.locations, 1)) ? 1 : 0

#     source = "./observers"
#     providers = {
#       aws      = aws.location_1
#       aws.peer = aws.location_1
#     }

#     tags = local.tags

#     nodes_count          = var.observers_config.nodes_count
#     instance_size        = var.observers_config.instance_size
#     iam_instance_profile = module.iam.iam_instance_profile

#     root_domain_name = var.observers_config.root_domain_name

#     enable_tls = var.observers_config.enable_tls

#     ssh_public_key_path  = var.ssh_public_key_path
#     ssh_private_key_path = var.ssh_private_key_path

#     location_index = 1
#     peer_vnet     = module.private_sentries[0].vnet
#   }

#   # Observers location 2
#   module "observers_2" {
#     count = (var.private_sentries_config.enable &&
#       var.observers_config.enable &&
#     contains(var.observers_config.locations, 2)) ? 1 : 0

#     source = "./observers"
#     providers = {
#       aws      = aws.location_2
#       aws.peer = aws.location_1
#     }

#     tags = local.tags

#     nodes_count          = var.observers_config.nodes_count
#     instance_size        = var.observers_config.instance_size
#     iam_instance_profile = module.iam.iam_instance_profile

#     root_domain_name = var.observers_config.root_domain_name

#     enable_tls = var.observers_config.enable_tls

#     ssh_public_key_path  = var.ssh_public_key_path
#     ssh_private_key_path = var.ssh_private_key_path

#     location_index = 2
#     peer_vnet     = module.private_sentries[0].vnet
#   }

#   module "prometheus" {
#     count = local.prometheus_enabled ? 1 : 0

#     source = "./prometheus"
#     providers = {
#       aws = aws.location_1
#     }

#     tags = local.tags

#     instance_size = var.prometheus_config.instance_size

#     endpoints = local.prometheus_endpoints

#     ssh_public_key_path  = var.ssh_public_key_path
#     ssh_private_key_path = var.ssh_private_key_path

#     vnet = module.private_sentries[0].vnet
#   }
