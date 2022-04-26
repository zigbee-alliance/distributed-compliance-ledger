variable "region_1" {
  description = "AWS Region 1"
  default     = "us-west-1"
}

variable "region_2" {
  description = "AWS Region 2"
  default     = "us-east-2"
}

variable "root_domain_name" {
  description = "Root domain name for dcl observer endpoints"
  default     = ""
}

variable "enable_tls" {
  description = "Enable tls for observer endpoints"
  default     = false
}