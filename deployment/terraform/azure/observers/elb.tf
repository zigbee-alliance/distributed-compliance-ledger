resource "azurerm_lb" "this_nlb" {
  name               = "observers-network-lb"
  internal           = false
  load_balancer_type = "network"
  subnets            = module.this_vpc.public_subnets

  enable_cross_zone_load_balancing = true
  # enable_deletion_protection = true

  tags = {
    Name = "Observers NLB"
  }
}

locals {
  tls_cert_arn = var.enable_tls ? aws_acm_certificate_validation.this_acm_cert_validation[0].certificate_arn : ""
  ssl_policy   = "ELBSecurityPolicy-TLS13-1-2-2021-06" # TLS 1.3 (recommended)
}

resource "azurerm_lb_rule" "rest" {
  count = local.enable_tls ? 0 : 1

  load_balancer_arn = azurerm_lb.this_nlb.arn
  port              = "80"
  protocol          = "TCP"

  default_action {
    type             = "forward"
    backend_address_pool_id = azurerm_lb_backend_address_pool.rest.id
  }
}

resource "azurerm_lb_rule" "grpc" {
  count = local.enable_tls ? 0 : 1

  load_balancer_arn = azurerm_lb.this_nlb.arn
  port              = "9090"
  protocol          = "TCP"

  default_action {
    type             = "forward"
    backend_address_pool_id = azurerm_lb_backend_address_pool.grpc.id
  }
}

resource "azurerm_lb_rule" "rpc" {
  count = local.enable_tls ? 0 : 1

  load_balancer_arn = azurerm_lb.this_nlb.arn
  port              = "8080"
  protocol          = "TCP"

  default_action {
    type             = "forward"
    backend_address_pool_id = azurerm_lb_backend_address_pool.rpc.id
  }
}

resource "azurerm_lb_rule" "tls_rest" {
  count = local.enable_tls ? 1 : 0

  load_balancer_arn = azurerm_lb.this_nlb.arn
  port              = "443"
  protocol          = "TLS"
  certificate_arn   = local.tls_cert_arn
  ssl_policy        = local.ssl_policy

  default_action {
    type             = "forward"
    backend_address_pool_id = azurerm_lb_backend_address_pool.rest.id
  }

  depends_on = [
    aws_acm_certificate_validation.this_acm_cert_validation[0]
  ]
}

resource "azurerm_lb_rule" "tls_grpc" {
  count = local.enable_tls ? 1 : 0

  load_balancer_arn = azurerm_lb.this_nlb.arn
  port              = "8443"
  protocol          = "TLS"
  certificate_arn   = local.tls_cert_arn
  ssl_policy        = local.ssl_policy

  default_action {
    type             = "forward"
    backend_address_pool_id = azurerm_lb_backend_address_pool.grpc.id
  }

  depends_on = [
    aws_acm_certificate_validation.this_acm_cert_validation[0]
  ]
}

resource "azurerm_lb_rule" "tls_rpc" {
  count = local.enable_tls ? 1 : 0

  load_balancer_arn = azurerm_lb.this_nlb.arn
  port              = "26657"
  protocol          = "TLS"
  certificate_arn   = local.tls_cert_arn
  ssl_policy        = local.ssl_policy

  default_action {
    type             = "forward"
    backend_address_pool_id = azurerm_lb_backend_address_pool.rpc.id
  }

  depends_on = [
    aws_acm_certificate_validation.this_acm_cert_validation[0]
  ]
}

resource "azurerm_lb_backend_address_pool" "rest" {
  name               = "observers-rest-target-group"
  port               = 1317
  protocol           = "TCP"
  vpc_id             = module.this_vpc.vpc_id
  preserve_client_ip = false
}

resource "azurerm_lb_backend_address_pool" "grpc" {
  name               = "observers-grpc-target-group"
  port               = 9090
  protocol           = "TCP"
  vpc_id             = module.this_vpc.vpc_id
  preserve_client_ip = false
}

resource "azurerm_lb_backend_address_pool" "rpc" {
  name               = "observers-rpc-target-group"
  port               = 26657
  protocol           = "TCP"
  vpc_id             = module.this_vpc.vpc_id
  preserve_client_ip = false
}

resource "aws_lb_target_group_attachment" "rest_targets" {
  count = length(azurerm_instance.this_nodes)

  backend_address_pool_id = azurerm_lb_backend_address_pool.rest.id
  target_id        = azurerm_instance.this_nodes[count.index].id
  port             = 1317
}

resource "aws_lb_target_group_attachment" "grpc_targets" {
  count = length(azurerm_instance.this_nodes)

  backend_address_pool_id = azurerm_lb_backend_address_pool.grpc.id
  target_id        = azurerm_instance.this_nodes[count.index].id
  port             = 9090
}

resource "aws_lb_target_group_attachment" "rpc_targets" {
  count = length(azurerm_instance.this_nodes)

  backend_address_pool_id = azurerm_lb_backend_address_pool.rpc.id
  target_id        = azurerm_instance.this_nodes[count.index].id
  port             = 26657
}