# TODO is it possible to get from provider ???
variable "region" {
  type = string
}

variable "project_id" {
  description = "GCP project ID"
  type        = string
}

variable "os_family" {
  description = "Node base image family"
  type        = string
  default     = "ubuntu-2004-lts" # TODO ubuntu 20.04 is deprecated
}

variable "labels" {
  description = "A map of labels to add to all applicable resources"
  type        = map(string)
  default     = {}
}

variable "ssh_public_key_path" {
  description = "SSH public key file path"
  default     = "~/.ssh/id_rsa.pub"
}

variable "ssh_private_key_path" {
  description = "SSH private key file path"
  default     = "~/.ssh/id_rsa"
}

variable "ssh_username" {
  description = "SSH username"
  default     = "ubuntu"
}

variable "peer_vpc" {
  description = "Peer VPC"
}

variable "nodes_count" {
  description = "Number of Private Sentry nodes"
}

variable "instance_type" {
  description = "Type of GCP compute instances"
}
