resource "azurerm_network_security_group" "this" {
  name                = "validator-private-security-group"
  resource_group_name = local.resource_group_name
  location            = local.location
}

resource "azurerm_network_security_rule" "sg-dev-inbound-ssh" {
  name                       = "AllowInboundSSH"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                     = "Allow"
  destination_address_prefix = "*"
  destination_port_range     = "22"
  direction                  = "Inbound"
  priority                   = 100
  protocol                   = "Tcp"
  source_address_prefix      = "*"
  source_port_range          = "*"
}

resource "azurerm_network_security_rule" "sg-dev-inbound-icmp" {
  name                       = "AllowInboundICMP"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                     = "Allow"
  destination_address_prefix = "*"
  destination_port_range     = "*"
  direction                  = "Inbound"
  priority                   = 101
  protocol                   = "Icmp"
  source_address_prefix      = "*"
  source_port_range          = "*"
}

resource "azurerm_network_security_rule" "sg-dev-outbound-all" {
  name                       = "AllowOutboundAll"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                     = "Allow"
  destination_address_prefix = "*"
  destination_port_range     = "*"
  direction                  = "Outbound"
  priority                   = 102
  protocol                   = "*"
  source_address_prefix      = "*"
  source_port_range          = "*"
}

resource "azurerm_network_security_rule" "sg-inbound-private-p2p" {
  name                       = "AllowInboundP2PFromInternalIPs"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                     = "Allow"
  destination_address_prefix = "*"
  destination_port_range     = local.p2p_port
  direction                  = "Inbound"
  priority                   = 103
  protocol                   = "Tcp"
  source_address_prefix      = "${local.vpc_network_prefix}.0.0/8"
  source_port_range          = local.p2p_port
}

resource "azurerm_network_security_rule" "sg-inbound-private-rpc" {
  name                       = "AllowInboundRPCFromInternalIPs"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                     = "Allow"
  destination_address_prefix = "*"
  destination_port_range     = local.rpc_port
  direction                  = "Inbound"
  priority                   = 104
  protocol                   = "Tcp"
  source_address_prefix      = "${local.vpc_network_prefix}.0.0/8"
  source_port_range          = local.rpc_port
}

# TODO
resource "azurerm_network_security_rule" "sg-inbound-private-prometheus" {
  name                       = "AllowInboundPrometheusFromInternalIPs"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                     = "Allow"
  destination_address_prefix = "*"
  destination_port_range     = local.prometheus_port
  direction                  = "Inbound"
  priority                   = 105
  protocol                   = "Tcp"
  source_address_prefix      = "${local.vpc_network_prefix}.0.0/8"
  source_port_range          = local.prometheus_port
}
