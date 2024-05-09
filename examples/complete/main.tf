module "service" {
  source = "../.."

  name                           = var.service_name
  namespace_id                   = module.namespace.id
  ttl                            = var.ttl
  dns_record_type                = var.dns_record_type
  routing_policy                 = var.routing_policy
  health_check_failure_threshold = var.health_check_failure_threshold

  tags = var.tags
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 5.1.1"

  name                 = var.vpc_name
  cidr                 = var.vpc_cidr
  private_subnets      = var.private_subnets
  azs                  = var.availability_zones
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = var.tags
}

module "namespace" {
  source = "git::https://github.com/launchbynttdata/tf-aws-module_primitive-private_dns_namespace.git?ref=1.0.0"

  vpc_id      = module.vpc.vpc_id
  name        = var.namespace_name
  description = var.description
  tags        = var.tags
}
