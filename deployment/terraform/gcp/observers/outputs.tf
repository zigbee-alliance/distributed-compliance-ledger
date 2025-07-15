output "instance_group_name" {
  value = google_compute_region_instance_group_manager.observer.name
}

output "instance_ips" {
  value = {
    for k, v in google_compute_instance.observer : k => {
      external_ip = v.network_interface[0].access_config[0].nat_ip
      internal_ip = v.network_interface[0].network_ip
    }
  }
}
