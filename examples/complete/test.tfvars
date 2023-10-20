vpc_name           = "vpc-ns-sd-786000"
vpc_cidr           = "10.2.0.0/16"
private_subnets    = ["10.2.1.0/24", "10.2.2.0/24", "10.2.3.0/24"]
availability_zones = ["us-east-2a", "us-east-2b", "us-east-2c"]
namespace_name     = "example75300.local"
description        = "Example Namespace creation for Service Discovery"
service_name       = "test-service70238"
tags = {
  provisioner = "Terraform"
  example     = "true"
}
