# TODO aws logic explicitly defines route ids to exchange this_rts_ids / peer_rts_ids

resource "google_compute_network_peering" "this_observers_to_private_sentries_vpc_peering" {
  name         = "observers-to-private-sentries-vpc-peering"
  network      = local.vpc.network_self_link
  peer_network = var.peer_vpc.network_self_link
}

resource "google_compute_network_peering" "this_private_sentries_to_observers_vpc_peering" {
  name         = "private-sentries-to-observers-vpc-peering"
  network      = var.peer_vpc.network_self_link
  peer_network = local.vpc.network_self_link
}

