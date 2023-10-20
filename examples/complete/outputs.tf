output "id" {
  description = "ID of the namespace"
  value       = module.namespace.id
}

output "arn" {
  description = "ARN of the namespace"
  value       = module.namespace.arn
}

output "hosted_zone" {
  description = "Hosted Zone for the namespace"
  value       = module.namespace.hosted_zone
}

output "vpc_id" {
  description = "VPC ID"
  value       = module.vpc.vpc_id
}

output "service_id" {
  description = "Service Discovery Service ID"
  value       = module.service.id
}
