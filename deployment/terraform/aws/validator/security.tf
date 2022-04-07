module "validator_sg" {
    source  = "terraform-aws-modules/security-group/aws"
    version = "~> 4.0"

    name        = "validator-security-group"
    description = "Security group for Validator nodes from outside"
    vpc_id      = module.validator_vpc.vpc_id

    ingress_cidr_blocks = ["0.0.0.0/0"]
    ingress_rules       = ["all-icmp", "ssh-tcp"]
    egress_rules        = ["all-all"]
    ingress_with_cidr_blocks = [
        {
            from_port   = 26656
            to_port     = 26656
            protocol    = "tcp"
            description = "Validator p2p"
            cidr_blocks = "10.0.0.0/8"
        },
        {
            from_port   = 26657
            to_port     = 26657
            protocol    = "tcp"
            description = "Validator RPC"
            cidr_blocks = "10.0.0.0/8"
        },
  ]
}