# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed
- Rebuilt test implementation to use AWS SDK for direct resource validation
- Migrated from output iteration pattern to AWS API-based validation
- Updated test implementation to validate all Service Discovery Service configuration settings
- Enhanced test coverage to include DNS configuration, health check settings, and tags validation
- Modified validation to use actual AWS resource state instead of tfvars file parsing

### Added
- Comprehensive documentation for test implementation approach
- Detailed README files for test directories explaining AWS SDK validation methodology
- Service Discovery Service feature descriptions and use cases in main README
- Complete example documentation with actual test.tfvars values
- AWS SDK for Go v2 servicediscovery client dependency

### Fixed
- Test implementation no longer relies on tfOpts.Vars which is not populated from -var-file
- Tests now validate AWS resource state directly without needing to parse tfvars files
- Tags validation correctly retrieves service ARN for tag listing
- Validation functions now check for valid ranges and values instead of exact matches
- Documented requirement for sequential test execution (`-p 1` flag) to avoid parallel execution conflicts
- Removed usage of deprecated AWS SDK v2 fields (`DnsConfig.NamespaceId`, `HealthCheckCustomConfig.FailureThreshold`)
- Fixed linter warnings: removed redundant nil check before len() operation

### Known Limitations
- When running `go test ./tests/...`, the `-p 1` flag must be used to run test packages sequentially
- Both test suites share the same `examples/complete` directory, which causes conflicts during parallel execution
- Individual test suites can be run independently without this limitation

## Description of Changes

This update refactors the test implementation to align with best practices for AWS resource validation. Instead of relying solely on Terraform outputs, tests now:

1. Query AWS Service Discovery Service directly using the AWS SDK
2. Validate all configured parameters against actual deployed resources
3. Ensure Terraform configurations are actually applied in AWS (not just reflected in state)
4. Provide comprehensive coverage of service name, DNS configuration (record type, TTL, routing policy), health check settings, and resource tags

This approach provides stronger validation guarantees and ensures the module correctly provisions AWS resources as intended.
