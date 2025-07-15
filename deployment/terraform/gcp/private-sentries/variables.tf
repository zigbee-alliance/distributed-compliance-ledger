variable "labels" {
  description = "A map of labels to add to all applicable resources"
  type        = map(string)
  default     = {}
}

variable "project_id" {
  type = string
}

variable "region" {
  type = string
}

variable "subnetwork" {
  type = string
}

variable "service_account_email" {
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

variable "instance_count" {
  type    = number
  default = 2
}
