# tf-aws-module_primitive-service_discovery_service

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY-NC-ND 4.0](https://img.shields.io/badge/License-CC_BY--NC--ND_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nc-nd/4.0/)

## Overview

This module provisions an AWS Cloud Map Service Discovery Service that can be referenced by ECS Services or other applications to enable service discovery. The service uses DNS-based service discovery with configurable health checks.

### Key Features

- DNS-based service discovery using AWS Cloud Map
- Configurable DNS record types (A, AAAA, SRV, CNAME)
- Support for MULTIVALUE and WEIGHTED routing policies
- Custom health check configuration with adjustable failure thresholds
- Resource tagging support
- Integration with private DNS namespaces

### Use Cases

- Service discovery for ECS tasks within a VPC
- Microservices communication in containerized environments
- Internal service mesh implementations
- Application discovery without hard-coded endpoints

## Pre-Commit hooks

[.pre-commit-config.yaml](.pre-commit-config.yaml) file defines certain `pre-commit` hooks that are relevant to terraform, golang and common linting tasks. There are no custom hooks added.

`commitlint` hook enforces commit message in certain format. The commit contains the following structural elements, to communicate intent to the consumers of your commit messages:

- **fix**: a commit of the type `fix` patches a bug in your codebase (this correlates with PATCH in Semantic Versioning).
- **feat**: a commit of the type `feat` introduces a new feature to the codebase (this correlates with MINOR in Semantic Versioning).
- **BREAKING CHANGE**: a commit that has a footer `BREAKING CHANGE:`, or appends a `!` after the type/scope, introduces a breaking API change (correlating with MAJOR in Semantic Versioning). A BREAKING CHANGE can be part of commits of any type.
footers other than BREAKING CHANGE: <description> may be provided and follow a convention similar to git trailer format.
- **build**: a commit of the type `build` adds changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)
- **chore**: a commit of the type `chore` adds changes that don't modify src or test files
- **ci**: a commit of the type `ci` adds changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)
- **docs**: a commit of the type `docs` adds documentation only changes
- **perf**: a commit of the type `perf` adds code change that improves performance
- **refactor**: a commit of the type `refactor` adds code change that neither fixes a bug nor adds a feature
- **revert**: a commit of the type `revert` reverts a previous commit
- **style**: a commit of the type `style` adds code changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- **test**: a commit of the type `test` adds missing tests or correcting existing tests

