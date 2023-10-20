// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

### VPC related variables

variable "vpc_name" {
  default = "test-vpc-015935234"
}

variable "vpc_cidr" {
  default = "10.1.0.0/16"
}

variable "private_subnets" {
  description = "List of private subnet cidrs"
  default     = ["10.1.1.0/24", "10.1.2.0/24", "10.1.3.0/24"]
}

variable "availability_zones" {
  description = "List of availability zones for the VPC"
  default     = ["us-east-2a", "us-east-2b", "us-east-2c"]
}

### Namespace related variables
variable "namespace_name" {
  description = "Name of the namespace. For example example.local"
  type        = string
}

variable "description" {
  description = "Description about the namespace"
  default     = ""
  type        = string
}


## Service related variables
variable "service_name" {
  description = "Name of the service. The service will be discoverable by this name - <service_name>.<namespace_name>"
  type        = string
}


variable "dns_record_type" {
  description = "DNS record type for the service. Default is A record"
  type        = string
  default     = "A"
}

variable "ttl" {
  description = "The amount of time, in seconds, that you want DNS resolvers to cache the settings for this resource record set."
  type        = number
  default     = 10
}

variable "routing_policy" {
  description = "he routing policy that you want to apply to all records that Route 53 creates when you register an instance and specify the service. Valid Values: MULTIVALUE, WEIGHTED"
  type        = string
  default     = "MULTIVALUE"
}

variable "health_check_failure_threshold" {
  description = "The number of 30-second intervals that you want service discovery to wait before it changes the health status of a service instance. Maximum value of 10"
  type        = number
  default     = 1
}

variable "tags" {
  description = "A map of custom tags to be attached to this resource"
  type        = map(string)
  default     = {}
}
