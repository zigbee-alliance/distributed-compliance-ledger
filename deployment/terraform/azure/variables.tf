variable "resource_group_name" {
  description = "Azure Resource Group name"
  default = null
}

variable "resource_group_name_prefix" {
  description = "Prefix to use for Azure Resource Group names. Ignored if 'resource_group_name' is set"
  default   = "dcl-resource-group"
}

variable "common_tags" {
  description = "Common tags for resources created in AWS."
  type = object({
    project     = optional(string) # default: DCL
    environment = optional(string) # default: workspace name
    created-by  = optional(string) # e.g. email address
    purpose     = optional(string)
  })
}

variable "location_1" {
  description = "Azure location 1"
}

variable "location_2" {
  description = "Azure location 2"
}


variable "ssh_public_key_path" {
  description = "SSH public key path"
}

variable "ssh_private_key_path" {
  description = "SSH private key path"
}

variable "validator_config" {
  type = object({
    instance_size = string
    is_genesis    = bool
    enable_encryption_at_host = bool
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
    instance_size = string
  })

  description = "Private Sentries config"
}

variable "public_sentries_config" {
  type = object({
    enable        = bool
    enable_ipv6   = bool
    nodes_count   = number
    instance_size = string
    locations     = set(number)
  })

  description = "Public Sentries config"
}

variable "observers_config" {
  type = object({
    enable           = bool
    nodes_count      = number
    instance_size    = string
    root_domain_name = string
    enable_tls       = bool
    locations        = set(number)
    azs              = optional(list(set(number)))
  })

  description = "Observers config"
}

variable "prometheus_config" {
  type = object({
    enable        = bool
    instance_size = string
  })
}
