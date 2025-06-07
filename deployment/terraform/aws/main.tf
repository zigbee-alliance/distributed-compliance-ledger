provider "aws" {
  alias  = "region_1"
  region = var.region_1
}

provider "aws" {
  alias  = "region_2"
  region = var.region_2
}

# IAM
module "iam" {
  source = "./iam"
  providers = {
    aws = aws.region_1
  }
}

# Validator
module "validator" {
  source = "./validator"
  providers = {
    aws = aws.region_1
  }

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  instance_type        = var.validator_config.instance_type
  iam_instance_profile = module.iam.iam_instance_profile
}

# Private Sentries
module "private_sentries" {
  count = var.private_sentries_config.enable ? 1 : 0

  source = "./private-sentries"

  nodes_count          = var.private_sentries_config.nodes_count
  instance_type        = var.private_sentries_config.instance_type
  iam_instance_profile = module.iam.iam_instance_profile

  providers = {
    aws      = aws.region_1
    aws.peer = aws.region_1
  }

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  peer_vpc = module.validator.vpc
}

# Public Sentries region 1
module "public_sentries_1" {
  count = (var.private_sentries_config.enable &&
    var.public_sentries_config.enable &&
  contains(var.public_sentries_config.regions, 1)) ? 1 : 0

  source = "./public-sentries"

  nodes_count          = var.public_sentries_config.nodes_count
  instance_type        = var.public_sentries_config.instance_type
  iam_instance_profile = module.iam.iam_instance_profile

  enable_ipv6 = var.public_sentries_config.enable_ipv6

  providers = {
    aws      = aws.region_1
    aws.peer = aws.region_1
  }

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  region_index = 1
  peer_vpc     = module.private_sentries[0].vpc
}

# Public Sentries region 2
module "public_sentries_2" {
  count = (var.private_sentries_config.enable &&
    var.public_sentries_config.enable &&
  contains(var.public_sentries_config.regions, 2)) ? 1 : 0

  source = "./public-sentries"

  nodes_count          = var.public_sentries_config.nodes_count
  instance_type        = var.public_sentries_config.instance_type
  iam_instance_profile = module.iam.iam_instance_profile

  enable_ipv6 = var.public_sentries_config.enable_ipv6

  providers = {
    aws      = aws.region_2
    aws.peer = aws.region_1
  }

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  region_index = 2
  peer_vpc     = module.private_sentries[0].vpc
}

# Observers region 1
module "observers_1" {
  count = (var.private_sentries_config.enable &&
    var.observers_config.enable &&
  contains(var.observers_config.regions, 1)) ? 1 : 0

  source = "./observers"

  nodes_count          = var.observers_config.nodes_count
  instance_type        = var.observers_config.instance_type
  iam_instance_profile = module.iam.iam_instance_profile

  root_domain_name = var.observers_config.root_domain_name

  enable_tls = var.observers_config.enable_tls

  providers = {
    aws      = aws.region_1
    aws.peer = aws.region_1
  }

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  region_index = 1
  peer_vpc     = module.private_sentries[0].vpc
}

# Observers region 2
module "observers_2" {
  count = (var.private_sentries_config.enable &&
    var.observers_config.enable &&
  contains(var.observers_config.regions, 2)) ? 1 : 0

  source = "./observers"

  nodes_count          = var.observers_config.nodes_count
  instance_type        = var.observers_config.instance_type
  iam_instance_profile = module.iam.iam_instance_profile

  root_domain_name = var.observers_config.root_domain_name

  enable_tls = var.observers_config.enable_tls

  providers = {
    aws      = aws.region_2
    aws.peer = aws.region_1
  }

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  region_index = 2
  peer_vpc     = module.private_sentries[0].vpc
}

module "prometheus" {
  count = local.prometheus_enabled ? 1 : 0

  source        = "./prometheus"
  instance_type = var.prometheus_config.instance_type

  endpoints = local.prometheus_endpoints

  providers = {
    aws = aws.region_1
  }

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  vpc = module.private_sentries[0].vpc
}
