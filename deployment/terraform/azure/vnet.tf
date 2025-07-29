# VNet for Azure infrastructure
module "vnet_azure" {
  source  = "Azure/vnet/azurerm"
  version = "4.0.0"

  resource_group_name = var.azure_resource_group_name
  location            = var.azure_locations[0]

  name            = "dcl-vnet"
  address_space   = ["10.0.0.0/16"]
  subnet_prefixes = [
    "10.0.1.0/24",  # Public subnet region 1
    "10.0.2.0/24",  # Private subnet region 1
    "10.0.3.0/24",  # Public subnet region 2 (if needed)
    "10.0.4.0/24"   # Private subnet region 2 (if needed)
  ]

  subnet_names = [
    "public-subnet-1",
    "private-subnet-1", 
    "public-subnet-2",
    "private-subnet-2"
  ]

  dns_servers        = ["168.63.129.16"] # Azure default DNS
  enable_nat_gateway = true

  tags = local.tags
}

# Additional VNet for region 2 if different region
module "vnet_azure_region2" {
  count   = length(var.azure_locations) > 1 ? 1 : 0
  source  = "Azure/vnet/azurerm"
  version = "4.0.0"

  resource_group_name = var.azure_resource_group_name
  location            = var.azure_locations[1]

  name            = "dcl-vnet-region2"
  address_space   = ["10.1.0.0/16"]
  subnet_prefixes = [
    "10.1.1.0/24",  # Public subnet region 2
    "10.1.2.0/24"   # Private subnet region 2
  ]

  subnet_names = [
    "public-subnet-region2",
    "private-subnet-region2"
  ]

  dns_servers        = ["168.63.129.16"] # Azure default DNS
  enable_nat_gateway = true

  tags = local.tags
} 