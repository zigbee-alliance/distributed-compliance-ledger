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
      private_ips = module.private_sentries[0].private_ips
      public_ips  = module.private_sentries[0].public_ips
      public_eips = module.private_sentries[0].public_eips
    } : {
      private_ips = []
      public_ips  = []
      public_eips = []
    }

    public_sentries = var.public_sentries_config.enable ? {
      private_ips = concat(
        length(module.public_sentries_1) > 0 ? module.public_sentries_1[0].private_ips : [],
        length(module.public_sentries_2) > 0 ? module.public_sentries_2[0].private_ips : []
      )
      public_ips = concat(
        length(module.public_sentries_1) > 0 ? module.public_sentries_1[0].public_ips : [],
        length(module.public_sentries_2) > 0 ? module.public_sentries_2[0].public_ips : []
      )
    } : {
      private_ips = []
      public_ips  = []
    }

    observers = var.observers_config.enable ? {
      private_ips = concat(
        length(module.observers_1) > 0 ? module.observers_1[0].private_ips : [],
        length(module.observers_2) > 0 ? module.observers_2[0].private_ips : []
      )
      public_ips = concat(
        length(module.observers_1) > 0 ? module.observers_1[0].public_ips : [],
        length(module.observers_2) > 0 ? module.observers_2[0].public_ips : []
      )
    } : {
      private_ips = []
      public_ips  = []
    }
  }

  # Build Ansible inventory structure
  ansible_inventory = {
    all = {
      children = {
        genesis = {
          hosts = var.validator_config.is_genesis ? local.nodes.validator.public_ips : []
        }
        validators = {
          hosts = local.nodes.validator.public_ips
        }
        private_sentries = {
          hosts = local.nodes.private_sentries.public_ips
        }
        public_sentries = {
          hosts = local.nodes.public_sentries.public_ips
        }
        observers = {
          hosts = local.nodes.observers.public_ips
        }
      }
    }
  }

  # Prometheus endpoints and enable flag
  prometheus_endpoints = concat(
    local.nodes.validator.private_ips,
    local.nodes.private_sentries.private_ips,
    local.nodes.public_sentries.private_ips,
    local.nodes.observers.private_ips
  )

  prometheus_enabled = var.prometheus_config.enable
  
  prometheus_endpoint = local.prometheus_enabled && length(module.prometheus) > 0 ? module.prometheus[0].prometheus_endpoint : null
}
