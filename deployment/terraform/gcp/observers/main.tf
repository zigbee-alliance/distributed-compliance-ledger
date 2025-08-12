locals {
  grpc_port = 9090
  rest_port = 1317
  p2p_port = 26656
  rpc_port = 26657
  prometheus_port = 26660

  vpc = var.vpc_name == null ? module.this_vpc : data.google_compute_network.vpc[0]

  vpc_network_prefix = "10.${30 + var.region_index}"
  internal_ips_prefix = "10.0"
  subnet_name_prefix = "observers-subnet"

  subnet_region = var.region
  subnet_output_key_prefix = "${local.subnet_region}/${local.subnet_name_prefix}"

  egress_inet_tag = "egress-inet"
  observer_tag = "observer"
}


data "google_compute_network" "vpc" {
  count = var.vpc_name == null ? 0 : 1
  name = var.vpc_name
}

data "google_compute_image" "ubuntu" {
  most_recent = true
  project = "ubuntu-os-cloud"
  filter = "(family = \"${var.os_family}\") AND (architecture = \"X86_64\")"
}


data "google_compute_zones" "available" {
}

resource "google_compute_instance" "this_nodes" {
  count = var.nodes_count

  name                = "observer-node-${count.index}" # FIXME copy-paste
  machine_type        = var.instance_type
  # FIXME 
  # - count.index might not fit available zones
  # - aws logic implements distibution across public zones in vpc
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
    subnetwork = local.vpc.subnets["${local.subnet_output_key_prefix}-${count.index % length(local.vpc.subnets)}"].name
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

  labels = var.labels # FIXME gcp.labels == aws.tags

  tags = [local.observer_tag, local.egress_inet_tag]
}
