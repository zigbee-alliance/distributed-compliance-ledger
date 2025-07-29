resource "azurerm_linux_virtual_machine" "this_nodes" {
  count               = var.nodes_count
  name                = "observer-node-${count.index}"
  resource_group_name = var.resource_group_name
  location            = var.location
  size                = var.instance_type
  admin_username      = var.ssh_username

  network_interface_ids = [
    azurerm_network_interface.this_nodes[count.index].id
  ]

  admin_ssh_key {
    username   = var.ssh_username
    public_key = file(var.ssh_public_key_path)
  }

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
    disk_size_gb         = 80
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-focal"
    sku       = "20_04-lts"
    version   = "latest"
  }

  tags = merge(var.tags, {
    Name = "Observer Node [${count.index}]"
  })

  provisioner "file" {
    content     = templatefile("./provisioner/cloudwatch-config.tpl", {})
    destination = "/tmp/cloudwatch-config.json"
  }

  provisioner "remote-exec" {
    script = "./provisioner/install-cloudwatch.sh"
  }

  provisioner "remote-exec" {
    script = "./provisioner/install-ansible-deps.sh"
  }
}

resource "azurerm_network_interface" "this_nodes" {
  count               = var.nodes_count
  name                = "observer-nic-${count.index}"
  location            = var.location
  resource_group_name = var.resource_group_name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.subnet_id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.this_nodes[count.index].id
  }
}

resource "azurerm_public_ip" "this_nodes" {
  count               = var.nodes_count
  name                = "observer-pip-${count.index}"
  location            = var.location
  resource_group_name = var.resource_group_name
  allocation_method   = "Dynamic"
  sku                 = "Basic"
}
