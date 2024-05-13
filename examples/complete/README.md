# Service Discovery Service setup example
This module will provision a Service Discovery service in Cloud Map. It will provision all dependent resources as shown below
- VPC
- Cloud Map Namespace

## Pre-requisites
- Need to create a `provider.tf` with below contents
    ```
    provider "aws" {
      profile = "<profile-name>"
      region  = "<aws-region>"
    }
    ```
- Command to execute the module
  `terraform apply -var-file test.tfvars`
<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.5.0, <= 1.5.5 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_service"></a> [service](#module\_service) | ../.. | n/a |
| <a name="module_vpc"></a> [vpc](#module\_vpc) | terraform-aws-modules/vpc/aws | ~> 5.1.1 |
| <a name="module_namespace"></a> [namespace](#module\_namespace) | git::https://github.com/launchbynttdata/tf-aws-module_primitive-private_dns_namespace.git | 1.0.0 |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_vpc_name"></a> [vpc\_name](#input\_vpc\_name) | n/a | `string` | `"test-vpc-015935234"` | no |
| <a name="input_vpc_cidr"></a> [vpc\_cidr](#input\_vpc\_cidr) | n/a | `string` | `"10.1.0.0/16"` | no |
| <a name="input_private_subnets"></a> [private\_subnets](#input\_private\_subnets) | List of private subnet cidrs | `list(string)` | <pre>[<br>  "10.1.1.0/24",<br>  "10.1.2.0/24",<br>  "10.1.3.0/24"<br>]</pre> | no |
| <a name="input_availability_zones"></a> [availability\_zones](#input\_availability\_zones) | List of availability zones for the VPC | `list(string)` | <pre>[<br>  "us-east-2a",<br>  "us-east-2b",<br>  "us-east-2c"<br>]</pre> | no |
| <a name="input_namespace_name"></a> [namespace\_name](#input\_namespace\_name) | Name of the namespace. For example example.local | `string` | n/a | yes |
| <a name="input_description"></a> [description](#input\_description) | Description about the namespace | `string` | `""` | no |
| <a name="input_service_name"></a> [service\_name](#input\_service\_name) | Name of the service. The service will be discoverable by this name - <service\_name>.<namespace\_name> | `string` | n/a | yes |
| <a name="input_dns_record_type"></a> [dns\_record\_type](#input\_dns\_record\_type) | DNS record type for the service. Default is A record | `string` | `"A"` | no |
| <a name="input_ttl"></a> [ttl](#input\_ttl) | The amount of time, in seconds, that you want DNS resolvers to cache the settings for this resource record set. | `number` | `10` | no |
| <a name="input_routing_policy"></a> [routing\_policy](#input\_routing\_policy) | he routing policy that you want to apply to all records that Route 53 creates when you register an instance and specify the service. Valid Values: MULTIVALUE, WEIGHTED | `string` | `"MULTIVALUE"` | no |
| <a name="input_health_check_failure_threshold"></a> [health\_check\_failure\_threshold](#input\_health\_check\_failure\_threshold) | The number of 30-second intervals that you want service discovery to wait before it changes the health status of a service instance. Maximum value of 10 | `number` | `1` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of custom tags to be attached to this resource | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | ID of the namespace |
| <a name="output_arn"></a> [arn](#output\_arn) | ARN of the namespace |
| <a name="output_hosted_zone"></a> [hosted\_zone](#output\_hosted\_zone) | Hosted Zone for the namespace |
| <a name="output_vpc_id"></a> [vpc\_id](#output\_vpc\_id) | VPC ID |
| <a name="output_service_id"></a> [service\_id](#output\_service\_id) | Service Discovery Service ID |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
