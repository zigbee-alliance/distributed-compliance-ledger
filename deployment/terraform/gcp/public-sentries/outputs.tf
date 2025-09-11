output "private_ips" {
  value = google_compute_instance.this_nodes.*.network_interface.0.network_ip
}

output "public_ips" {
  value = var.enable_ipv6 ? google_compute_instance.this_nodes.*.network_interface.0.ipv6_access_config.0.external_ipv6 : google_compute_instance.this_nodes.*.network_interface.0.access_config.0.nat_ip
}

output "seed_private_ips" {
  value = google_compute_instance.this_seed_nodes.*.network_interface.0.network_ip
}

output "seed_public_ips" {
  value = var.enable_ipv6 ? google_compute_instance.this_seed_nodes.*.network_interface.0.ipv6_access_config.0.external_ipv6 : google_compute_instance.this_seed_nodes.*.network_interface.0.access_config.0.nat_ip
}
