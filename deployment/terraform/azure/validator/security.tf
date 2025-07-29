resource "azurerm_network_security_group" "dev_sg" {
  name                = "validator-dev-security-group"
  location            = var.location
  resource_group_name = azurerm_resource_group.validator_rg.name
  tags                = var.tags

  security_rule {
    name                       = "allow_ssh"
    priority                   = 1000
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                  = "Tcp"
    source_port_range         = "*"
    destination_port_range    = "22"
    source_address_prefix     = "*"
    destination_address_prefix = "*"
  }

  security_rule {
    name                       = "allow_icmp"
    priority                   = 1010
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                  = "Icmp"
    source_port_range         = "*"
    destination_port_range    = "*"
    source_address_prefix     = "*"
    destination_address_prefix = "*"
  }
}

resource "azurerm_network_security_group" "private_sg" {
  name                = "validator-private-security-group"
  location            = var.location
  resource_group_name = azurerm_resource_group.validator_rg.name
  tags                = var.tags

  security_rule {
    name                       = "allow_p2p_internal"
    priority                   = 1000
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                  = "Tcp"
    source_port_range         = "*"
    destination_port_range    = "26656"
    source_address_prefix     = "10.0.0.0/8"
    destination_address_prefix = "*"
  }

  security_rule {
    name                       = "allow_rpc_internal"
    priority                   = 1010
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                  = "Tcp"
    source_port_range         = "*"
    destination_port_range    = "26657"
    source_address_prefix     = "10.0.0.0/8"
    destination_address_prefix = "*"
  }

  security_rule {
    name                       = "allow_prometheus_internal"
    priority                   = 1020
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                  = "Tcp"
    source_port_range         = "*"
    destination_port_range    = "26660"
    source_address_prefix     = "10.0.0.0/8"
    destination_address_prefix = "*"
  }

  security_rule {
    name                       = "allow_all_outbound"
    priority                   = 1100
    direction                  = "Outbound"
    access                     = "Allow"
    protocol                  = "*"
    source_port_range         = "*"
    destination_port_range    = "*"
    source_address_prefix     = "*"
    destination_address_prefix = "*"
  }
}
