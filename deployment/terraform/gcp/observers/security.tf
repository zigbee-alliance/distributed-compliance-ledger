# TODO
# - block project wide ssh keys

resource "google_compute_firewall" "this_dev_fw_ingress_rules" {
  name        = "observer-dev-fw-ingress-rules"
  description = "Validator firewall ingress rules for development"
  network     = local.vpc.network_name

  # TODO in aws source port 22 is also specified
  allow {
    protocol = "tcp"
    ports    = ["22"]
  }

  allow {
    protocol = "icmp"
  }

  source_ranges = ["0.0.0.0/0"]

  target_tags = [local.observer_tag]
}

resource "google_compute_firewall" "this_fw_egress_rules" {
  name        = "observer-fw-egress-rules"
  description = "Observer nodes firewall egress rules"
  network     = local.vpc.network_name

  allow {
    protocol = "all"
  }

  direction = "EGRESS"

  target_tags = [local.observer_tag]
}


resource "google_compute_firewall" "this_private_fw_ingress_rules" {
  name        = "observer-private-fw-ingress-rules"
  description = "Observer nodes firewall ingress rules for internal connections"
  network     = local.vpc.network_name

  # TODO source ports are not restricted

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

  # Allow gRPC from internal IPs
  allow {
    protocol = "tcp"
    ports    = [local.grpc_port]
  }

  # Allow REST from internal IPs
  allow {
    protocol = "tcp"
    ports    = [local.rest_port]
  }

  source_ranges = [local.internal_ips_range]

  target_tags = [local.observer_tag]
}

# TODO what about NLB proxy IP
# allow access from health check ranges
resource "google_compute_firewall" "default" {
  name    = "observer-nlb-hc-fw-ingress-rule"
  network = local.vpc.network_name
  allow {
    protocol = "tcp"
  }
  source_ranges = ["130.211.0.0/22", "35.191.0.0/16"]
  target_tags   = [local.observer_nlb_tag]
}
