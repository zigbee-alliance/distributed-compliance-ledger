# FIXME adjust multi-region related vars

resource_group_name = "<your-resource-group-name>"

location_1 = "eastus"
location_2 = "westus2"

ssh_public_key_path  = "~/.ssh/id_rsa.pub"
ssh_private_key_path = "~/.ssh/id_rsa"

validator_config = {
  instance_size             = "Standard_D2as_v5"
  is_genesis                = true
  enable_encryption_at_host = true
}

# set to `true` before the destruction procedure
disable_validator_protection = false

private_sentries_config = {
  enable        = true
  nodes_count   = 2
  instance_size = "Standard_D2as_v5"
}

public_sentries_config = {
  enable        = true
  enable_ipv6   = false # TODO
  nodes_count   = 2
  instance_size = "Standard_D2as_v5"

  locations = [
    1,
    2
  ]
}

observers_config = {
  enable           = true
  nodes_count      = 3
  instance_size    = "Standard_D2as_v5"
  root_domain_name = "matterprotocol.com"
  enable_tls       = false # TODO

  locations = [
    1,
    2
  ]
  azs = [[2], [2]]
}

prometheus_config = {
  enable        = false # TODO
  instance_size = "Standard_B1s"
}
