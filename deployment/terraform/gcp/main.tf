provider "google" {
  alias   = "region_1"
  region  = var.region_1
  project = var.project_id
}

provider "google" {
  alias   = "region_2"
  region  = var.region_2
  project = var.project_id
}

# Enable OS Login for all compute instances in the project
resource "google_compute_project_metadata" "default" {
  metadata = {
    enable-oslogin = "TRUE"
  }
}

# IAM
module "iam" {
  source = "./iam"
  providers = {
    google = google.region_1
  }
  project_id = var.project_id
}

# Validator
module "validator" {
  source = "./validator"
  providers = {
    google = google.region_1
  }

  labels = local.labels

  instance_type = var.validator_config.instance_type
  disable_instance_protection = local.disable_validator_protection
  service_account_email = module.iam.service_account_email

  boot_image = local.default_source_image
  subnetwork = ""
  project_id = var.project_id
}

# Private Sentries
module "private_sentries" {
  count = var.private_sentries_config.enable ? 1 : 0

  source = "./private-sentries"
  providers = {
    google      = google.region_1
    google.peer = google.region_1
  }

  labels = local.labels

  nodes_count           = var.private_sentries_config.nodes_count
  instance_type         = var.private_sentries_config.instance_type
  service_account_email = module.iam.service_account_email

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  peer_vpc = module.validator.vpc
  project_id = var.project_id
  region = var.region_1
}

# Public Sentries region 1
module "public_sentries_1" {
  count = (var.private_sentries_config.enable &&
    var.public_sentries_config.enable &&
  contains(var.public_sentries_config.regions, 1)) ? 1 : 0

  source = "./public-sentries"
  providers = {
    google      = google.region_1
    google.peer = google.region_1
  }

  labels = local.labels

  nodes_count           = var.public_sentries_config.nodes_count
  instance_type         = var.public_sentries_config.instance_type
  service_account_email = module.iam.service_account_email

  enable_ipv6 = var.public_sentries_config.enable_ipv6

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  region_index = 1
  peer_vpc     = module.private_sentries[0].vpc
}

# Public Sentries region 2
module "public_sentries_2" {
  count = (var.private_sentries_config.enable &&
    var.public_sentries_config.enable &&
  contains(var.public_sentries_config.regions, 2)) ? 1 : 0

  source = "./public-sentries"
  providers = {
    google      = google.region_2
    google.peer = google.region_1
  }

  labels = local.labels

  nodes_count           = var.public_sentries_config.nodes_count
  instance_type         = var.public_sentries_config.instance_type
  service_account_email = module.iam.service_account_email

  enable_ipv6 = var.public_sentries_config.enable_ipv6

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  region_index = 2
  peer_vpc     = module.private_sentries[0].vpc
}

# Observers region 1
module "observers_1" {
  count = (var.private_sentries_config.enable &&
    var.observers_config.enable &&
  contains(var.observers_config.regions, 1)) ? 1 : 0

  source = "./observers"
  providers = {
    google      = google.region_1
    google.peer = google.region_1
  }

  labels = local.labels

  nodes_count           = var.observers_config.nodes_count
  instance_type         = var.observers_config.instance_type
  service_account_email = module.iam.service_account_email

  root_domain_name = var.observers_config.root_domain_name
  enable_tls       = var.observers_config.enable_tls

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  region_index = 1
  peer_vpc     = module.private_sentries[0].vpc
}

# Observers region 2
module "observers_2" {
  count = (var.private_sentries_config.enable &&
    var.observers_config.enable &&
  contains(var.observers_config.regions, 2)) ? 1 : 0

  source = "./observers"
  providers = {
    google      = google.region_2
    google.peer = google.region_1
  }

  labels = local.labels

  nodes_count           = var.observers_config.nodes_count
  instance_type         = var.observers_config.instance_type
  service_account_email = module.iam.service_account_email

  root_domain_name = var.observers_config.root_domain_name
  enable_tls       = var.observers_config.enable_tls

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  region_index = 2
  peer_vpc     = module.private_sentries[0].vpc
}

module "prometheus" {
  count = local.prometheus_enabled ? 1 : 0

  source = "./prometheus"
  providers = {
    google = google.region_1
  }

  labels = local.labels

  instance_type = var.prometheus_config.instance_type
  endpoints     = local.prometheus_endpoints

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  vpc = module.private_sentries[0].vpc
}
