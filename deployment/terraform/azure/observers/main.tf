locals {
  location            = var.location == null ? data.azurerm_resource_group.this.location : var.location
  resource_group_name = data.azurerm_resource_group.this.name

  base_prefix = "observers"

  resource_prefix = (
    var.resource_suffix == null
    ? "${local.base_prefix}-${local.location}"
    : length(var.resource_suffix) > 0
    ? "${local.base_prefix}-${var.resource_suffix}" : local.base_prefix
  )

  p2p_port        = 26656
  rpc_port        = 26657
  prometheus_port = 26660
  rest_port       = 1317
  grpc_port       = 9090

  nlb_ports = [
    {
      name : "rest",
      port : local.rest_port,
      listen_port : 80,
      listen_port_tls : 443,
    },
    {
      name : "grpc",
      port : local.grpc_port,
      listen_port : 9090,
      listen_port_tls : 8443,
    },
    {
      name : "rpc",
      port : local.rpc_port,
      listen_port : 8080,
      listen_port_tls : 26657,
    },
  ]

  lb_ip_configuration_name  = "${local.resource_prefix}-lb-public-ip-configuration"
  nic_ip_configuration_name = "internal"

  enable_tls = var.enable_tls && var.root_domain_name != ""

  vnet_network_prefix = "10.${30 + var.location_index}"
  internal_ips_range  = "10.0.0.0/8"
  subnet_name         = "${local.resource_prefix}-subnet"

  azs        = [for zm in data.azurerm_location.this.zone_mappings : zm.logical_zone]
  azs_to_use = var.azs == null || length(var.azs) == 0 ? local.azs : [for zone in local.azs : zone if contains(var.azs, tonumber(zone))]

  node_zones = [
    for index in range(var.nodes_count) : local.azs_to_use[index % length(local.azs_to_use)]
  ]
}

data "azurerm_resource_group" "this" {
  name = var.resource_group_name
}

data "azurerm_location" "this" {
  location = local.location
}

resource "azurerm_public_ip" "node" {
  count = var.nodes_count

  name                = "${local.resource_prefix}-node-${count.index}-public-ip"
  allocation_method   = "Static"
  location            = local.location
  resource_group_name = local.resource_group_name
  sku                 = "Standard"

  tags = var.tags
}


resource "azurerm_network_interface" "this" {
  count = var.nodes_count

  name                = "${local.resource_prefix}-node-${count.index}-nic"
  location            = local.location
  resource_group_name = local.resource_group_name

  ip_configuration {
    name                          = local.nic_ip_configuration_name
    subnet_id                     = azurerm_subnet.this.id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.node[count.index].id
  }

  tags = var.tags
}

resource "azurerm_network_interface_application_security_group_association" "observers" {
  count = var.nodes_count

  network_interface_id          = azurerm_network_interface.this[count.index].id
  application_security_group_id = azurerm_application_security_group.observers.id
}

resource "azurerm_linux_virtual_machine" "this_nodes" {
  count = var.nodes_count

  name                = "${local.resource_prefix}-node-${count.index}"
  resource_group_name = local.resource_group_name
  location            = local.location
  # zone                       = local.node_zones[count.index]

  size = var.instance_size

  admin_username = var.ssh_username

  admin_ssh_key {
    username   = var.ssh_username
    public_key = file(var.ssh_public_key_path)
  }

  network_interface_ids = [
    azurerm_network_interface.this[count.index].id,
  ]

  os_disk {
    caching              = "ReadWrite" # TODO review
    storage_account_type = "StandardSSD_LRS"
    disk_size_gb         = 80
  }

  encryption_at_host_enabled = var.enable_encryption_at_host

  source_image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-focal"
    sku       = "20_04-lts"
    version   = "latest"
  }

  # to avoid changes in case new latest image released
  lifecycle {
    ignore_changes = [source_image_reference]
  }

  connection {
    type        = "ssh"
    host        = self.public_ip_address
    user        = var.ssh_username
    private_key = file(var.ssh_private_key_path)
  }

  // ansible
  provisioner "remote-exec" {
    script = "./provisioner/install-ansible-deps.sh"
  }

  tags = var.tags
}
