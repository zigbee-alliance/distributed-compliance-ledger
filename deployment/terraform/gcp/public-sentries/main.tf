resource "google_compute_instance_template" "sentry_template" {
  labels = var.labels

  name_prefix  = "sentry-template"
  machine_type = var.machine_type
  region       = var.region

  disk {
    source_image = var.boot_image
    auto_delete  = true
    boot         = true
  }

  network_interface {
    access_config {}
  }

  service_account {
    email  = var.service_account_email
    scopes = ["cloud-platform"]
  }

  tags = ["public-sentry"]
}

resource "google_compute_region_instance_group_manager" "sentry_group" {
  name                = "sentry-group"
  base_instance_name  = "sentry"
  region              = var.region
  version {
    instance_template = google_compute_instance_template.sentry_template.id
  }
  target_size         = var.instance_count
  auto_healing_policies {
    health_check      = google_compute_health_check.default.id
    initial_delay_sec = 30
  }
}

resource "google_compute_health_check" "default" {
  name               = "sentry-health-check"
  check_interval_sec = 10
  timeout_sec        = 5
  healthy_threshold  = 2
  unhealthy_threshold = 3

  http_health_check {
    port = 80
  }
}
