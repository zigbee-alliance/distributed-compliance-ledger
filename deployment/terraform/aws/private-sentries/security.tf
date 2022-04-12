module "this_dev_sg" {
    source  = "terraform-aws-modules/security-group/aws"
    version = "~> 4.0"

    name        = "private-sentry-maintenance-security-group"
    description = "Private Sentry nodes security group for development"
    vpc_id      = module.this_vpc.vpc_id

    ingress_cidr_blocks = ["0.0.0.0/0"]
    ingress_rules       = ["all-icmp", "ssh-tcp"]
    egress_rules        = ["all-all"]
}

module "this_private_sg" {
    source  = "terraform-aws-modules/security-group/aws"
    version = "~> 4.0"

    name        = "private-sentry-internal-security-group"
    description = "Private Sentry nodes security group for internal connections"
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

module "this_public_sg" {
    source  = "terraform-aws-modules/security-group/aws"
    version = "~> 4.0"

    name        = "private-sentry-internal-security-group"
    description = "Security group for Private Sentry nodes internal connections"
    vpc_id      = module.this_vpc.vpc_id

    # ingress_cidr_blocks = ["10.0.0.0/8"]
    egress_rules        = ["all-all"]


    ingress_with_cidr_blocks = [
        {
            from_port   = 26656
            to_port     = 26656
            protocol    = "tcp"
            description = "Allow P2P Some Organization"
            cidr_blocks = "10.1.1.1/32" # whitelist IP
        },
    ]
}