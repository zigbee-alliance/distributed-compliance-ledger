terraform {
    required_providers {
        aws = {
            source  = "hashicorp/aws"
            version = ">= 2.7.0"
        }
    }
}


data "aws_ami" "ubuntu" {
    most_recent = true
    owners      = ["099720109477"]

    filter {
        name   = "name"
        values = ["ubuntu-minimal/images/hvm-ssd/ubuntu-focal-20.04-amd64-minimal-*"]
    }

    filter {
        name   = "virtualization-type"
        values = ["hvm"]
    }
}

resource "aws_key_pair" "key_pair" {
    public_key = file(var.ssh_public_key_path)
}

resource "aws_instance" "validator_node" {
    ami           = data.aws_ami.ubuntu.id
    instance_type = "t3.medium"

    subnet_id              = element(module.validator_vpc.public_subnets, 0)
    vpc_security_group_ids = [module.validator_sg.security_group_id]

    key_name   = aws_key_pair.key_pair.id
    monitoring = true

    root_block_device {
        encrypted   = true
        volume_size = 30
    }
    metadata_options {
        http_endpoint = "enabled"
        http_tokens   = "required"
    }
}