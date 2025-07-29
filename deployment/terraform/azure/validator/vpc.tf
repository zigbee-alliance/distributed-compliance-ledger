module "this_vnet" {
  source  = "Azure/vnet/azurerm"
  version = "4.0.0"

  resource_group_name = azurerm_resource_group.validator_rg.name
  location            = var.location

  name            = "validator-vnet"
  address_space   = ["10.0.0.0/16"]
  subnet_prefixes = ["10.0.1.0/24"]

  dns_servers        = ["168.63.129.16"] # Azure default DNS
  enable_nat_gateway = true

  tags = var.tags
}
