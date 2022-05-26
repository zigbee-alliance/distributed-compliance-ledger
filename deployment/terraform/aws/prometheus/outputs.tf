output "prometheus_endpoint" {
  value = aws_prometheus_workspace.this_amp_workspace.prometheus_endpoint
}