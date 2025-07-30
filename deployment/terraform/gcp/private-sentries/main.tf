resource "google_compute_instance_template" "private_sentry_template" {
  labels = var.labels

  name_prefix  = "private-sentry-template"
  machine_type = var.machine_type
  region       = var.region

  disk {
    source_image = var.boot_image
    auto_delete  = true
    boot         = true
  }

  network_interface {
    subnetwork = var.subnetwork
  }

  service_account {
    email  = var.service_account_email
    scopes = ["cloud-platform"]
  }

  tags = ["private-sentry"]
}

resource "google_compute_region_instance_group_manager" "private_sentry_group" {
  name                = "private-sentry-group"
  base_instance_name  = "private-sentry"
  region              = var.region
  version {
    instance_template = google_compute_instance_template.private_sentry_template.id
  }
  target_size         = var.instance_count
}
