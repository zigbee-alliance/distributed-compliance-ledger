locals {
  nodes = {
    validator = {
      private_ips = module.validator.private_ips
      public_ips  = module.validator.public_ips
    },

    private_sentries = {
      private_ips = var.private_sentries_config.enable ? module.private_sentries[0].private_ips : []
      public_ips  = var.private_sentries_config.enable ? module.private_sentries[0].public_ips : []
      public_eips = var.private_sentries_config.enable ? module.private_sentries[0].public_eips : []
    },

    public_sentries = {
      private_ips = concat(
        (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.regions, 1)) ? module.public_sentries_1[0].private_ips : [],
        (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.regions, 2)) ? module.public_sentries_2[0].private_ips : [],
      )

      public_ips = concat(
        (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.regions, 1)) ? module.public_sentries_1[0].public_ips : [],
        (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.regions, 2)) ? module.public_sentries_2[0].public_ips : [],
      )
    },

    seeds = {
      private_ips = concat(
        (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.regions, 1)) ? module.public_sentries_1[0].seed_private_ips : [],
        (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.regions, 2)) ? module.public_sentries_2[0].seed_private_ips : [],
      )

      public_ips = concat(
        (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.regions, 1)) ? module.public_sentries_1[0].seed_public_ips : [],
        (var.private_sentries_config.enable && var.public_sentries_config.enable && contains(var.public_sentries_config.regions, 2)) ? module.public_sentries_2[0].seed_public_ips : [],
      )
    },

    observers = {
      private_ips = concat(
        (var.private_sentries_config.enable && var.observers_config.enable && contains(var.observers_config.regions, 1)) ? module.observers_1[0].private_ips : [],
        (var.private_sentries_config.enable && var.observers_config.enable && contains(var.observers_config.regions, 2)) ? module.observers_2[0].private_ips : [],
      )

      public_ips = concat(
        (var.private_sentries_config.enable && var.observers_config.enable && contains(var.observers_config.regions, 1)) ? module.observers_1[0].public_ips : [],
        (var.private_sentries_config.enable && var.observers_config.enable && contains(var.observers_config.regions, 2)) ? module.observers_2[0].public_ips : [],
      )
    },
  }

  prometheus_endpoints = concat(local.nodes.validator.private_ips, local.nodes.private_sentries.private_ips, local.nodes.public_sentries.private_ips, local.nodes.observers.private_ips, local.nodes.seeds.private_ips)

  prometheus_endpoint = var.prometheus_config.enable ? module.prometheus[0].prometheus_endpoint : null
}