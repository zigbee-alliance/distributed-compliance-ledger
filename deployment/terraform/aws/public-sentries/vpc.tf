data "aws_availability_zones" "available" {
  state = "available"
}

module "this_vpc" {
    source = "terraform-aws-modules/vpc/aws"
    version = "3.14.0"

    name = "public-sentries-vpc"
    cidr = "10.2.0.0/16"

    azs                = [data.aws_availability_zones.available.names[0]]

    # private_subnets    = ["10.1.1.0/24"]
    public_subnets     = ["10.2.1.0/24"]

    enable_nat_gateway = true
    enable_dns_hostnames = true
}
