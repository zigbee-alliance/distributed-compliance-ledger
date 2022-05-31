variable "region_1" {
  description = "AWS Region 1"
  # default     = "us-west-1"
}

variable "region_2" {
  description = "AWS Region 2"
  # default     = "us-east-2"
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