data "aws_availability_zones" "available" {
    state = "available"
}
module "validator_vpc" {
    source = "terraform-aws-modules/vpc/aws"

    name = "validator-vpc"
    cidr = "10.0.0.0/16"

    azs                = [data.aws_availability_zones.available.names[0]]
    private_subnets    = ["10.0.1.0/24"]
    public_subnets     = ["10.0.101.0/24"]

    enable_nat_gateway = true
    enable_dns_hostnames = true
}