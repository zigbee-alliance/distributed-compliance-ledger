locals {
  project_name_default = "DCL"

  base_tags = {
    project     = local.project_name_default
    environment = terraform.workspace
  }

  locations = [var.location_1, var.location_2]

  observers_azs = var.observers_config.azs == null ? [null, null] : var.observers_config.azs

  resource_group_names = (
    var.resource_group_name != null
    ? [var.resource_group_name, var.resource_group_name]
    : azurerm_resource_group.dcl.*.name
  )

  disable_validator_protection = tobool(var.disable_validator_protection) == true

  tags = merge(local.base_tags, { for k, v in var.common_tags : k => v if try(length(v), 0) > 0 })

  nodes = {
    validator = {
      private_ips = module.validator.private_ips
      public_ips  = module.validator.public_ips
    }

    private_sentries = {
      private_ips = var.private_sentries_config.enable ? module.private_sentries[0].private_ips : []
      public_ips  = var.private_sentries_config.enable ? module.private_sentries[0].public_ips : []
    }

    public_sentries = {
      private_ips = concat(
        (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.locations, 1)) ? module.public_sentries_1[0].private_ips : [],
        (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.locations, 2)) ? module.public_sentries_2[0].private_ips : [],
      )

      public_ips = concat(
        (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.locations, 1)) ? module.public_sentries_1[0].public_ips : [],
        (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.locations, 2)) ? module.public_sentries_2[0].public_ips : [],
      )
    }

    seeds = {
      private_ips = concat(
        (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.locations, 1)) ? module.public_sentries_1[0].seed_private_ips : [],
        (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.locations, 2)) ? module.public_sentries_2[0].seed_private_ips : [],
      )

      public_ips = concat(
        (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.locations, 1)) ? module.public_sentries_1[0].seed_public_ips : [],
        (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.locations, 2)) ? module.public_sentries_2[0].seed_public_ips : [],
      )
    }

    observers = {
      private_ips = concat(
        (var.private_sentries_config.enable && var.observers_config.enable && contains(var.observers_config.locations, 1)) ? module.observers_1[0].private_ips : [],
        (var.private_sentries_config.enable && var.observers_config.enable && contains(var.observers_config.locations, 2)) ? module.observers_2[0].private_ips : [],
      )

      public_ips = concat(
        (var.private_sentries_config.enable && var.observers_config.enable && contains(var.observers_config.locations, 1)) ? module.observers_1[0].public_ips : [],
        (var.private_sentries_config.enable && var.observers_config.enable && contains(var.observers_config.locations, 2)) ? module.observers_2[0].public_ips : [],
      )
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

  # TODO
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
