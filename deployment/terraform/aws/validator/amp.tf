resource "aws_prometheus_workspace" "this_amp_workspace" {
  count = var.enable_prometheus ? 1 : 0

  alias = "amp-workspace"
  tags = {
    Name = "DCL AMP Workspace"
  }
}

data "aws_region" "current" {}

locals {
    prometheus_config = var.enable_prometheus ? templatefile("./provisioner/prometheus.tpl", 
    { 
      monitor = "validator-monitor", 
      targets = "'${aws_instance.this_node.private_ip}:26660'", 
      prometheus_endpoint = aws_prometheus_workspace.this_amp_workspace[0].prometheus_endpoint,
      region = data.aws_region.current.name
    }) : null
}
