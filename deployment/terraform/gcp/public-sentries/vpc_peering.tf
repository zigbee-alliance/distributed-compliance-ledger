resource "google_compute_network_peering" "this_public_to_private_sentries_vpc_peering" {
  name         = "public-to-private-sentries-vpc-peering"
  network      = module.this_vpc.network_self_link
  peer_network = var.peer_vpc.network_self_link
}

resource "google_compute_network_peering" "this_private_to_public_sentries_vpc_peering" {
  name         = "private-to-public-sentries-vpc-peering"
  provider     = google.peer
  network      = var.peer_vpc.network_self_link
  peer_network = module.this_vpc.network_self_link
}

