resource "azurerm_linux_virtual_machine" "this_node" {
  name                = "prometheus-server-node"
  resource_group_name = var.resource_group_name
  location            = var.location
  size                = var.instance_type
  admin_username      = var.ssh_username

  network_interface_ids = [
    azurerm_network_interface.this_node.id
  ]

  admin_ssh_key {
    username   = var.ssh_username
    public_key = file(var.ssh_public_key_path)
  }

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
    disk_size_gb         = 20
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-focal"
    sku       = "20_04-lts"
    version   = "latest"
  }

  tags = merge(var.tags, {
    Name = "Prometheus Server Node"
  })

  provisioner "file" {
    content     = local.prometheus_config
    destination = "/tmp/prometheus.yml"
  }

  provisioner "remote-exec" {
    script = "./provisioner/install-prometheus-service.sh"
  }
}

resource "azurerm_network_interface" "this_node" {
  name                = "prometheus-nic"
  location            = var.location
  resource_group_name = var.resource_group_name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.subnet_id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.this_node.id
  }
}

resource "azurerm_public_ip" "this_node" {
  name                = "prometheus-pip"
  location            = var.location
  resource_group_name = var.resource_group_name
  allocation_method   = "Dynamic"
  sku                 = "Basic"
}
