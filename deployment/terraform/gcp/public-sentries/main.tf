locals {
  p2p_port        = 26656
  rpc_port        = 26657
  prometheus_port = 26660

  subnet_name        = "public-sentries-subnet"
  internal_ips_range = "10.0.0.0/8"

  egress_inet_tag        = "egress-inet"
  public_sentry_tag      = "public-sentry"
  public_sentry_seed_tag = "public-sentry-seed"

  subnets = [for index, config in var.region_config :
    {
      subnet_name      = "${local.subnet_name}"
      subnet_ip        = "10.${20 + index}.1.0/24"
      subnet_region    = "${config.region}"
      stack_type       = var.enable_ipv6 ? "IPV4_IPV6" : "IPV4_ONLY"
      ipv6_access_type = var.enable_ipv6 ? "EXTERNAL" : null
    }
  ]

  regions = [for config in var.region_config : config.region]

  nodes = flatten([for region_index, config in var.region_config : [
    for node_index in range(config.nodes_count) : {
      region     = config.region
      subnet_key = "${config.region}/${local.subnet_name}"
    zone = data.google_compute_zones.available[region_index].names[node_index % length(data.google_compute_zones.available[region_index].names)], }
  ]])

  seed_nodes = [for region_index, config in var.region_config : {
    region     = config.region
    subnet_key = "${config.region}/${local.subnet_name}"
    zone       = data.google_compute_zones.available[region_index].names[0]
  } if config.nodes_count > 0]
}

data "google_compute_image" "ubuntu" {
  most_recent = true
  project     = "ubuntu-os-cloud"
  filter      = "(family = \"${var.os_family}\") AND (architecture = \"X86_64\")"
}


data "google_compute_zones" "available" {
  count  = length(local.regions)
  region = local.regions[count.index]
}

resource "google_compute_address" "this_static_ips" {
  count      = length(local.nodes)
  name       = "public-sentry-node-${count.index}-static-ip"
  region     = local.nodes[count.index].region
  ip_version = var.enable_ipv6 ? "IPV6" : "IPV4"
  #subnetwork = module.this_vpc.subnets[local.subnet_output_key].name
}

resource "google_compute_address" "this_seed_static_ips" {
  count      = length(local.seed_nodes)
  name       = "public-sentry-seed-node-${count.index}-static-ip"
  region     = local.seed_nodes[count.index].region
  ip_version = var.enable_ipv6 ? "IPV6" : "IPV4"
  #subnetwork = module.this_vpc.subnets[local.subnet_output_key].name
}


resource "google_compute_instance" "this_nodes" {
  count = length(local.nodes)

  name         = "public-sentry-node-${count.index}"
  machine_type = var.instance_type
  zone         = local.nodes[count.index].zone

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
    stack_type = var.enable_ipv6 ? "IPV4_IPV6" : "IPV4_ONLY"
    subnetwork = module.this_vpc.subnets[local.nodes[count.index].subnet_key].name
    access_config {
      nat_ip = var.enable_ipv6 ? null : google_compute_address.this_static_ips[count.index].address # static external IP
    }

    dynamic "ipv6_access_config" {
      for_each = var.enable_ipv6 ? [{}] : []
      content {
        network_tier  = "PREMIUM"                                                   # required, Only PREMIUM tier is valid
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

  labels = var.labels

  tags = [local.public_sentry_tag, local.egress_inet_tag]
}

resource "google_compute_instance" "this_seed_nodes" {
  count = length(local.seed_nodes)

  name         = "public-sentry-seed-node-${count.index}"
  machine_type = var.instance_type
  zone         = local.seed_nodes[count.index].zone

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
    stack_type = var.enable_ipv6 ? "IPV4_IPV6" : "IPV4_ONLY"
    subnetwork = module.this_vpc.subnets[local.seed_nodes[count.index].subnet_key].name
    access_config {
      nat_ip = var.enable_ipv6 ? null : google_compute_address.this_seed_static_ips[count.index].address # static external IP
    }

    dynamic "ipv6_access_config" {
      for_each = var.enable_ipv6 ? [{}] : []
      content {
        network_tier  = "PREMIUM"                                                        # required, Only PREMIUM tier is validA
        external_ipv6 = google_compute_address.this_seed_static_ips[count.index].address # static external IPv6
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

  labels = var.labels

  tags = [local.public_sentry_seed_tag, local.egress_inet_tag]
}
