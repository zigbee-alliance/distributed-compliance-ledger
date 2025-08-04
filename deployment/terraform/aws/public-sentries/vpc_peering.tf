module "this_vpc_peering" {
  source  = "grem11n/vpc-peering/aws"
  version = "4.1.0"

  providers = {
    aws.this = aws
    aws.peer = aws.peer
  }

  this_vpc_id = module.this_vpc.vpc_id
  peer_vpc_id = var.peer_vpc.vpc_id

  this_rts_ids = [element(module.this_vpc.public_route_table_ids, 0)]
  peer_rts_ids = [element(var.peer_vpc.public_route_table_ids, 0)]

  auto_accept_peering = true

  tags = merge(var.tags, {
    Name = "Public Sentries to Private Sentries VPC peering"
  })
}
