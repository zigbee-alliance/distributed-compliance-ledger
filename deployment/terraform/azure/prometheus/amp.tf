resource "aws_prometheus_workspace" "this_amp_workspace" {
  alias = "amp-workspace"
  tags = {
    Name = "DCL AMP Workspace"
  }
}

data "aws_region" "current" {}

locals {
  prometheus_config = templatefile("./provisioner/prometheus.tpl",
    {
      monitor             = "dcl-monitor",
      targets             = "${join(",", formatlist("'%s'", var.endpoints))}",
      prometheus_endpoint = azurerm_prometheus_workspace.this_amp_workspace.prometheus_endpoint,
      region              = data.aws_region.current.name
  })
}
