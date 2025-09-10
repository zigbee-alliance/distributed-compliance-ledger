provider "google" {
  alias          = "region_1"
  region         = var.region_1
  project        = var.project_id
  default_labels = local.labels
}

provider "google" {
  alias          = "region_2"
  region         = var.region_2
  project        = var.project_id
  default_labels = local.labels
}

# TODO
# Enable OS Login for all compute instances in the project
#resource "google_compute_project_metadata" "default" {
#  metadata = {
#    enable-oslogin = "TRUE"
#  }
#}

# Validator
module "validator" {
  source = "./validator"
  providers = {
    google = google.region_1
  }

  region = var.region_1
  labels = local.labels

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  instance_type               = var.validator_config.instance_type
  disable_instance_protection = local.disable_validator_protection

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

  region = var.region_1
  labels = local.labels

  nodes_count   = var.private_sentries_config.nodes_count
  instance_type = var.private_sentries_config.instance_type

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  peer_vpc   = module.validator.vpc
  project_id = var.project_id
}

# Public Sentries
module "public_sentries" {
  count = (var.private_sentries_config.enable &&
  var.public_sentries_config.enable) ? 1 : 0

  source = "./public-sentries"
  providers = {
    google      = google.region_1
    google.peer = google.region_1
  }

  region_config = [
    {
      region      = var.region_1
      nodes_count = var.public_sentries_config.region1_nodes_count
    },
    {
      region      = var.region_2
      nodes_count = var.public_sentries_config.region2_nodes_count
    }
  ]

  labels = local.labels

  instance_type = var.public_sentries_config.instance_type

  enable_ipv6 = var.public_sentries_config.enable_ipv6

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  peer_vpc   = module.private_sentries[0].vpc
  project_id = var.project_id
}

# Observers
module "observers" {
  count = (var.private_sentries_config.enable &&
  var.observers_config.enable) ? 1 : 0

  source = "./observers"

  providers = {
    google      = google.region_1
    google.peer = google.region_1
  }

  region_config = [
    {
      region      = var.region_1
      nodes_count = var.observers_config.region1_nodes_count
    },
    {
      region      = var.region_2
      nodes_count = var.observers_config.region2_nodes_count
    }
  ]

  labels = local.labels

  instance_type = var.observers_config.instance_type

  root_domain_name = var.observers_config.root_domain_name
  enable_tls       = var.observers_config.enable_tls

  ssh_public_key_path  = var.ssh_public_key_path
  ssh_private_key_path = var.ssh_private_key_path

  peer_vpc   = module.private_sentries[0].vpc
  project_id = var.project_id
}


#   module "prometheus" {
#     count = local.prometheus_enabled ? 1 : 0

#     source = "./prometheus"
#     providers = {
#       google = google.region_1
#     }

#     labels = local.labels

#     instance_type = var.prometheus_config.instance_type
#     endpoints     = local.prometheus_endpoints

#     ssh_public_key_path  = var.ssh_public_key_path
#     ssh_private_key_path = var.ssh_private_key_path

#     vpc = module.private_sentries[0].vpc
#   }
