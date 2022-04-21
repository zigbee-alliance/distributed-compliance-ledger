variable "region_1" {
    description = "AWS Region 1"
    # default     = "us-west-1"
}

variable "region_2" {
    description = "AWS Region 2"
    # default     = "us-east-2"
}

variable "private_sentries_config" {
    type = object({ 
        enable = bool, 
        nodes_count = number,
    })

    description = "Public Sentries config"
}

variable "public_sentries_config" {
    type = object({ 
        enable = bool,
        enable_ipv6 = bool,
        nodes_count = number,
        regions = set(number)
    })

    description = "Public Sentries config"
}

variable "observers_config" {
    type = object({ 
        enable = bool, 
        nodes_count = number,
        root_domain_name = string
        enable_tls = bool
        regions = set(number)
    })

    description = "Observers config"
}

variable "enable_observers_tls" {
    description = "Enable tls for observer endpoints"
    default     = true
}

variable "root_domain_name" {
    description = "Root domain name for dcl observer endpoints"
    default     = ""
}