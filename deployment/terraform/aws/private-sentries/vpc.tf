data "aws_availability_zones" "available" {
  state = "available"
}

locals {
  vpc_network_prefix = "10.10"
}

module "this_vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "3.19.0"

  name = "private-sentries-vpc"
  cidr = "${local.vpc_network_prefix}.0.0/16"

  azs = [data.aws_availability_zones.available.names[0]]

  public_subnets = ["${local.vpc_network_prefix}.1.0/24"]

  enable_nat_gateway   = true
  enable_dns_hostnames = true
  tags                 = var.tags
}
