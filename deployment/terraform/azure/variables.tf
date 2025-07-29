variable "common_tags" {
  description = "Common tags for resources across all Azure resources."
  type        = map(string)
  default     = {}
}

variable "azure_resource_group_name" {
  description = "Name of the Azure Resource Group"
  type        = string
}

variable "azure_locations" {
  description = "List of Azure regions to deploy into (e.g. ['East US', 'West Europe'])"
  type        = list(string)
}

variable "ssh_public_key_path" {
  description = "Path to the SSH public key"
  type        = string
}

variable "ssh_private_key_path" {
  description = "Path to the SSH private key"
  type        = string
}

variable "validator_config" {
  description = "Configuration for validator instances"
  type = object({
    instance_type = string
    is_genesis    = bool
  })
}

variable "disable_validator_protection" {
  description = "Disable protection that prevents validator VMs from accidental deletion"
  type        = bool
  default     = false
}

variable "private_sentries_config" {
  description = "Configuration for private sentries"
  type = object({
    enable        = bool
    nodes_count   = number
    instance_type = string
  })
}

variable "public_sentries_config" {
  description = "Configuration for public sentries"
  type = object({
    enable        = bool
    enable_ipv6   = bool
    nodes_count   = number
    instance_type = string
    locations     = set(string)
  })
}

variable "observers_config" {
  description = "Configuration for observer nodes"
  type = object({
    enable           = bool
    nodes_count      = number
    instance_type    = string
    root_domain_name = string
    enable_tls       = bool
    locations        = set(string)
  })
}

variable "prometheus_config" {
  description = "Configuration for Prometheus monitoring instances"
  type = object({
    enable        = bool
    instance_type = string
  })
}
