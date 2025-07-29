region_1 = "us-east-2"
region_2 = "us-west-1"

ssh_public_key_path  = "~/.ssh/id_rsa.pub"
ssh_private_key_path = "~/.ssh/id_rsa"

validator_config = {
  instance_type = "t3.medium"
  is_genesis    = true
}

# set to `true` before the destruction procedure
disable_validator_protection = false

private_sentries_config = {
  enable        = true
  nodes_count   = 2
  instance_type = "t3.medium"
}

public_sentries_config = {
  enable        = true
  enable_ipv6   = false
  nodes_count   = 2
  instance_type = "t3.medium"

  regions = [
    1,
    2
  ]
}

observers_config = {
  enable           = true
  nodes_count      = 3
  instance_type    = "t3.medium"
  root_domain_name = "matterprotocol.com"
  enable_tls       = true

  regions = [
    1,
    2
  ]
}

prometheus_config = {
  enable        = true
  instance_type = "t3.small"
}