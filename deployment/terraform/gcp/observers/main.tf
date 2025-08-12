locals {
  grpc_port = 9090
  rest_port = 1317
  p2p_port = 26656
  rpc_port = 26657
  prometheus_port = 26660

  vpc = module.this_vpc

  subnet_name_prefix = "observers-subnet"
  internal_ips_prefix = "10.0"

  egress_inet_tag = "egress-inet"
  observer_tag = "observer"

  subnets = flatten([ for index, config in var.region_config : [
    {
      subnet_name     = "${local.subnet_name_prefix}-0"
      subnet_ip       = "10.${30 + index}.1.0/24"
      subnet_region   = "${config.region}"
    },
    {
      subnet_name     = "${local.subnet_name_prefix}-1"
      subnet_ip       = "10.${30 + index}.2.0/24"
      subnet_region   = "${config.region}"
    },
  ]])

  nodes = flatten([ for region_index, config in var.region_config : [ 
    for node_index in range(config.nodes_count) : {
      region_index = region_index
      subnet_key = "${config.region}/${local.subnet_name_prefix}-${node_index % 2}"
      zone = data.google_compute_zones.available[region_index].names[node_index % length(data.google_compute_zones.available[region_index].names)]
    }
  ]])

  regions = [ for config in var.region_config : config.region ]
}


data "google_compute_image" "ubuntu" {
  most_recent = true
  project = "ubuntu-os-cloud"
  filter = "(family = \"${var.os_family}\") AND (architecture = \"X86_64\")"
}


data "google_compute_zones" "available" {
  count = length(local.regions)
  region = local.regions[count.index]
}

resource "google_compute_instance" "this_nodes" {
  count = length(local.nodes)

  name          = "observer-node-${count.index}" # FIXME copy-paste
  machine_type  = var.instance_type
  zone          = local.nodes[count.index].zone

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
    subnetwork = local.vpc.subnets[local.nodes[count.index].subnet_key].name
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
