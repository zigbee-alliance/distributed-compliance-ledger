variable "location" {
  description = "Azure region (location) for all resources"
  # default   = "eastus"
}

variable "ssh_public_key_path" {
  description = "SSH public key path"
}

variable "ssh_private_key_path" {
  description = "SSH private key path"
}

variable "resource_group_name" {
  description = "Azure Resource Group name"
  # default   = "dcl-resource-group"
}

variable "validator_config" {
  type = object({
    instance_type = string
    is_genesis    = bool
  })
}

variable "private_sentries_config" {
  type = object({
    enable        = bool
    nodes_count   = number
    instance_type = string
  })

  description = "Private Sentries config"
}

variable "public_sentries_config" {
  type = object({
    enable        = bool
    enable_ipv6   = bool
    nodes_count   = number
    instance_type = string
    regions       = set(number) # для будущей поддержки мульти-регионов (см. ниже)
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
    regions          = set(number) # для будущей поддержки мульти-регионов (см. ниже)
  })

  description = "Observers config"
}

variable "prometheus_config" {
  type = object({
    enable        = bool
    instance_type = string
  })
}