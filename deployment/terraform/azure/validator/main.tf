resource "azurerm_public_ip" "validator_public_ip" {
  name                = "validator-public-ip"
  location            = var.location
  resource_group_name = var.resource_group_name
  allocation_method   = "Dynamic"
  tags                = var.tags
}

resource "azurerm_network_interface" "nic" {
  name                = "nic-validator"
  location            = var.location
  resource_group_name = var.resource_group_name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.subnet_id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.validator_public_ip.id
  }
}

resource "azurerm_linux_virtual_machine" "validator_vm" {
  name                = "validator-node"
  resource_group_name = var.resource_group_name
  location            = var.location
  size                = var.instance_type
  admin_username      = var.ssh_username

  network_interface_ids = [azurerm_network_interface.nic.id]
  disable_password_authentication = true

  admin_ssh_key {
    username   = var.ssh_username
    public_key = file(var.ssh_public_key_path)
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-focal"
    sku       = "20_04-lts"
    version   = "latest"
  }

  os_disk {
    name                 = "validator-osdisk"
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
    disk_size_gb         = 80
  }

  custom_data = base64encode(templatefile("./cloud-init/validator-init.tpl", {}))

  tags = merge(var.tags, {
    Name = "Validator Node"
  })
}
