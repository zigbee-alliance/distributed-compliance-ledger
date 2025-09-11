resource "azurerm_network_security_group" "this" {
  name                = "${local.resource_prefix}-security-group"
  resource_group_name = local.resource_group_name
  location            = local.location

  tags = var.tags
}

resource "azurerm_application_security_group" "observers" {
  name                = "${local.resource_prefix}-appsecurity-group"
  resource_group_name = local.resource_group_name
  location            = local.location
  tags                = var.tags
}

resource "azurerm_network_security_rule" "sg_dev_inbound_ssh" {
  name                        = "AllowInboundSSH"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                      = "Allow"
  destination_application_security_group_ids = [
    azurerm_application_security_group.observers.id,
  ]
  destination_port_range = "22"
  direction              = "Inbound"
  priority               = 100
  protocol               = "Tcp"
  source_address_prefix  = "*"
  source_port_range      = "*"

  # TODO a workaround, there is some issue (likely on Azure side) that leads to infinite state drifts
  lifecycle {
    ignore_changes = [destination_application_security_group_ids]
  }
}

resource "azurerm_network_security_rule" "sg_dev_inbound_icmp" {
  name                        = "AllowInboundICMP"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                      = "Allow"
  destination_application_security_group_ids = [
    azurerm_application_security_group.observers.id,
  ]
  destination_port_range = "*"
  direction              = "Inbound"
  priority               = 101
  protocol               = "Icmp"
  source_address_prefix  = "*"
  source_port_range      = "*"

  # TODO a workaround, there is some issue (likely on Azure side) that leads to infinite state drifts
  lifecycle {
    ignore_changes = [destination_application_security_group_ids]
  }
}

resource "azurerm_network_security_rule" "sg_outbound_all" {
  name                        = "AllowOutboundAll"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                      = "Allow"
  destination_address_prefix  = "*"
  destination_port_range      = "*"
  direction                   = "Outbound"
  priority                    = 102
  protocol                    = "*"
  source_application_security_group_ids = [
    azurerm_application_security_group.observers.id,
  ]
  source_port_range = "*"

  # TODO a workaround, there is some issue (likely on Azure side) that leads to infinite state drifts
  lifecycle {
    ignore_changes = [source_application_security_group_ids]
  }
}

resource "azurerm_network_security_rule" "sg_inbound_private_p2p" {
  name                        = "AllowInboundP2PFromInternalIPs"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                      = "Allow"
  destination_application_security_group_ids = [
    azurerm_application_security_group.observers.id,
  ]
  destination_port_range = local.p2p_port
  direction              = "Inbound"
  priority               = 103
  protocol               = "Tcp"
  source_address_prefix  = local.internal_ips_range
  source_port_range      = local.p2p_port

  # TODO a workaround, there is some issue (likely on Azure side) that leads to infinite state drifts
  lifecycle {
    ignore_changes = [destination_application_security_group_ids]
  }
}

resource "azurerm_network_security_rule" "sg_inbound_private_rpc" {
  name                        = "AllowInboundRPCFromInternalIPs"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                      = "Allow"
  destination_application_security_group_ids = [
    azurerm_application_security_group.observers.id,
  ]
  destination_port_range = local.rpc_port
  direction              = "Inbound"
  priority               = 104
  protocol               = "Tcp"
  source_address_prefix  = local.internal_ips_range
  source_port_range      = local.rpc_port

  # TODO a workaround, there is some issue (likely on Azure side) that leads to infinite state drifts
  lifecycle {
    ignore_changes = [destination_application_security_group_ids]
  }
}

resource "azurerm_network_security_rule" "sg_inbound_private_grpc" {
  name                        = "AllowInboundGRPCFromInternalIPs"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                      = "Allow"
  destination_application_security_group_ids = [
    azurerm_application_security_group.observers.id,
  ]
  destination_port_range = local.grpc_port
  direction              = "Inbound"
  priority               = 105
  protocol               = "Tcp"
  source_address_prefix  = local.internal_ips_range
  source_port_range      = local.grpc_port

  # TODO a workaround, there is some issue (likely on Azure side) that leads to infinite state drifts
  lifecycle {
    ignore_changes = [destination_application_security_group_ids]
  }
}

resource "azurerm_network_security_rule" "sg_inbound_private_rest" {
  name                        = "AllowInboundRestFromInternalIPs"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                      = "Allow"
  destination_application_security_group_ids = [
    azurerm_application_security_group.observers.id,
  ]
  destination_port_range = local.rest_port
  direction              = "Inbound"
  priority               = 106
  protocol               = "Tcp"
  source_address_prefix  = local.internal_ips_range
  source_port_range      = local.rest_port

  # TODO a workaround, there is some issue (likely on Azure side) that leads to infinite state drifts
  lifecycle {
    ignore_changes = [destination_application_security_group_ids]
  }
}

# TODO
resource "azurerm_network_security_rule" "sg_inbound_private_prometheus" {
  name                        = "AllowInboundPrometheusFromInternalIPs"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                      = "Allow"
  destination_application_security_group_ids = [
    azurerm_application_security_group.observers.id,
  ]
  destination_port_range = local.prometheus_port
  direction              = "Inbound"
  priority               = 107
  protocol               = "Tcp"
  source_address_prefix  = local.internal_ips_range
  source_port_range      = local.prometheus_port

  # TODO a workaround, there is some issue (likely on Azure side) that leads to infinite state drifts
  lifecycle {
    ignore_changes = [destination_application_security_group_ids]
  }
}
