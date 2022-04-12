provider "aws" {
    alias   = "reg1"
    region  = var.region_1
}

provider "aws" {
    alias  = "reg2"
    region = var.region_2
}

# Validator
module "validator" {
    source    = "./validator"
    providers = {
        aws = aws.reg1
    }
}

# Private Sentries
module "private_sentries" {
    source      = "./private-sentries"

    providers = {
        aws = aws.reg1
        aws.peer = aws.reg1
    }

    validator_vpc = module.validator.vpc
}

# Public Sentries
module "public_sentries" {
    source      = "./public-sentries"
    # node_count  = 3
    providers   = {
        aws = aws.reg1
    }
}

# Observers region 1
module "observers_region_1" {
    source      = "./observers"
    providers   = {
        aws = aws.reg1
    }
}

# Observers region 2
module "observers_region_2" {
    source      = "./observers"
    providers   = {
        aws = aws.reg2
    }
}