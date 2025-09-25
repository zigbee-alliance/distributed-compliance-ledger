locals {
  p2p_port        = 26656
  rpc_port        = 26657
  prometheus_port = 26660

  vnet_network_prefix = "10.0"
  internal_ips_range  = "10.0.0.0/8"
  subnet_name         = "validator-subnet"

  location            = var.location == null ? data.azurerm_resource_group.this.location : var.location
  resource_group_name = data.azurerm_resource_group.this.name
}


data "azurerm_resource_group" "this" {
  name = var.resource_group_name
}


resource "azurerm_public_ip" "node" {
  name                = "validator-public-ip"
  allocation_method   = "Static"
  location            = local.location
  resource_group_name = local.resource_group_name
  sku                 = "Standard"

  tags = var.tags
}


resource "azurerm_network_interface" "this" {
  name                = "validator-node-nic"
  location            = local.location
  resource_group_name = local.resource_group_name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.this.id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.node.id
  }

  tags = var.tags
}


resource "azurerm_linux_virtual_machine" "this_node" {
  name                = "validator-node"
  resource_group_name = local.resource_group_name
  location            = local.location

  size = var.instance_size

  admin_username = var.ssh_username

  admin_ssh_key {
    username   = var.ssh_username
    public_key = file(var.ssh_public_key_path)
  }

  network_interface_ids = [
    azurerm_network_interface.this.id,
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
    sku       = "20_04-lts-gen2"
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

# TODO verify, doesn't protect against terraform destroy
resource "azurerm_management_lock" "this_node_lock" {
  count = var.disable_instance_protection ? 0 : 1

  name       = "validator-node-lock"
  lock_level = "CanNotDelete"
  scope      = azurerm_linux_virtual_machine.this_node.id
}
