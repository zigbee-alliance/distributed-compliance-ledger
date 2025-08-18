variable "labels" {
  description = "A map of labels to add to all applicable resources"
  type        = map(string)
  default     = {}
}

variable "project_id" {
  type = string
}

variable "zone" {
  type = string
}

variable "machine_type" {
  type    = string
  default = "e2-medium"
}

variable "boot_image" {
  type    = string
  default = "debian-cloud/debian-11"
}

variable "subnetwork" {
  type = string
}

variable "service_account_email" {
  type = string
}

variable "startup_script" {
  type        = string
  description = "Startup script to install and configure Prometheus"
  default     = <<-EOT
#!/bin/bash
sudo apt-get update
sudo apt-get install -y prometheus
sudo systemctl enable prometheus
sudo systemctl start prometheus
EOT
}
