locals {
  p2p_port = 26656
  rpc_port = 26657
  prometheus_port = 26660

  vpc_network_prefix = "10.${20 + var.region_index}"
  internal_ips_prefix = "10.0"
  subnet_name = "public-sentries-subnet"

  subnet_region = var.region
  subnet_output_key = "${local.subnet_region}/${local.subnet_name}"

  egress_inet_tag = "egress-inet"
  public_sentry_tag = "public-sentry"
  public_sentry_seed_tag = "public-sentry-seed"
}

data "google_compute_image" "ubuntu" {
  most_recent = true
  project = "ubuntu-os-cloud"
  filter = "(family = \"${var.os_family}\") AND (architecture = \"X86_64\")"
}


data "google_compute_zones" "available" {
}

# FIXME vpc = true in aws
resource "google_compute_address" "this_static_ips" {
  count = var.nodes_count
  name  = "public-sentry-node-${count.index}-static-ip"
  ip_version = var.enable_ipv6 ? "IPV6" : "IPV4"
}

# FIXME vpc = true in aws
resource "google_compute_address" "this_seed_static_ip" {
  name  = "public-sentry-seed-node-static-ip"
  ip_version = var.enable_ipv6 ? "IPV6" : "IPV4"
}


resource "google_compute_instance" "this_nodes" {
  count = var.nodes_count

  name                = "public-sentry-node-${count.index}" # FIXME copy-paste
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
    stack_type            = var.enable_ipv6 ? "IPV4_IPV6" : "IPV4_ONLY"
    subnetwork = module.this_vpc.subnets[local.subnet_output_key].name
    access_config {
      nat_ip = var.enable_ipv6 ? null : google_compute_address.this_static_ips[count.index].address # static external IP
    }

    dynamic "ipv6_access_config" {
      for_each = var.enable_ipv6 ? [{}] : []
      content {
        network_tier = "PREMIUM" # required, Only PREMIUM tier is valid
        external_ipv6 = google_compute_address.this_static_ips[count.index].address # static external IPv6
      }
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

  tags = [local.public_sentry_tag, local.egress_inet_tag]
}

resource "google_compute_instance" "this_seed_node" {
  name                = "public-sentry-seed-node"
  machine_type        = var.instance_type
  zone               = data.google_compute_zones.available.names[0]

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
    stack_type            = var.enable_ipv6 ? "IPV4_IPV6" : "IPV4_ONLY"
    subnetwork = module.this_vpc.subnets[local.subnet_output_key].name
    access_config {
      nat_ip = var.enable_ipv6 ? google_compute_address.this_static_ips.0.address : null # static external IP
    }

    dynamic "ipv6_access_config" {
      for_each = var.enable_ipv6 ? [{}] : []
      content {
        network_tier = "PREMIUM" # required, Only PREMIUM tier is validA
        external_ipv6 = google_compute_address.this_seed_static_ip.address # static external IPv6
      }
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

  tags = [local.public_sentry_seed_tag, local.egress_inet_tag]
}
