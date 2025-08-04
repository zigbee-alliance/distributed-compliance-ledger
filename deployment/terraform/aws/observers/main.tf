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
  tags       = var.tags
}

resource "aws_instance" "this_nodes" {
  count = var.nodes_count

  ami           = data.aws_ami.ubuntu.id
  instance_type = var.instance_type

  iam_instance_profile = var.iam_instance_profile.name

  subnet_id = element(module.this_vpc.public_subnets, count.index % length(module.this_vpc.public_subnets))
  vpc_security_group_ids = [
    module.this_dev_sg.security_group_id,
    module.this_private_sg.security_group_id
  ]

  key_name   = aws_key_pair.key_pair.id
  monitoring = true

  lifecycle {
    ignore_changes = [ami]
  }

  connection {
    type        = "ssh"
    host        = self.public_ip
    user        = var.ssh_username
    private_key = file(var.ssh_private_key_path)
  }

  provisioner "file" {
    content     = templatefile("./provisioner/cloudwatch-config.tpl", {})
    destination = "/tmp/cloudwatch-config.json"
  }

  provisioner "remote-exec" {
    script = "./provisioner/install-cloudwatch.sh"
  }

  provisioner "remote-exec" {
    script = "./provisioner/install-ansible-deps.sh"
  }

  tags = merge(var.tags, {
    Name = "Observer Node [${count.index}]"
  })

  root_block_device {
    encrypted   = true
    volume_size = 80
  }

  metadata_options {
    http_endpoint = "enabled"
    http_tokens   = "required"
  }
}
