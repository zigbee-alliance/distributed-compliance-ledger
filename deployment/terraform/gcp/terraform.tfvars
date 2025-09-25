region_1 = "us-east2"
region_2 = "us-west1"

ssh_public_key_path  = "~/.ssh/id_rsa.pub"
ssh_private_key_path = "~/.ssh/id_rsa"

project_id = "<your-project-id>"

validator_config = {
  instance_type = "n2-standard-2"
  is_genesis    = true
}

# set to `true` before the destruction procedure
disable_validator_protection = false

private_sentries_config = {
  enable        = true
  nodes_count   = 2
  instance_type = "n2-standard-2"
}

public_sentries_config = {
  enable              = true
  enable_ipv6         = false # TODO
  region1_nodes_count = 2
  region2_nodes_count = 2
  instance_type       = "n2-standard-2"
}

observers_config = {
  enable              = true
  region1_nodes_count = 3
  region2_nodes_count = 3
  instance_type       = "n2-standard-2"
  root_domain_name    = "matterprotocol.com"
  enable_tls          = false # TODO
}

prometheus_config = {
  enable        = false # TODO
  instance_type = "n2-medium"
}
