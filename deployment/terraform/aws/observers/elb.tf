resource "aws_lb" "this_nlb" {
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

resource "aws_lb_listener" "rest" {
    load_balancer_arn = aws_lb.this_nlb.arn
    port              = "80"
    protocol          = "TCP"
    # certificate_arn   = "arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4"
    # alpn_policy       = "HTTP2Preferred"

    default_action {
        type = "forward"
        target_group_arn = aws_lb_target_group.rest.arn
    }
}

resource "aws_lb_listener" "grpc" {
    load_balancer_arn = aws_lb.this_nlb.arn
    port              = "9090"
    protocol          = "TCP"
    # certificate_arn   = "arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4"
    # alpn_policy       = "HTTP2Preferred"

    default_action {
        type = "forward"
        target_group_arn = aws_lb_target_group.grpc.arn
    }
}

resource "aws_lb_listener" "rpc" {
    load_balancer_arn = aws_lb.this_nlb.arn
    port              = "26657"
    protocol          = "TCP"
    # certificate_arn   = "arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4"
    # alpn_policy       = "HTTP2Preferred"

    default_action {
        type = "forward"
        target_group_arn = aws_lb_target_group.rpc.arn
    }
}

resource "aws_lb_target_group" "rest" {
    name     = "observers-rest-target-group"
    port     = 1317
    protocol = "TCP"
    vpc_id   = module.this_vpc.vpc_id
    preserve_client_ip = false
}

resource "aws_lb_target_group" "grpc" {
    name     = "observers-grpc-target-group"
    port     = 9090
    protocol = "TCP"
    vpc_id   = module.this_vpc.vpc_id
    preserve_client_ip = false
}

resource "aws_lb_target_group" "rpc" {
    name     = "observers-rpc-target-group"
    port     = 26657
    protocol = "TCP"
    vpc_id   = module.this_vpc.vpc_id
    preserve_client_ip = false
}

resource "aws_lb_target_group_attachment" "rest_targets" {
    count = length(aws_instance.this_nodes)

    target_group_arn = aws_lb_target_group.rest.arn
    target_id        = aws_instance.this_nodes[count.index].id
    port             = 80
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