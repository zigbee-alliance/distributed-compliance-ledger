output "nodes" {
  value = local.nodes
}

output "common_labels" {
  value = local.labels
}

# FIXME
#output "prometheus_endpoint" {
#  value = local.prometheus_endpoint
#}

output "ansible_inventory" {
  value = local.ansible_inventory
}

output "ansible_inventory_yaml" {
  value = yamlencode(local.ansible_inventory)
}
