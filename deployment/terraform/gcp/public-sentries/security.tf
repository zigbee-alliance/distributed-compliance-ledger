# TODO
# - block project wide ssh keys

resource "google_compute_firewall" "this_dev_fw_ingress_rules" {
  name        = "public-sentry-dev-fw-ingress-rules"
  description = "Public Sentry nodes firewall ingress rules for development"
  network     = module.this_vpc.network_name

  # TODO in aws source port 22 is also specified
  allow {
    protocol = "tcp"
    ports    = ["22"]
  }

  allow {
    protocol = "icmp"
  }

  source_ranges = ["0.0.0.0/0"]

  target_tags = [local.public_sentry_tag, local.public_sentry_seed_tag]
}

resource "google_compute_firewall" "this_dev_fw_ingress_rules_ipv6" {
  count       = var.enable_ipv6 ? 1 : 0
  name        = "public-sentry-dev-fw-ingress-rules-ipv6"
  description = "Public Sentry nodes firewall IPv6 ingress rules for development"
  network     = module.this_vpc.network_name

  # TODO in aws source port 22 is also specified
  allow {
    protocol = "tcp"
    ports    = ["22"]
  }

  allow {
    protocol = "58" # ICMPv6, https://github.com/hashicorp/terraform-provider-google/issues/16600
  }

  source_ranges = ["::/0"]

  target_tags = [local.public_sentry_tag, local.public_sentry_seed_tag]
}



resource "google_compute_firewall" "this_fw_egress_rules" {
  name        = "public-sentry-fw-egress-rules"
  description = "Public Sentry nodes firewall egress rules"
  network     = module.this_vpc.network_name

  allow {
    protocol = "all"
  }

  direction = "EGRESS"

  target_tags = [local.public_sentry_tag, local.public_sentry_seed_tag]
}


resource "google_compute_firewall" "this_public_fw_ingress_rules" {
  name        = "public-sentry-public-fw-ingress-rules"
  description = "Public Sentry nodes firewall ingress rules for external connections"
  network     = module.this_vpc.network_name

  # Allow p2p from all external IPs
  allow {
    protocol = "tcp"
    ports    = [local.p2p_port]
  }

  # Allow RPC from all external IPs
  allow {
    protocol = "tcp"
    ports    = [local.rpc_port]
  }

  source_ranges = ["0.0.0.0/0"]

  target_tags = [local.public_sentry_tag]
}

resource "google_compute_firewall" "this_public_fw_ingress_rules_ipv6" {
  count       = var.enable_ipv6 ? 1 : 0
  name        = "public-sentry-public-fw-ingress-rules-ipv6"
  description = "Public Sentry nodes firewall IPv6 ingress rules for external connections"
  network     = module.this_vpc.network_name

  # Allow p2p from all external IPs
  allow {
    protocol = "tcp"
    ports    = [local.p2p_port]
  }

  # Allow RPC from all external IPs
  allow {
    protocol = "tcp"
    ports    = [local.rpc_port]
  }

  source_ranges = ["::/0"]

  target_tags = [local.public_sentry_tag]
}


resource "google_compute_firewall" "this_prometheus_fw_ingress_rules" {
  name        = "public-sentry-prometheus-fw-ingress-rules"
  description = "Public Sentry nodes firewall ingress rules for internal prometheus connections"
  network     = module.this_vpc.network_name

  # TODO source ports are not restricted

  # Allow Prometheus from internal IPs
  allow {
    protocol = "tcp"
    ports    = [local.prometheus_port]
  }

  source_ranges = [local.internal_ips_range]

  target_tags = [local.public_sentry_tag]
}

resource "google_compute_firewall" "this_public_seed_fw_ingress_rules" {
  name        = "public-sentry-public-seed-fw-ingress-rules"
  description = "Public Sentry Seed node firewall ingress rules for external connections"
  network     = module.this_vpc.network_name

  # Allow P2P from all
  allow {
    protocol = "tcp"
    ports    = [local.p2p_port]
  }

  source_ranges = ["0.0.0.0/0"]

  target_tags = [local.public_sentry_seed_tag]
}
