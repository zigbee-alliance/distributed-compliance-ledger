output "prometheus_ip" {
  value = google_compute_instance.prometheus.network_interface[0].access_config[0].nat_ip
}
