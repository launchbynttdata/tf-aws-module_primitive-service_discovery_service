# Pre-Deploy Functional Tests

This directory contains pre-deployment validation tests that run before infrastructure is provisioned.

## Overview

Pre-deploy functional tests validate:

- Terraform syntax and configuration
- Module input validation
- Provider configuration
- Policy compliance (OPA/Conftest)
- Static code analysis

## Purpose

These tests catch issues early in the development cycle without requiring actual AWS resource provisioning:

- **Syntax Validation**: Ensures Terraform code is syntactically correct
- **Type Checking**: Validates variable types and values
- **Policy Checks**: Enforces organizational policies via OPA
- **Linting**: Checks code style and best practices

## Running Tests

From the repository root:

```bash
make check
```

This will run pre-deploy validations as part of the full test suite.

## Test Structure

```
tests/
├── pre_deploy_functional/
│   ├── main_test.go           # Test harness
│   └── README.md              # This file
```

## What Gets Validated

1. **Terraform Commands**:
   - `terraform fmt -check`
   - `terraform validate`
   - `terraform plan`

2. **Policy Checks** (if configured):
   - OPA policies in `/components/policy/`
   - Custom validation rules

3. **Module Configuration**:
   - Required variables are provided
   - Variable types match declarations
   - Default values are appropriate

## Integration with CI/CD

Pre-deploy tests are ideal for CI/CD pipelines as they:
- Run quickly (no infrastructure provisioning)
- Provide fast feedback
- Catch configuration errors early
- Don't incur AWS costs

These tests should be run on every pull request before post-deploy functional tests.
