resource "google_compute_instance" "prometheus" {
  labels = var.labels

  name         = "dcl-prometheus"
  machine_type = var.machine_type
  zone         = var.zone

  boot_disk {
    initialize_params {
      image = var.boot_image
    }
  }

  network_interface {
    subnetwork         = var.subnetwork
    access_config {}
  }

  service_account {
    email  = var.service_account_email
    scopes = ["monitoring", "logging-write", "cloud-platform"]
  }

  metadata_startup_script = var.startup_script

  tags = ["prometheus"]
}
