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

resource "aws_instance" "this_nodes" {
  count = var.nodes_count

  ami           = data.aws_ami.ubuntu.id
  instance_type = var.instance_type

  subnet_id          = element(module.this_vpc.public_subnets, 0)
  ipv6_address_count = var.enable_ipv6 ? 1 : 0

  vpc_security_group_ids = [
    module.this_dev_sg.security_group_id,
    module.this_public_sg.security_group_id
  ]

  key_name   = aws_key_pair.key_pair.id
  monitoring = true

  connection {
    type        = "ssh"
    host        = self.public_ip
    user        = var.ssh_username
    private_key = file(var.ssh_private_key_path)
  }

  provisioner "remote-exec" {
    script = "./provisioner/install-ansible-deps.sh"
  }

  tags = {
    Name = "Public Sentry Node ${count.index}"
  }

  root_block_device {
    encrypted   = true
    volume_size = 30
  }

  metadata_options {
    http_endpoint = "enabled"
    http_tokens   = "required"
  }
}

resource "aws_instance" "this_seed_node" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = var.instance_type

  subnet_id          = element(module.this_vpc.public_subnets, 0)
  ipv6_address_count = var.enable_ipv6 ? 1 : 0

  vpc_security_group_ids = [
    module.this_dev_sg.security_group_id,
    module.this_seed_sg.security_group_id
  ]

  key_name   = aws_key_pair.key_pair.id
  monitoring = true

  tags = {
    Name = "Public Sentries' Seed Node"
  }

  root_block_device {
    encrypted   = true
    volume_size = 30
  }

  metadata_options {
    http_endpoint = "enabled"
    http_tokens   = "required"
  }
}

resource "aws_eip" "this_nodes_eips" {
  count = var.enable_ipv6 ? 0 : length(aws_instance.this_nodes)

  instance = aws_instance.this_nodes[count.index].id
  vpc      = true

  tags = {
    Name = "Public Sentry Node [${count.index}] Elastic IP"
  }
}

resource "aws_eip" "this_seed_eip" {
  count = var.enable_ipv6 ? 0 : 1

  instance = aws_instance.this_seed_node.id
  vpc      = true

  tags = {
    Name = "Public Sentries' Seed Node Elastic IP"
  }
}
