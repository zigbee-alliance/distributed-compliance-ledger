variable "tags" {
  description = "A map of tags to add to all resources"
  type        = map(string)
  default     = {}
}

variable "scope_id" {
  type        = string
  description = "ID области (subscription, resource group или resource) для назначения роли"
}