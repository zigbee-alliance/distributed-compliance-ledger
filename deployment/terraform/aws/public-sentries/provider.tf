terraform {
    required_providers {
        aws = {
            source                = "hashicorp/aws"
            version               = ">= 3.72"
            configuration_aliases = [aws, aws.peer]
        }
    }
}