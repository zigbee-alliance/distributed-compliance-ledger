output "nodes" {
  value = local.nodes
}

output "prometheus_endpoint" {
  value = local.prometheus_endpoint
}

output "ansible_inventory" {
  value = {
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
}