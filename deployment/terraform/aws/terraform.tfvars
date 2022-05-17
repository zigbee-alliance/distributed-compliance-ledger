region_1 = "us-west-1"
region_2 = "us-east-2"

validator_config = {
  instance_type = "t3.medium"
  enable_prometheus = true
}

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