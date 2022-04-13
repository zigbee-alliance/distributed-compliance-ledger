provider "aws" {
    alias   = "region_1"
    region  = var.region_1
}

provider "aws" {
    alias  = "region_2"
    region = var.region_2
}

# Validator
module "validator" {
    source    = "./validator"
    providers = {
        aws = aws.region_1
    }
}

# Private Sentries
module "private_sentries" {
    source      = "./private-sentries"

    providers = {
        aws = aws.region_1
        aws.peer = aws.region_1
    }

    peer_vpc = module.validator.vpc
}

# Public Sentries
module "public_sentries" {
    source      = "./public-sentries"
    # node_count  = 3
    providers = {
        aws = aws.region_2
        aws.peer = aws.region_1
    }

    peer_vpc = module.private_sentries.vpc
}

# Observers region 1
module "observers_1" {
    source      = "./observers"
    
    providers = {
        aws = aws.region_1
        aws.peer = aws.region_1
    }

    region_index = 1
    peer_vpc = module.private_sentries.vpc
}

# Observers region 2
module "observers_2" {
    source      = "./observers"
    
    providers = {
        aws = aws.region_2
        aws.peer = aws.region_1
    }

    region_index = 2
    peer_vpc = module.private_sentries.vpc
}