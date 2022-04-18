data "aws_availability_zones" "available" {
    state = "available"
}

locals {
    vpc_network_prefix = "10.${20 + var.region_index}"
}

module "this_vpc" {
    source = "terraform-aws-modules/vpc/aws"
    version = "3.14.0"

    name = "public-sentries-vpc"
    cidr = "${local.vpc_network_prefix}.0.0/16"

    enable_ipv6                                     = var.enable_ipv6
    public_subnet_assign_ipv6_address_on_creation   = var.enable_ipv6
    public_subnet_ipv6_prefixes                     = var.enable_ipv6 ? [1] : []

    azs                = [data.aws_availability_zones.available.names[0]]

    public_subnets     = ["${local.vpc_network_prefix}.1.0/24"]

    enable_nat_gateway = true
    enable_dns_hostnames = true
}
