output "vpc" {
  value = module.this_vpc
}

output "private_ips" {
  value = [google_compute_instance.this_node.network_interface.0.network_ip]
}

output "public_ips" {
  value = [google_compute_instance.this_node.network_interface.0.access_config.0.nat_ip]
}
