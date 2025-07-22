variable "resource_group_name" {
  description = "Azure Resource Group name"
}

variable "location" {
  description = "Azure region"
}

variable "nodes_count" {
  description = "Number of observer nodes"
  type        = number
}

variable "instance_type" {
  description = "VM size for observer nodes"
  type        = string
}

variable "admin_username" {
  description = "SSH username"
  type        = string
  default     = "ubuntu"
}

variable "ssh_public_key_path" {
  description = "Path to SSH public key"
  type        = string
}

variable "subnet_id" {
  description = "Subnet ID for observer nodes"
  type        = string
}

variable "network_security_group_id" {
  description = "NSG for observer nodes"
  type        = string
}

resource "azurerm_public_ip" "observer_public_ip" {
  count               = var.nodes_count
  name                = "observer-public-ip-${count.index}"
  location            = var.location
  resource_group_name = var.resource_group_name
  allocation_method   = "Dynamic"
}

resource "azurerm_network_interface" "observer_nic" {
  count               = var.nodes_count
  name                = "observer-nic-${count.index}"
  location            = var.location
  resource_group_name = var.resource_group_name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.subnet_id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.observer_public_ip[count.index].id
  }

  network_security_group_id = var.network_security_group_id
}

resource "azurerm_linux_virtual_machine" "observer_vm" {
  count                = var.nodes_count
  name                 = "observer-vm-${count.index}"
  resource_group_name  = var.resource_group_name
  location             = var.location
  size                 = var.instance_type
  admin_username       = var.admin_username
  network_interface_ids = [azurerm_network_interface.observer_nic[count.index].id]

  admin_ssh_key {
    username   = var.admin_username
    public_key = file(var.ssh_public_key_path)
  }

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
    disk_size_gb         = 30
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "20.04-LTS"
    version   = "latest"
  }
}