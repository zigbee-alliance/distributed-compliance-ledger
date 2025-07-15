variable "labels" {
  description = "A map of labels to add to all applicable resources"
  type        = map(string)
  default     = {}
}

variable "disable_instance_protection" {
  description = "Disable the protection that prevents the validator instance from being accidentally terminated"
  type        = bool
  default     = false
}

variable "project_id" {
  type = string
}

variable "instance_type" {
  type = string
}

variable "boot_image" {
  description = "Image to use for the boot disk"
  type        = string
  default     = local.default_source_image
}

variable "subnetwork" {
  description = "The subnetwork to deploy the instance into"
  type        = string
}

variable "service_account_email" {
  description = "IAM service account email"
  type        = string
}
