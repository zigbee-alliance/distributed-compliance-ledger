resource "google_compute_network_peering" "this_private_sentries_to_validator_vpc_peering" {
  name         = "private-sentries-to-validator-vpc-peering"
  network      = module.this_vpc.network_self_link
  peer_network = var.peer_vpc.network_self_link
}

resource "google_compute_network_peering" "this_validator_to_private_sentries_vpc_peering" {
  name         = "validator-to-private-sentries-vpc-peering"
  provider     = google.peer
  network      = var.peer_vpc.network_self_link
  peer_network = module.this_vpc.network_self_link
}

