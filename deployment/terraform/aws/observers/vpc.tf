data "aws_availability_zones" "available" {
    state = "available"
}

module "this_vpc" {
    source = "terraform-aws-modules/vpc/aws"
    version = "3.14.0"

    name = "observers-vpc-1"
    cidr = "10.3.0.0/16"

    azs                = [data.aws_availability_zones.available.names[0], data.aws_availability_zones.available.names[1]]

    public_subnets     = ["10.3.1.0/24", "10.3.2.0/24"]

    enable_nat_gateway = true
    enable_dns_hostnames = true
}
