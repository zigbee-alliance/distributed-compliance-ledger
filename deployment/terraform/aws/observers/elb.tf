resource "aws_lb" "this_nlb" {
  name               = "observers-network-lb"
  internal           = false
  load_balancer_type = "network"
  subnets            = module.this_vpc.public_subnets

  enable_cross_zone_load_balancing = true
  # enable_deletion_protection = true

  tags = merge(var.tags, {
    Name = "Observers NLB"
  })
}

locals {
  tls_cert_arn = local.enable_tls ? aws_acm_certificate_validation.this_acm_cert_validation[0].certificate_arn : ""
  ssl_policy   = "ELBSecurityPolicy-TLS13-1-2-2021-06" # TLS 1.3 (recommended)
}

resource "aws_lb_listener" "rest" {
  count = local.enable_tls ? 0 : 1

  load_balancer_arn = aws_lb.this_nlb.arn
  port              = "80"
  protocol          = "TCP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.rest.arn
  }
  tags = var.tags
}

resource "aws_lb_listener" "grpc" {
  count = local.enable_tls ? 0 : 1

  load_balancer_arn = aws_lb.this_nlb.arn
  port              = "9090"
  protocol          = "TCP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.grpc.arn
  }
  tags = var.tags
}

resource "aws_lb_listener" "rpc" {
  count = local.enable_tls ? 0 : 1

  load_balancer_arn = aws_lb.this_nlb.arn
  port              = "8080"
  protocol          = "TCP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.rpc.arn
  }
  tags = var.tags
}

resource "aws_lb_listener" "tls_rest" {
  count = local.enable_tls ? 1 : 0

  load_balancer_arn = aws_lb.this_nlb.arn
  port              = "443"
  protocol          = "TLS"
  certificate_arn   = local.tls_cert_arn
  ssl_policy        = local.ssl_policy

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.rest.arn
  }

  depends_on = [
    aws_acm_certificate_validation.this_acm_cert_validation[0]
  ]
  tags = var.tags
}

resource "aws_lb_listener" "tls_grpc" {
  count = local.enable_tls ? 1 : 0

  load_balancer_arn = aws_lb.this_nlb.arn
  port              = "8443"
  protocol          = "TLS"
  certificate_arn   = local.tls_cert_arn
  ssl_policy        = local.ssl_policy

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.grpc.arn
  }

  depends_on = [
    aws_acm_certificate_validation.this_acm_cert_validation[0]
  ]
  tags = var.tags
}

resource "aws_lb_listener" "tls_rpc" {
  count = local.enable_tls ? 1 : 0

  load_balancer_arn = aws_lb.this_nlb.arn
  port              = "26657"
  protocol          = "TLS"
  certificate_arn   = local.tls_cert_arn
  ssl_policy        = local.ssl_policy

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.rpc.arn
  }

  depends_on = [
    aws_acm_certificate_validation.this_acm_cert_validation[0]
  ]
  tags = var.tags
}

resource "aws_lb_target_group" "rest" {
  name               = "observers-rest-target-group"
  port               = 1317
  protocol           = "TCP"
  vpc_id             = module.this_vpc.vpc_id
  preserve_client_ip = false
  tags               = var.tags
}

resource "aws_lb_target_group" "grpc" {
  name               = "observers-grpc-target-group"
  port               = 9090
  protocol           = "TCP"
  vpc_id             = module.this_vpc.vpc_id
  preserve_client_ip = false
  tags               = var.tags
}

resource "aws_lb_target_group" "rpc" {
  name               = "observers-rpc-target-group"
  port               = 26657
  protocol           = "TCP"
  vpc_id             = module.this_vpc.vpc_id
  preserve_client_ip = false
  tags               = var.tags
}

resource "aws_lb_target_group_attachment" "rest_targets" {
  count = length(aws_instance.this_nodes)

  target_group_arn = aws_lb_target_group.rest.arn
  target_id        = aws_instance.this_nodes[count.index].id
  port             = 1317
}

resource "aws_lb_target_group_attachment" "grpc_targets" {
  count = length(aws_instance.this_nodes)

  target_group_arn = aws_lb_target_group.grpc.arn
  target_id        = aws_instance.this_nodes[count.index].id
  port             = 9090
}

resource "aws_lb_target_group_attachment" "rpc_targets" {
  count = length(aws_instance.this_nodes)

  target_group_arn = aws_lb_target_group.rpc.arn
  target_id        = aws_instance.this_nodes[count.index].id
  port             = 26657
}
