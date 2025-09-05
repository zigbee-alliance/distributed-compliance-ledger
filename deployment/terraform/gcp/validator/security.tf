# FIXME
# - aws tags (aka labels here)
# - block project wirde ssh keys

resource "google_compute_firewall" "this_dev_fw_ingress_rules" {
  name        = "validator-dev-fw-ingress-rules"
  description = "Validator firewall ingress rules for development"
  network     = module.this_vpc.network_name

  # FIXME in aws source port 22 is also specified
  allow {
    protocol = "tcp"
    ports    = ["22"]
  }

  allow {
    protocol = "icmp"
  }

  source_ranges = ["0.0.0.0/0"]

  target_tags   = [local.validator_tag]
}

resource "google_compute_firewall" "this_fw_egress_rules" {
  name        = "validator-fw-egress-rules"
  description = "Validator node firewall egress rules"
  network     = module.this_vpc.network_name

  allow {
    protocol = "all"
  }

  direction = "EGRESS"

  target_tags   = [local.validator_tag]
}


resource "google_compute_firewall" "this_private_fw_ingress_rules" {
  name        = "validator-private-fw-ingress-rules"
  description = "Validator node firewall ingress rules for internal connections"
  network     = module.this_vpc.network_name

  # FIXME source ports are not restricted

  # Allow p2p from internal IPs
  allow {
    protocol = "tcp"
    ports    = [local.p2p_port]
  }

  # Allow RPC from internal IPs
  allow {
    protocol = "tcp"
    ports    = [local.rpc_port]
  }

  # Allow Prometheus from internal IPs
  allow {
    protocol = "tcp"
    ports    = [local.prometheus_port]
  }

  source_ranges = ["${local.vpc_network_prefix}.0.0/8"]

  target_tags   = [local.validator_tag]
}
