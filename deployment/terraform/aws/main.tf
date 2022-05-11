provider "aws" {
  alias  = "region_1"
  region = var.region_1
}

provider "aws" {
  alias  = "region_2"
  region = var.region_2
}

# Validator
module "validator" {
  source = "./validator"
  providers = {
    aws = aws.region_1
  }

  instance_type = var.validator_config.instance_type
}

# Private Sentries
module "private_sentries" {
  count = var.private_sentries_config.enable ? 1 : 0

  source = "./private-sentries"

  nodes_count   = var.private_sentries_config.nodes_count
  instance_type = var.private_sentries_config.instance_type
  providers = {
    aws      = aws.region_1
    aws.peer = aws.region_1
  }

  peer_vpc = module.validator.vpc
}

# Public Sentries region 1
module "public_sentries_1" {
  count = (var.private_sentries_config.enable &&
    var.public_sentries_config.enable &&
  contains(var.public_sentries_config.regions, 1)) ? 1 : 0

  source = "./public-sentries"

  nodes_count   = var.public_sentries_config.nodes_count
  instance_type = var.public_sentries_config.instance_type

  enable_ipv6 = var.public_sentries_config.enable_ipv6

  providers = {
    aws      = aws.region_1
    aws.peer = aws.region_1
  }

  region_index = 1
  peer_vpc     = module.private_sentries[0].vpc
}

# Public Sentries region 2
module "public_sentries_2" {
  count = (var.private_sentries_config.enable &&
    var.public_sentries_config.enable &&
  contains(var.public_sentries_config.regions, 2)) ? 1 : 0

  source = "./public-sentries"

  nodes_count   = var.public_sentries_config.nodes_count
  instance_type = var.public_sentries_config.instance_type

  enable_ipv6 = var.public_sentries_config.enable_ipv6

  providers = {
    aws      = aws.region_2
    aws.peer = aws.region_1
  }

  region_index = 2
  peer_vpc     = module.private_sentries[0].vpc
}

# Observers region 1
module "observers_1" {
  count = (var.private_sentries_config.enable &&
    var.observers_config.enable &&
  contains(var.observers_config.regions, 1)) ? 1 : 0

  source = "./observers"

  nodes_count   = var.observers_config.nodes_count
  instance_type = var.observers_config.instance_type

  root_domain_name = var.observers_config.root_domain_name

  enable_tls = var.observers_config.enable_tls

  providers = {
    aws      = aws.region_1
    aws.peer = aws.region_1
  }

  region_index = 1
  peer_vpc     = module.private_sentries[0].vpc
}

# Observers region 2
module "observers_2" {
  count = (var.private_sentries_config.enable &&
    var.observers_config.enable &&
  contains(var.observers_config.regions, 2)) ? 1 : 0

  source = "./observers"

  nodes_count   = var.observers_config.nodes_count
  instance_type = var.observers_config.instance_type

  root_domain_name = var.observers_config.root_domain_name

  enable_tls = var.observers_config.enable_tls

  providers = {
    aws      = aws.region_2
    aws.peer = aws.region_1
  }

  region_index = 2
  peer_vpc     = module.private_sentries[0].vpc
}