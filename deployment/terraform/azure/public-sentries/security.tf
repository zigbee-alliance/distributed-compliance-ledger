resource "azurerm_network_security_group" "this" {
  name                = "${local.resource_prefix}-security-group"
  resource_group_name = local.resource_group_name
  location            = local.location
  tags                = var.tags
}

resource "azurerm_application_security_group" "sentries" {
  name                = "${local.resource_prefix}-appsecurity-group"
  resource_group_name = local.resource_group_name
  location            = local.location
  tags                = var.tags
}

resource "azurerm_application_security_group" "seeds" {
  name                = "${local.resource_prefix}-seeds-appsecurity-group"
  resource_group_name = local.resource_group_name
  location            = local.location
  tags                = var.tags
}

# TODO
# - dev ipv6 ingress from "::/0" (ssh, icmp)
# - p2p and rcp ipv6 ingress from "::/0"

resource "azurerm_network_security_rule" "sg_dev_inbound_ssh" {
  name                        = "AllowInboundSSH"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                      = "Allow"
  destination_application_security_group_ids = [
    azurerm_application_security_group.sentries.id,
    azurerm_application_security_group.seeds.id,
  ]
  destination_port_range = "22"
  direction              = "Inbound"
  priority               = 100
  protocol               = "Tcp"
  source_address_prefix  = "*"
  source_port_range      = "*"

  # TODO a workaround, there is some issue (likely on Azure side) that leads to infinite state drifts:
  # - ids repassed with upper-cases
  # - on cloud they are lowercased partly
  # - so terraform plan is always with the changes
  lifecycle {  
    ignore_changes = [ destination_application_security_group_ids ]
  }
}

resource "azurerm_network_security_rule" "sg_dev_inbound_icmp" {
  name                        = "AllowInboundICMP"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                      = "Allow"
  destination_application_security_group_ids = [
    azurerm_application_security_group.sentries.id,
    azurerm_application_security_group.seeds.id,
  ]
  destination_port_range = "*"
  direction              = "Inbound"
  priority               = 101
  protocol               = "Icmp"
  source_address_prefix  = "*"
  source_port_range      = "*"

  # TODO a workaround, there is some issue (likely on Azure side) that leads to infinite state drifts
  lifecycle {  
    ignore_changes = [ destination_application_security_group_ids ]
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
    azurerm_application_security_group.sentries.id,
    azurerm_application_security_group.seeds.id,
  ]
  source_port_range = "*"

  # TODO a workaround, there is some issue (likely on Azure side) that leads to infinite state drifts
  lifecycle {  
    ignore_changes = [ source_application_security_group_ids ]
  }
}

resource "azurerm_network_security_rule" "sg_inbound_public_p2p" {
  name                        = "AllowInboundP2PFromAll"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                      = "Allow"
  destination_application_security_group_ids = [
    azurerm_application_security_group.sentries.id,
  ]
  destination_port_range = local.p2p_port
  direction              = "Inbound"
  priority               = 103
  protocol               = "Tcp"
  source_address_prefix  = "*"
  source_port_range      = local.p2p_port

  # TODO a workaround, there is some issue (likely on Azure side) that leads to infinite state drifts
  lifecycle {  
    ignore_changes = [ destination_application_security_group_ids ]
  }
}

resource "azurerm_network_security_rule" "sg_inbound_public_rpc" {
  name                        = "AllowInboundRPCFromAll"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                      = "Allow"
  destination_application_security_group_ids = [
    azurerm_application_security_group.sentries.id,
  ]
  destination_port_range = local.rpc_port
  direction              = "Inbound"
  priority               = 104
  protocol               = "Tcp"
  source_address_prefix  = "*"
  source_port_range      = local.rpc_port

  # TODO a workaround, there is some issue (likely on Azure side) that leads to infinite state drifts
  lifecycle {  
    ignore_changes = [ destination_application_security_group_ids ]
  }
}

# TODO
resource "azurerm_network_security_rule" "sg_inbound_public_prometheus" {
  name                        = "AllowInboundPrometheusFromInternalIPs"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                      = "Allow"
  destination_application_security_group_ids = [
    azurerm_application_security_group.sentries.id,
  ]
  destination_port_range = local.prometheus_port
  direction              = "Inbound"
  priority               = 106
  protocol               = "Tcp"
  source_address_prefix  = local.internal_ips_range
  source_port_range      = local.prometheus_port

  # TODO a workaround, there is some issue (likely on Azure side) that leads to infinite state drifts
  lifecycle {  
    ignore_changes = [ destination_application_security_group_ids ]
  }
}

resource "azurerm_network_security_rule" "sg_inbound_seed_public_p2p" {
  name                        = "AllowInboundSeedP2PFromAll"
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.this.name
  access                      = "Allow"
  destination_application_security_group_ids = [
    azurerm_application_security_group.seeds.id,
  ]
  destination_port_range = local.p2p_port
  direction              = "Inbound"
  priority               = 105
  protocol               = "Tcp"
  source_address_prefix  = "*"
  source_port_range      = local.p2p_port

  # TODO a workaround, there is some issue (likely on Azure side) that leads to infinite state drifts
  lifecycle {  
    ignore_changes = [ destination_application_security_group_ids ]
  }
}
