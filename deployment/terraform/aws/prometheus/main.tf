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

resource "aws_instance" "this_node" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = var.instance_type

  subnet_id = element(var.vpc.public_subnets, 0)

  vpc_security_group_ids = [
    module.this_dev_sg.security_group_id,
  ]

  key_name   = aws_key_pair.key_pair.id
  monitoring = true

  iam_instance_profile = aws_iam_instance_profile.this_amp_role_profile.name

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
    content     = local.prometheus_config
    destination = "/tmp/prometheus.yml"
  }

  provisioner "remote-exec" {
    script = "./provisioner/install-prometheus-service.sh"
  }

  tags = merge(var.tags, {
    Name = "Prometheus Server Node"
  })

  root_block_device {
    encrypted   = true
    volume_size = 20
  }

  metadata_options {
    http_endpoint = "enabled"
    http_tokens   = "required"
  }
}
