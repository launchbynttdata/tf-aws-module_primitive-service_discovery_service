# Post-Deploy Functional Tests

This directory contains the functional test harness that validates the AWS Service Discovery Service module after deployment.

## Overview

The post-deploy functional tests use the [lcaf-component-terratest](https://github.com/launchbynttdata/lcaf-component-terratest) framework to:

1. Deploy the Terraform configuration from the `examples/` directory
2. Execute validation tests using the AWS SDK
3. Tear down the infrastructure after testing

## Test Implementation

The actual test logic is implemented in `/tests/testimpl/test_impl.go` and uses the AWS SDK for Go v2 to validate:

### Service Configuration
- Service exists in AWS Cloud Map
- Service ID matches Terraform output
- Service name is not empty
- Service ARN is generated correctly

### DNS Configuration
- DNS configuration exists
- Routing policy is valid (MULTIVALUE or WEIGHTED)
- DNS records are created (exactly one record)
- TTL is positive
- DNS record type is valid (A, AAAA, SRV, or CNAME)
- **Note**: `NamespaceId` field in `DnsConfig` is deprecated in AWS SDK v2

### Health Check Configuration
- Custom health check configuration exists (confirms feature is enabled)
- **Note**: `FailureThreshold` field is deprecated in AWS SDK v2 and always returns 1

### Tagging
- Tags are applied to the service
- Tag keys and values are not empty

## Running Tests

From the repository root:

```bash
make check
```

Or run tests directly:

```bash
# Run this test suite only
cd tests/post_deploy_functional
go test -v -timeout 30m

# Run all test suites sequentially (from repo root)
go test -v -p 1 ./tests/...
```

**Important**: When running multiple test suites with `go test ./tests/...`, use the `-p 1` flag to prevent parallel execution. Both `post_deploy_functional` and `pre_deploy_functional` use the same `examples/complete` directory, and parallel execution causes file conflicts.

## Test Structure

```
tests/
├── post_deploy_functional/
│   ├── main_test.go           # Test harness (DO NOT MODIFY)
│   └── README.md              # This file
├── testimpl/
│   ├── test_impl.go           # Test implementation with AWS SDK validation
│   └── types.go               # Type definitions
```

## AWS SDK Validation Approach

Unlike traditional Terratest patterns that iterate through Terraform outputs, this module uses direct AWS SDK calls to:

1. Retrieve the service using `GetService` API
2. Query service tags using `ListTagsForResource` API
3. Validate actual AWS resource state against Terraform configuration
4. Ensure deployed resources match the intended configuration

This approach provides stronger validation by confirming that Terraform settings are actually applied in AWS, not just reflected in state files.

## Environment Variables

Tests use AWS credentials from:
- AWS profile configured in `provider.tf`
- Environment variables (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`)
- IAM role if running in AWS environment
- SSO session if configured

## Dependencies

- Go 1.25+
- AWS SDK for Go v2 (`github.com/aws/aws-sdk-go-v2/service/servicediscovery`)
- Terratest (`github.com/gruntwork-io/terratest`)
- lcaf-component-terratest (`github.com/launchbynttdata/lcaf-component-terratest`)
