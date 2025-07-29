locals {
  # Default project name and base tags
  project_name_default = "DCL"
  base_tags = {
    project     = local.project_name_default
    environment = terraform.workspace
  }

  # Validator protection flag
  disable_validator_protection = var.disable_validator_protection

  # Merge any user‑supplied common tags
  tags = merge(
    local.base_tags,
    var.common_tags
  )

  # Gather IPs from each module (Azure)
  nodes = {
    validator = {
      private_ips = module.validator.private_ips
      public_ips  = module.validator.public_ips
    }

    private_sentries = var.private_sentries_config.enable ? {
      private_ips = module.private_sentries.private_ips
      public_ips  = module.private_sentries.public_ips
      public_eips = module.private_sentries.public_eips
    } : {
      private_ips = []
      public_ips  = []
      public_eips = []
    }

    public_sentries = var.public_sentries_config.enable ? {
      private_ips = module.public_sentries.private_ips
      public_ips  = module.public_sentries.public_ips
    } : {
      private_ips = []
      public_ips  = []
    }

    seeds = var.public_sentries_config.enable ? {
      private_ips = module.public_sentries.seed_private_ips
      public_ips  = module.public_sentries.seed_public_ips
    } : {
      private_ips = []
      public_ips  = []
    }

   # observers = var.observers_config.enable ? {
   #   private_ips = module.observers.private_ips
   #   public_ips  = module.observers.public_ips
   # } : {
   #   private_ips = []
   #   public_ips  = []
   # }
  }

  # Build Ansible inventory structure
  ansible_inventory = {
    all = {
      children = {
        genesis = {
          hosts = var.validator_config.is_genesis
            ? { for ip in local.nodes.validator.public_ips : ip => null }
            : {}
        }
        validators        = { hosts = { for ip in local.nodes.validator.public_ips       : ip => null } }
        private_sentries  = { hosts = { for ip in local.nodes.private_sentries.public_ips : ip => null } }
        public_sentries   = { hosts = { for ip in local.nodes.public_sentries.public_ips  : ip => null } }
        seeds             = { hosts = { for ip in local.nodes.seeds.public_ips            : ip => null } }
       # observers         = { hosts = { for ip in local.nodes.observers.public_ips        : ip => null } }
      }
    }
  }

  # Prometheus endpoints and enable flag
  prometheus_endpoints = concat(
    local.nodes.validator.private_ips,
    local.nodes.private_sentries.private_ips,
    local.nodes.public_sentries.private_ips,
   # local.nodes.observers.private_ips,
    local.nodes.seeds.private_ips,
  )

  prometheus_enabled  = var.private_sentries_config.enable && var.prometheus_config.enable
  prometheus_endpoint = local.prometheus_enabled
    ? module.prometheus.prometheus_endpoint
    : null
}
