# FIXME adjust multi-region related vars

resource_group_name = "<your-resource-group-name>"

location_1 = "eastus"
location_2 = "westus2"

ssh_public_key_path  = "~/.ssh/id_rsa.pub"
ssh_private_key_path = "~/.ssh/id_rsa"

validator_config = {
  instance_size             = "Standard_B2s"
  is_genesis                = true
  enable_encryption_at_host = true
}

# set to `true` before the destruction procedure
disable_validator_protection = false

private_sentries_config = {
  enable        = true
  nodes_count   = 2
  instance_size = "Standard_B2s"
}

public_sentries_config = {
  enable        = true
  enable_ipv6   = false
  nodes_count   = 2
  instance_size = "Standard_B2s"

  locations = [
    1,
    2
  ]
}

observers_config = {
  enable           = true
  nodes_count      = 3
  instance_size    = "Standard_B2s"
  root_domain_name = "matterprotocol.com"
  enable_tls       = true

  locations = [
    1,
    2
  ]
  azs = [[2], [2]]
}

prometheus_config = {
  enable        = true
  instance_size = "Standard_B1s"
}
