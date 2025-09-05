# FIXME
# - no tags support

module "this_vpc" {
    source  = "terraform-google-modules/network/google"
    version = "~> 11.1"
   
    project_id = var.project_id

    network_name = "observers-vpc"

    subnets = local.subnets

    routes = [
        {
            name                   = "observers-egress-internet"
            destination_range      = "0.0.0.0/0"
            tags                   = local.egress_inet_tag 
            next_hop_internet      = "true"
        },
    ]
}
