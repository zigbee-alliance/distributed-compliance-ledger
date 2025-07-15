variable "labels" {
  description = "A map of labels to add to all applicable resources"
  type        = map(string)
  default     = {}
}

variable "project_id" {
  description = "The GCP Project ID"
  type = string
}

variable "machine_type" {
  type    = string
}

variable "instance_count" {
  type    = number
  default = 2
}

variable "boot_image" {
  description = "Image to use for the boot disk"
  type        = string
  default     = local.default_source_image
}

variable "region_index" {
  description = "Observer Region Index"
}

variable "enable_tls" {
  description = "Enable TLS on LB listeners"
}

variable "root_domain_name" {
  description = "Root domain name"
}

variable "peer_vpc" {
  description = "Peer VPC"
}

variable "service_account_email" {
  type = string
}
