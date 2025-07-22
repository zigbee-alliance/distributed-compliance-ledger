resource "aws_acm_certificate" "this_acm_cert" {
  count = local.enable_tls ? 1 : 0

  domain_name       = "on.${data.aws_route53_zone.this_zone[0].name}"
  validation_method = "DNS"
}

resource "aws_route53_record" "this_acm_val_records" {
  count = local.enable_tls ? length(aws_acm_certificate.this_acm_cert[0].domain_validation_options) : 0

  name    = tolist(aws_acm_certificate.this_acm_cert[0].domain_validation_options)[count.index].resource_record_name
  records = [tolist(aws_acm_certificate.this_acm_cert[0].domain_validation_options)[count.index].resource_record_value]
  type    = tolist(aws_acm_certificate.this_acm_cert[0].domain_validation_options)[count.index].resource_record_type

  allow_overwrite = true
  ttl             = 60
  zone_id         = data.aws_route53_zone.this_zone[0].zone_id
}

resource "aws_acm_certificate_validation" "this_acm_cert_validation" {
  count = local.enable_tls ? 1 : 0

  certificate_arn         = aws_acm_certificate.this_acm_cert[0].arn
  validation_record_fqdns = aws_route53_record.this_acm_val_records[*].fqdn
}
