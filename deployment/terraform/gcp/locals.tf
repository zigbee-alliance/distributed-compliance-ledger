locals {

  project_name_default = "dcl"

  base_labels = {
    project     = local.project_name_default
    environment = terraform.workspace
  }

  disable_validator_protection = tobool(var.disable_validator_protection) == true

  labels = merge(local.base_labels, { for k, v in var.common_labels : k => v if try(length(v), 0) > 0 })

  nodes = {
    validator = {
      private_ips = module.validator.private_ips
      public_ips  = module.validator.public_ips
    }

    private_sentries = {
      private_ips       = var.private_sentries_config.enable ? module.private_sentries[0].private_ips : []
      public_ips        = var.private_sentries_config.enable ? module.private_sentries[0].public_ips : []
      public_static_ips = var.private_sentries_config.enable ? module.private_sentries[0].public_static_ips : []
    }

    public_sentries = {
      private_ips = (var.private_sentries_config.enable && var.public_sentries_config.enable) ? module.public_sentries[0].private_ips : [],
      public_ips  = (var.private_sentries_config.enable && var.public_sentries_config.enable) ? module.public_sentries[0].public_ips : [],
    }

    seeds = {
      private_ips = (var.private_sentries_config.enable && var.public_sentries_config.enable) ? module.public_sentries[0].seed_private_ips : []
      public_ips  = (var.private_sentries_config.enable && var.public_sentries_config.enable) ? module.public_sentries[0].seed_public_ips : []
    }

    observers = {
      private_ips = (var.private_sentries_config.enable && var.observers_config.enable) ? module.observers[0].private_ips : []
      public_ips  = (var.private_sentries_config.enable && var.observers_config.enable) ? module.observers[0].public_ips : []
    }
  }

  ansible_inventory = {
    all = {
      children = {
        genesis = {
          hosts = var.validator_config.is_genesis ? { for host in local.nodes.validator.public_ips : host => null } : null
        }

        validators = {
          hosts = { for host in local.nodes.validator.public_ips : host => null }
        }

        private_sentries = {
          hosts = { for host in local.nodes.private_sentries.public_ips : host => null }
        }

        public_sentries = {
          hosts = { for host in local.nodes.public_sentries.public_ips : host => null }
        }

        seeds = {
          hosts = { for host in local.nodes.seeds.public_ips : host => null }
        }

        observers = {
          hosts = { for host in local.nodes.observers.public_ips : host => null }
        }
      }
    }
  }

  # TODO metrics collection configuration
  # prometheus_endpoints = concat(
  #   local.nodes.validator.private_ips,
  #   local.nodes.private_sentries.private_ips,
  #   local.nodes.public_sentries.private_ips,
  #   local.nodes.observers.private_ips,
  #   local.nodes.seeds.private_ips
  # )

  # prometheus_enabled = var.private_sentries_config.enable && var.prometheus_config.enable

  # prometheus_endpoint = local.prometheus_enabled ? module.prometheus[0].prometheus_endpoint : null
}
