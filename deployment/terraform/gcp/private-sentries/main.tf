locals {
  p2p_port = 26656
  rpc_port = 26657
  prometheus_port = 26660

  vpc_network_prefix = "10.10"
  internal_ips_prefix = "10.0"
  subnet_name = "private-sentries-subnet"

  subnet_region = var.region
  subnet_output_key = "${local.subnet_region}/${local.subnet_name}"

  private_sentry_tag = "private-sentry"
}

data "google_compute_image" "ubuntu" {
  most_recent = true
  project = "ubuntu-os-cloud"
  filter = "(family = \"${var.os_family}\") AND (architecture = \"X86_64\")"
}


data "google_compute_zones" "available" {
}

resource "google_compute_address" "this_static_ips" {
  count = var.nodes_count > 0 ? 1 : 0
  name  = "private-sentry-node-${count.index}-static-ip"
}

resource "google_compute_instance" "this_nodes" {
  count = var.nodes_count

  name                = "private-sentry-node-${count.index}" # FIXME copy-paste
  machine_type        = var.instance_type
  zone               = data.google_compute_zones.available.names[count.index]

  boot_disk {
    initialize_params {
      image = data.google_compute_image.ubuntu.self_link
      size = 80
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
    access_config {
      nat_ip = count.index == 0 ? google_compute_address.this_static_ips.0.address : null # static external IP
    }
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

  labels = var.labels # FIXME gcp.labels == aws.tags

  tags = [local.private_sentry_tag]
}
