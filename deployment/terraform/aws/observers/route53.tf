locals {
  enable_routing = var.root_domain_name == "" ? 0 : 1
}

data "aws_route53_zone" "this_zone" {
  count = local.enable_routing
  name  = var.root_domain_name
}

data "aws_region" "current" {}

resource "aws_route53_record" "on" {
  count = local.enable_routing

  zone_id = data.aws_route53_zone.this_zone[0].zone_id
  name    = "on.${data.aws_route53_zone.this_zone[0].name}"
  type    = "CNAME"
  ttl     = "300"

  latency_routing_policy {
    region = data.aws_region.current.name
  }

  set_identifier = "Observers NLB [${var.region_index}]"
  records        = ["${aws_lb.this_nlb.dns_name}"]
}