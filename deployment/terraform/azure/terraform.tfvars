azure_resource_group_name = "dcl-resource-group"
azure_locations = ["East US", "West Europe"]

ssh_public_key_path  = "~/.ssh/id_rsa.pub"
ssh_private_key_path = "~/.ssh/id_rsa"

validator_config = {
  instance_type = "Standard_B2s"
  is_genesis    = true
}

# set to `true` before the destruction procedure
disable_validator_protection = false

private_sentries_config = {
  enable        = true
  nodes_count   = 2
  instance_type = "Standard_B2s"
}

public_sentries_config = {
  enable        = true
  enable_ipv6   = false
  nodes_count   = 2
  instance_type = "Standard_B2s"

  locations = ["East US", "West Europe"]
}

observers_config = {
  enable           = true
  nodes_count      = 3
  instance_type    = "Standard_B2s"
  root_domain_name = "matterprotocol.com"
  enable_tls       = true

  locations = ["East US", "West Europe"]
}

prometheus_config = {
  enable        = true
  instance_type = "Standard_B1s"
}