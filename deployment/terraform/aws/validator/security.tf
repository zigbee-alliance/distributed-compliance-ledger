module "this_dev_sg" {
    source  = "terraform-aws-modules/security-group/aws"
    version = "~> 4.0"

    name        = "validator-dev-security-group"
    description = "Validator security group for development"
    vpc_id      = module.this_vpc.vpc_id

    ingress_cidr_blocks = ["0.0.0.0/0"]
    ingress_rules       = ["all-icmp", "ssh-tcp"]
    egress_rules        = ["all-all"]
}
module "this_private_sg" {
    source  = "terraform-aws-modules/security-group/aws"
    version = "~> 4.0"

    name        = "validator-private-security-group"
    description = "Validator node security group for internal connections"
    vpc_id      = module.this_vpc.vpc_id

    egress_rules        = ["all-all"]
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
    ]
}