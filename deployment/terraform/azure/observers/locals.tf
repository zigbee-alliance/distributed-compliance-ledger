locals {
  enable_tls = var.enable_tls && var.root_domain_name != ""
}