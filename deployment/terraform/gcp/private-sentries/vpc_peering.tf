resource "google_compute_network_peering" "this_local_to_peer_vpc_peering" {
  name         = "private-sentries-local-to-peer-vpc-peering"
  network      = module.this_vpc.network_self_link
  peer_network = var.peer_vpc.network_self_link
}

resource "google_compute_network_peering" "this_peer_to_local_vpc_peering" {
  name         = "private-sentries-peer-to-local-vpc-peering"
  provider     = google.peer
  network      = var.peer_vpc.network_self_link
  peer_network = module.this_vpc.network_self_link
}

