provider "aws" {
    alias   = "usw1"
    region  = "us-west-1"
}

provider "aws" {
    alias  = "euw1"
    region = "eu-west-1"
}

# Validator
module "validator" {
    source    = "./validator"
    providers = {
        aws = aws.usw1
    }
}

# Private Sentries
module "private_sentries" {
    source      = "./private-sentries"
    # node_count  = 2
    providers   = {
        aws = aws.usw1
    }
}

# Public Sentries
module "public_sentries" {
    source      = "./public-sentries"
    # node_count  = 3
    providers   = {
        aws = aws.usw1
    }
}

# Observers region 1
module "observers_region_1" {
    source      = "./observers"
    providers   = {
        aws = aws.usw1
    }
}

# Observers region 2
module "observers_region_2" {
    source      = "./observers"
    providers   = {
        aws = aws.euw1
    }
}