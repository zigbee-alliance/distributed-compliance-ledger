# reserved IP address
resource "google_compute_global_address" "this_static_ip" {
  name = "observers-nlb-static-ip"
}


# FIXME EXTERNAL vs EXTERNAL_MANAGED and other properties
# backend service GRPC
resource "google_compute_backend_service" "this_nlb_backend_service" {
  count = length(local.nodes) > 0 ? length(local.nlb_ports) : 0

  name                  = "observer-nlb-backend-service-${local.nlb_ports[count.index].name}"
  protocol              = "SSL"
  port_name             = local.nlb_ports[count.index].name
  load_balancing_scheme = "EXTERNAL"
  timeout_sec           = 10 # FIXME
  health_checks         = [google_compute_health_check.default[count.index].id]

  dynamic "backend" {
    for_each = range(length(local.zones))
    content {
      group           = google_compute_instance_group.this_nodes_group[backend.value].id
      balancing_mode  = "UTILIZATION"
      max_utilization = 1.0
      capacity_scaler = 1.0
    }
  }

}

resource "google_compute_target_ssl_proxy" "this_nlb_target_ssl_proxy" {
  count = local.enable_tls ? length(local.nlb_ports) : 0

  name             = "observer-nlb-target-ssl-proxy-${local.nlb_ports[count.index].name}"
  backend_service  = google_compute_backend_service.this_nlb_backend_service[count.index].id
  ssl_certificates = [null] # [google_compute_ssl_certificate.this_cert[0].id] TODO need to implement support for TLS
}

resource "google_compute_target_tcp_proxy" "this_nlb_target_tcp_proxy" {
  count = local.enable_tls ? 0 : length(local.nlb_ports)

  name            = "observer-nlb-target-tcp-proxy-${local.nlb_ports[count.index].name}"
  backend_service = google_compute_backend_service.this_nlb_backend_service[count.index].id
}


# FIXME properties
resource "google_compute_health_check" "default" {
  count = length(local.nlb_ports)

  name                = "observer-nlb-health-check-${local.nlb_ports[count.index].name}"
  timeout_sec         = 10
  check_interval_sec  = 30
  healthy_threshold   = 5
  unhealthy_threshold = 2

  tcp_health_check {
    port = local.nlb_ports[count.index].port
  }
}

# forwarding rule
resource "google_compute_global_forwarding_rule" "this_nlb_global_forwarding_rule" {
  count = length(local.nlb_ports)

  name                  = "observer-nlb-forwarding-rule-${local.nlb_ports[count.index].name}"
  ip_protocol           = "TCP"
  load_balancing_scheme = "EXTERNAL"
  port_range = (
    local.enable_tls
    ? local.nlb_ports[count.index].listen_port_tls
    : local.nlb_ports[count.index].listen_port
  )
  target = (
    local.enable_tls
    ? google_compute_target_ssl_proxy.this_nlb_target_ssl_proxy[count.index].id
    : google_compute_target_tcp_proxy.this_nlb_target_tcp_proxy[count.index].id
  )
  ip_address = google_compute_global_address.this_static_ip.id
}
