resource "google_compute_instance_template" "observer" {
  labels = var.labels

  name_prefix  = "observer-template"
  machine_type = var.machine_type

  disk {
    source_image = "ubuntu-os-pro-cloud/ubuntu-minimal-pro-2004-lts"

    auto_delete = true
    boot        = true
  }

  network_interface {
    access_config {}
  }

  service_account {
    email  = var.service_account_email
    scopes = ["cloud-platform"]
  }

  tags = ["observer"]
}

resource "google_compute_region_instance_group_manager" "observer" {
  name               = "observer-group"
  base_instance_name = "observer"
  version {
    instance_template = google_compute_instance_template.observer.id
  }
  target_size = var.instance_count
  auto_healing_policies {
    health_check      = google_compute_health_check.observer.id
    initial_delay_sec = 30
  }
}

data "google_compute_region_instance_group" "observer" {
  name = google_compute_region_instance_group_manager.observer.name
}

resource "google_compute_instance" "observer" {
  for_each     = toset(data.google_compute_region_instance_group.observer.instances)
  name         = basename(each.value)
  zone         = google_compute_region_instance_group_manager.observer.region
  project      = var.project_id
  machine_type = var.machine_type
  boot_disk {
    initialize_params {
      image = var.boot_image
    }
  }
  network_interface {
    access_config {
    }
  }
}

resource "google_compute_health_check" "observer" {
  name                = "observer-health-check"
  check_interval_sec  = 10
  timeout_sec         = 5
  healthy_threshold   = 2
  unhealthy_threshold = 3

  http_health_check {
    port = 80
  }
}
