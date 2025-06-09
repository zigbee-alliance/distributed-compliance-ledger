module "this_dev_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = "private-sentry-dev-security-group"
  description = "Private Sentry nodes security group for development"
  vpc_id      = module.this_vpc.vpc_id

  ingress_cidr_blocks = ["0.0.0.0/0"]
  ingress_rules       = ["all-icmp", "ssh-tcp"]
  egress_rules        = ["all-all"]
  tags = var.tags
}

module "this_private_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = "private-sentry-private-security-group"
  description = "Private Sentry nodes security group for internal connections"
  vpc_id      = module.this_vpc.vpc_id

  egress_rules = ["all-all"]
  ingress_with_cidr_blocks = [
    {
      from_port   = 26656
      to_port     = 26656
      protocol    = "tcp"
      description = "Allow p2p from internal IPs"
      cidr_blocks = "10.0.0.0/8"
    },
    {
      from_port   = 26657
      to_port     = 26657
      protocol    = "tcp"
      description = "Allow RPC from internal IPs"
      cidr_blocks = "10.0.0.0/8"
    },
    {
      from_port   = 26660
      to_port     = 26660
      protocol    = "tcp"
      description = "Allow Prometheus from internal IPs"
      cidr_blocks = "10.0.0.0/8"
    },
  ]
  tags = var.tags
}

module "this_public_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = "private-sentry-public-security-group"
  description = "Private Sentry nodes security group for external connections"
  vpc_id      = module.this_vpc.vpc_id

  # ingress_cidr_blocks = ["10.0.0.0/8"]
  egress_rules = ["all-all"]


  ingress_with_cidr_blocks = [
    {
      from_port   = 26656
      to_port     = 26656
      protocol    = "tcp"
      description = "Allow P2P from Some Organization"
      cidr_blocks = "10.1.1.1/32" # whitelist IP
    },
  ]
  tags = var.tags
}
