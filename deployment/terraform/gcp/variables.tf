variable "common_labels" {
  description = "Common tags for resources created in AWS."
  type = object({
    project     = optional(string) # default: DCL
    environment = optional(string) # default: workspace name
    created-by  = optional(string) # e.g. email address # FIXME @ is not allowed for labels
    purpose     = optional(string)
  })
}

variable "project_id" {
  description = "GCP project ID"
  type        = string
  default     = "DCL"
}

# FIXME default regions

variable "region_1" {
  type    = string
  default = "us-east1"
}

variable "region_2" {
  type    = string
  default = "us-west1"
}

# FIXME
variable "zone" {
  type    = string
  default = "us-east1-b"
}

variable "ssh_public_key_path" {
  description = "SSH public key path"
}

variable "ssh_private_key_path" {
  description = "SSH private key path"
}

# FIXME
#variable "subnetwork" {
#  type = string
#}

# FIXME
#variable "service_account_email" {
#  type = string
#}

variable "validator_config" {
  type = object({
    instance_type = string
    is_genesis    = bool
  })
}

variable "disable_validator_protection" {
  description = "Disable the protection that prevents the validator instance from being accidentally terminated"
  type        = bool
  default     = false
}

variable "private_sentries_config" {
  type = object({
    enable        = bool
    nodes_count   = number
    instance_type = string
  })

  description = "Public Sentries config"
}

variable "public_sentries_config" {
  type = object({
    enable        = bool
    enable_ipv6   = bool
    nodes_count   = number
    instance_type = string
    regions       = set(number)
  })

  description = "Public Sentries config"
}

variable "observers_config" {
  type = object({
    enable           = bool
    nodes_count      = number
    instance_type    = string
    root_domain_name = string
    enable_tls       = bool
    regions          = set(number)
  })

  description = "Observers config"
}

variable "prometheus_config" {
  type = object({
    enable        = bool
    instance_type = string
  })
}
