module "this_dev_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = "public-sentry-dev-security-group"
  description = "Public Sentry nodes security group for development"
  vpc_id      = module.this_vpc.vpc_id

  ingress_cidr_blocks      = ["0.0.0.0/0"]
  ingress_ipv6_cidr_blocks = ["::/0"]
  ingress_rules            = ["all-icmp", "ssh-tcp"]
  egress_rules             = ["all-all"]
}

module "this_public_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = "public-sentry-public-security-group"
  description = "Public Sentry nodes security group for external connections"
  vpc_id      = module.this_vpc.vpc_id

  egress_rules = ["all-all"]
  ingress_with_cidr_blocks = [
    {
      from_port   = 26656
      to_port     = 26656
      protocol    = "tcp"
      description = "Allow p2p from all external IPs"
      cidr_blocks = "0.0.0.0/0"
    },
    {
      from_port   = 26657
      to_port     = 26657
      protocol    = "tcp"
      description = "Allow RPC from all external IPs"
      cidr_blocks = "0.0.0.0/0"
    },
    {
      from_port   = 26660
      to_port     = 26660
      protocol    = "tcp"
      description = "Allow Prometheus from internal IPs"
      cidr_blocks = "10.0.0.0/8"
    },
  ]

  ingress_with_ipv6_cidr_blocks = [
    {
      from_port        = 26656
      to_port          = 26656
      protocol         = "tcp"
      description      = "Allow p2p from all external IPs"
      ipv6_cidr_blocks = "::/0"
    },
    {
      from_port        = 26657
      to_port          = 26657
      protocol         = "tcp"
      description      = "Allow RPC from all external IPs"
      ipv6_cidr_blocks = "::/0"
    },
  ]
}

module "this_seed_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = "public-sentries-seed-security-group"
  description = "Public Sentries Seed node security group for external connections"
  vpc_id      = module.this_vpc.vpc_id

  # ingress_cidr_blocks = ["10.0.0.0/8"]
  egress_rules = ["all-all"]


  ingress_with_cidr_blocks = [
    {
      from_port   = 26656
      to_port     = 26656
      protocol    = "tcp"
      description = "Allow P2P from all"
      cidr_blocks = "0.0.0.0/0"
    },
  ]
}
