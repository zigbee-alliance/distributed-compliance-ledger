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

resource "aws_instance" "this_node" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = "t3.medium"

  subnet_id = element(module.this_vpc.public_subnets, 0)
  vpc_security_group_ids = [
    module.this_dev_sg.security_group_id,
    module.this_private_sg.security_group_id
  ]

  key_name   = aws_key_pair.key_pair.id
  monitoring = true

  tags = {
    Name = "Validator Node"
  }

  root_block_device {
    encrypted   = true
    volume_size = 30
  }

  metadata_options {
    http_tokens = "required"
  }
}
