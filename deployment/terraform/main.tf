terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 3.72"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  region = var.region
}

module "dcl_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = "dcl-security_group"
  description = "Security group for accessing DCL nodes from outside"
  vpc_id      = module.network_lab.vpc_id

  ingress_cidr_blocks = ["0.0.0.0/0"]
  ingress_rules       = ["all-icmp", "ssh-tcp"]
  egress_rules        = ["all-all"]
  ingress_with_cidr_blocks = [
    {
      from_port   = 26656
      to_port     = 26656
      protocol    = "tcp"
      description = "DCL p2p"
      cidr_blocks = "0.0.0.0/0"
    },
    {
      from_port   = 26657
      to_port     = 26657
      protocol    = "tcp"
      description = "DCL RPC"
      cidr_blocks = "0.0.0.0/0"
    },
  ]
}

resource "aws_instance" "dcl_node" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = "c5.4xlarge"

  subnet_id              = element(module.network_lab.public_subnets, 0)
  vpc_security_group_ids = [module.dcl_sg.security_group_id]

  key_name   = aws_key_pair.key_pair.id
  monitoring = true

  root_block_device {
    encrypted   = true
    volume_size = 20
  }

  connection {
    type        = "ssh"
    host        = self.public_ip
    user        = var.ssh_username
    private_key = file(var.ssh_private_key_path)
  }

  provisioner "remote-exec" {
    inline = [
      "sudo apt-get update",
      "sudo apt-get install -y --no-install-recommends python3",
    ]
  }

  provisioner "local-exec" {
    command = "ansible-playbook -i ../ansible/inventory/aws_ec2.yml -u ${var.ssh_username} ../ansible/deploy.yml"
    environment = {
      ANSIBLE_HOST_KEY_CHECKING = "False"
    }
  }

  metadata_options {
    http_endpoint = "enabled"
    http_tokens   = "required"
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

data "aws_availability_zones" "available" {
  state = "available"
}

module "network_lab" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 3.0"

  name = "dcl-network"
  cidr = "10.0.0.0/16"

  azs                = [data.aws_availability_zones.available.names[0]]
  private_subnets    = ["10.0.1.0/24"]
  public_subnets     = ["10.0.101.0/24"]
  enable_nat_gateway = true
}
