module "this_dev_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = "observer-dev-security-group"
  description = "Observer nodes security group for development"
  tags        = var.tags

  vpc_id = module.this_vpc.vpc_id

  ingress_cidr_blocks = ["0.0.0.0/0"]
  ingress_rules       = ["all-icmp", "ssh-tcp"]
  egress_rules        = ["all-all"]
}

module "this_private_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = "observer-private-security-group"
  description = "Observer nodes security group for internal connections"
  tags        = var.tags

  vpc_id = module.this_vpc.vpc_id

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
      from_port   = 9090
      to_port     = 9090
      protocol    = "tcp"
      description = "Allow gRPC from internal IPs"
      cidr_blocks = "10.0.0.0/8"
    },
    {
      from_port   = 1317
      to_port     = 1317
      protocol    = "tcp"
      description = "Allow REST from internal IPs"
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
}
