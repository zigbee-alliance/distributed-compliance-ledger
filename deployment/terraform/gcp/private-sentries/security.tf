# FIXME
# - aws tags (aka labels here)
# - block project wirde ssh keys

resource "google_compute_firewall" "this_dev_fw_ingress_rules" {
  name        = "private-sentry-dev-fw-ingress-rules"
  description = "Private Sentry nodes firewall ingress rules for development"
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

  target_tags   = [local.private_sentry_tag]
}

resource "google_compute_firewall" "this_fw_egress_rules" {
  name        = "private-sentry-fw-egress-rules"
  description = "Private Sentry nodes firewall egress rules"
  network     = module.this_vpc.network_name

  allow {
    protocol = "all"
  }

  direction = "EGRESS"

  target_tags   = [local.private_sentry_tag]
}


resource "google_compute_firewall" "this_private_fw_ingress_rules" {
  name        = "private-sentry-private-fw-ingress-rules"
  description = "Private Sentry nodes firewall ingress rules for internal connections"
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

  source_ranges = [local.internal_ips_range]

  target_tags   = [local.private_sentry_tag]
}

resource "google_compute_firewall" "this_public_fw_ingress_rules" {
  name        = "private-sentry-public-fw-ingress-rules"
  description = "Private Sentry nodes firewall ingress rules for external connections"
  network     = module.this_vpc.network_name

  # FIXME source ports are not restricted

  # Allow P2P from Some Organization
  allow {
    protocol = "tcp"
    ports    = [local.p2p_port]
  }

  # TODO make a variable to link to an external scope
  source_ranges = ["10.1.1.1/32"] # whitelist IP

  target_tags   = [local.private_sentry_tag]
}
