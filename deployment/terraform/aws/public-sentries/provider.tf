terraform {
  required_providers {
    aws = {
      source                = "hashicorp/aws"
      version               = ">= 4.1"
      configuration_aliases = [aws, aws.peer]
    }
  }
}
