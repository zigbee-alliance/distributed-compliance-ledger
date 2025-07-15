resource "google_compute_instance" "validator" {
  labels = var.labels

  name                = "dcl-validator"
  machine_type        = var.instance_type
  deletion_protection = !var.disable_instance_protection

  boot_disk {
    initialize_params {
      image = var.boot_image
    }
  }

  network_interface {
    subnetwork = var.subnetwork
    access_config {} # enables external IP
  }

  service_account {
    email  = var.service_account_email
    scopes = ["cloud-platform"]
  }

  tags = ["validator"]
}
