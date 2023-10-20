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

resource "aws_service_discovery_service" "this" {
  name = var.name

  dns_config {
    namespace_id = var.namespace_id

    dns_records {
      ttl  = var.ttl
      type = var.dns_record_type
    }

    routing_policy = var.routing_policy
  }

  health_check_custom_config {
    failure_threshold = var.health_check_failure_threshold
  }

  tags = var.tags
}
