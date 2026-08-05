resource "azurerm_virtual_network" "this" {
  name                = "validator-vnet"
  location            = local.location
  resource_group_name = local.resource_group_name
  address_space       = ["${local.vnet_network_prefix}.0.0/16"]

  tags = var.tags
}

resource "azurerm_subnet" "this" {
  name                 = local.subnet_name
  resource_group_name  = local.resource_group_name
  virtual_network_name = azurerm_virtual_network.this.name
  address_prefixes     = ["${local.vnet_network_prefix}.1.0/24"]
}

resource "azurerm_subnet_network_security_group_association" "this" {
  subnet_id                 = azurerm_subnet.this.id
  network_security_group_id = azurerm_network_security_group.this.id
}

resource "azurerm_nat_gateway" "this" {
  count = var.enable_nay_gw ? 1 : 0

  name                = "validator-vnet-nat-gw"
  resource_group_name = local.resource_group_name
  location            = local.location

  tags = var.tags
}

resource "azurerm_subnet_nat_gateway_association" "this" {
  count = var.enable_nay_gw ? 1 : 0

  subnet_id      = azurerm_subnet.this.id
  nat_gateway_id = azurerm_nat_gateway.this[0].id
}

resource "azurerm_public_ip" "nat_gw" {
  count = var.enable_nay_gw ? 1 : 0

  name                = "validator-nat-gw-public-ip"
  location            = local.location
  resource_group_name = local.resource_group_name
  allocation_method   = "Static"
  sku                 = "Standard"

  tags = var.tags
}

resource "azurerm_nat_gateway_public_ip_association" "this" {
  count = var.enable_nay_gw ? 1 : 0

  nat_gateway_id       = azurerm_nat_gateway.this[0].id
  public_ip_address_id = azurerm_public_ip.nat_gw[0].id
}
