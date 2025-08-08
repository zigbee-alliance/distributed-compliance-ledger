# FIXME
# - no tags support

module "this_vpc" {
    source  = "terraform-google-modules/network/google"
    version = "~> 11.1"
   
    project_id = var.project_id

    network_name = "private-sentries-vpc"
    routing_mode = "REGIONAL" // TODO for now keeps similar to AWS

    subnets = [
        {
            subnet_name           = "${local.subnet_name}"
            subnet_ip             = "${local.vpc_network_prefix}.1.0/24"
            subnet_region         = "${local.subnet_region}"
        },
    ]

    routes = [
        {
            name                   = "private-sentries-egress-internet"
            destination_range      = "0.0.0.0/0"
            tags                   = local.egress_inet_tag
            next_hop_internet      = "true"
        },
    ]
}