Base configuration used for this project is [commitlint-config-conventional (based on the Angular convention)](https://github.com/conventional-changelog/commitlint/tree/master/@commitlint/config-conventional#type-enum)

If you are a developer using vscode, [this](https://marketplace.visualstudio.com/items?itemName=joshbolduc.commitlint) plugin may be helpful.

`detect-secrets-hook` prevents new secrets from being introduced into the baseline. TODO: INSERT DOC LINK ABOUT HOOKS

In order for `pre-commit` hooks to work properly

- You need to have the pre-commit package manager installed. [Here](https://pre-commit.com/#install) are the installation instructions.
- `pre-commit` would install all the hooks when commit message is added by default except for `commitlint` hook. `commitlint` hook would need to be installed manually using the command below

```
pre-commit install --hook-type commit-msg
```

## To test the module locally

1. For development/enhancements to this module locally, you'll need to install all of its components. This is controlled by the `configure` target in the project's [`Makefile`](./Makefile). Before you can run `configure`, familiarize yourself with the variables in the `Makefile` and ensure they're pointing to the right places.

```
make configure
```

This adds in several files and directories that are ignored by `git`. They expose many new Make targets.

2. The first target you care about is `env`. This is the common interface for setting up environment variables. The values of the environment variables will be used to authenticate with cloud provider from local development workstation.

`make configure` command will bring down `aws_env.sh` file on local workstation. Developer would need to modify this file, replace the environment variable values with relevant values.

These environment variables are used by `terratest` integration suite.

Then run this make target to set the environment variables on developer workstation.

```
make env
```

3. Run the `check` target to validate the module.

**Pre-requisites**

Before running this target it is important to ensure that, developer has created files mentioned below on local workstation under root directory of git repository that contains code for primitives/segments. Note that these files are `aws` specific. If primitive/segment under development uses any other cloud provider than AWS, this section may not be relevant.

- A file named `provider.tf` with contents below

```
provider "aws" {
  profile = "<profile_name>"
  region  = "<region_name>"
}
```

- A file named `terraform.tfvars` which contains key value pair of variables used.

Note that since these files are added in `gitignore` they would not be checked in into primitive/segment's git repo.

After creating these files, for running tests associated with the primitive/segment, run

```
make check
```

If `make check` target is successful, developer is good to commit the code to primitive/segment's git repo.

`make check` target

- runs `terraform commands` to `lint`,`validate` and `plan` terraform code.
- runs `conftests`. `conftests` make sure `policy` checks are successful.
- runs `terratest`. This is integration test suite.
- runs `opa` tests

## Testing

This module uses comprehensive AWS SDK-based testing to validate deployed resources. The test implementation:

- **Direct AWS API Validation**: Uses the AWS SDK for Go v2 (`servicediscovery` client) to query actual deployed resources
- **State Verification**: Validates that AWS resources exist and have valid configurations
- **Comprehensive Coverage**: Tests service existence, DNS configuration (record type, TTL, routing policy), health check settings, and tags
- **No tfvars Parsing**: Tests validate actual AWS state rather than comparing against input variables

### Test Coverage

The test suite validates:

- ✅ Service Discovery Service existence and basic attributes (ID, ARN, name)
- ✅ DNS configuration (record type validity, TTL positivity, routing policy, single DNS record)
- ✅ Health check custom configuration (presence confirms feature is enabled)
- ✅ Resource tagging (tags exist with non-empty keys and values)

**Note**: Some AWS SDK fields like `NamespaceId` in `DnsConfig` and `FailureThreshold` in `HealthCheckCustomConfig` are deprecated in AWS SDK v2. The tests validate the overall configuration correctness rather than individual deprecated fields.

### Running Tests

Tests are located in `/tests/testimpl/` and are executed via the standard test harness:

```bash
make check
```

For manual test execution:

```bash
# Run tests sequentially (recommended to avoid parallel execution conflicts)
go test -v -p 1 ./tests/...

# Or run individual test suites
cd tests/post_deploy_functional
go test -v -timeout 30m
```

**Note**: When running `go test ./tests/...`, use the `-p 1` flag to run test packages sequentially. Both test suites use the same `examples/complete` directory, so parallel execution causes file conflicts.

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.5 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.100 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.100.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_service_discovery_service.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/service_discovery_service) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_name"></a> [name](#input\_name) | Name of the service. The service will be discoverable by this name - <service\_name>.<namespace\_name> | `string` | n/a | yes |
| <a name="input_namespace_id"></a> [namespace\_id](#input\_namespace\_id) | ID of the Cloud Map namespace | `string` | n/a | yes |
| <a name="input_dns_record_type"></a> [dns\_record\_type](#input\_dns\_record\_type) | DNS record type for the service. Default is A record | `string` | `"A"` | no |
| <a name="input_ttl"></a> [ttl](#input\_ttl) | The amount of time, in seconds, that you want DNS resolvers to cache the settings for this resource record set. | `number` | `10` | no |
| <a name="input_routing_policy"></a> [routing\_policy](#input\_routing\_policy) | he routing policy that you want to apply to all records that Route 53 creates when you register an instance and specify the service. Valid Values: MULTIVALUE, WEIGHTED | `string` | `"MULTIVALUE"` | no |
| <a name="input_health_check_failure_threshold"></a> [health\_check\_failure\_threshold](#input\_health\_check\_failure\_threshold) | The number of 30-second intervals that you want service discovery to wait before it changes the health status of a service instance. Maximum value of 10 | `number` | `1` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of custom tags to be attached to this resource | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | ID of the service |
| <a name="output_arn"></a> [arn](#output\_arn) | ARN of the service |
<!-- END_TF_DOCS -->
