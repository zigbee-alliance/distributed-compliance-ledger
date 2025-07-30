output "service_account_email" {
  description = "Service account email"
  value       = google_service_account.default.email
}
