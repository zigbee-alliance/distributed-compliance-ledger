locals {
  p2p_port        = 26656
  rpc_port        = 26657
  prometheus_port = 26660

  vpc_network_prefix = "10.0"
  internal_ips_range = "10.0.0.0/8"
  subnet_name        = "validator-subnet"
  subnet_region      = var.region
  subnet_output_key  = "${local.subnet_region}/${local.subnet_name}"

  egress_inet_tag = "egress-inet"
  validator_tag   = "validator"
}

data "google_compute_image" "ubuntu" {
  most_recent = true
  project     = "ubuntu-os-cloud"
  filter      = "(family = \"${var.os_family}\") AND (architecture = \"X86_64\")"
}


data "google_compute_zones" "available" {
}

resource "google_compute_instance" "this_node" {
  name                = "validator-node"  # FIXME copy-paste
  machine_type        = var.instance_type # FIXME default instance type
  deletion_protection = !var.disable_instance_protection
  zone                = data.google_compute_zones.available.names[0]

  boot_disk {
    initialize_params {
      image = data.google_compute_image.ubuntu.self_link
      size  = 80
    }
  }

  # to avoid changes in case new latest image released or zone changes
  lifecycle {
    ignore_changes = [boot_disk, zone]
  }

  metadata = {
    ssh-keys = "${var.ssh_username}:${file(var.ssh_public_key_path)}"
  }

  network_interface {
    subnetwork = module.this_vpc.subnets[local.subnet_output_key].name
    access_config {} # enables external IP
  }

  #   service_account {
  #     email  = var.service_account_email
  #     scopes = ["cloud-platform"]
  #   }

  connection {
    type        = "ssh"
    host        = self.network_interface.0.access_config.0.nat_ip
    user        = var.ssh_username
    private_key = file(var.ssh_private_key_path)
  }

  // ansible
  provisioner "remote-exec" {
    script = "./provisioner/install-ansible-deps.sh"
  }

  labels = var.labels

  tags = [local.validator_tag, local.egress_inet_tag]
}
