resource "azurerm_public_ip" "lb" {
  name                = "${local.resource_prefix}-lb-public-ip"
  location            = local.location
  resource_group_name = local.resource_group_name
  allocation_method   = "Static"
  sku                 = "Standard"
  tags                = var.tags
}

resource "azurerm_lb" "this" {
  name                = "${local.resource_prefix}-lb"
  location            = local.location
  resource_group_name = local.resource_group_name
  sku                 = "Standard"

  frontend_ip_configuration {
    name                 = local.lb_ip_configuration_name
    public_ip_address_id = azurerm_public_ip.lb.id
  }
  tags = var.tags
}

resource "azurerm_lb_backend_address_pool" "this" {
  name            = "${local.resource_prefix}-backend-address-pool"
  loadbalancer_id = azurerm_lb.this.id
}

resource "azurerm_lb_probe" "this" {
  count           = length(local.nlb_ports)
  name            = "${local.resource_prefix}-lb-probe-${local.nlb_ports[count.index].name}"
  loadbalancer_id = azurerm_lb.this.id
  protocol        = "Tcp"
  port            = local.nlb_ports[count.index].port
  probe_threshold = 2 # NOTE no way to set different health/unhealth
  # threasholds like for AWS or GCP
  interval_in_seconds = 30
}

resource "azurerm_lb_rule" "this" {
  count           = length(local.nlb_ports)
  loadbalancer_id = azurerm_lb.this.id
  name            = "${local.resource_prefix}-lb-rule-${local.nlb_ports[count.index].name}"
  protocol        = "Tcp"
  frontend_port = (
    local.enable_tls
    ? local.nlb_ports[count.index].listen_port_tls
    : local.nlb_ports[count.index].listen_port
  )
  backend_port                   = local.nlb_ports[count.index].port
  disable_outbound_snat          = true
  frontend_ip_configuration_name = local.lb_ip_configuration_name
  probe_id                       = azurerm_lb_probe.this[count.index].id
  backend_address_pool_ids       = [azurerm_lb_backend_address_pool.this.id]
}

# Associate Network Interface to the Backend Pool of the Load Balancer
resource "azurerm_network_interface_backend_address_pool_association" "this" {
  count                   = var.nodes_count
  ip_configuration_name   = local.nic_ip_configuration_name
  network_interface_id    = azurerm_network_interface.this[count.index].id
  backend_address_pool_id = azurerm_lb_backend_address_pool.this.id
}
