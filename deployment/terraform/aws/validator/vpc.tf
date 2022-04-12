data "aws_availability_zones" "available" {
    state = "available"
}
module "this_vpc" {
    source = "terraform-aws-modules/vpc/aws"
    version = "3.14.0"

    name = "validator-vpc"
    cidr = "10.0.0.0/16"

    azs                = [data.aws_availability_zones.available.names[0]]

    # private_subnets    = ["10.0.1.0/24"]
    public_subnets     = ["10.0.1.0/24"]

    enable_nat_gateway = true
    enable_dns_hostnames = true
}