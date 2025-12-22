# GitHub Copilot Instructions

## File Editing Rules

- **NEVER remove license headers** from files when making edits.
- Always preserve the Apache 2.0 license header at the top of all source files (`.tf`, `.go`, `.sh`, etc.).
- When editing files, include the full license header in replacements if modifying code within the first 20 lines of the file.

### Expected License Header

The following Apache 2.0 license header must be present at the top of all source files. The comment syntax varies by file type:

**Go files (`.go`):**

```go
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
```

**Terraform (`.tf`) and Shell (`.sh`) files:**
Use `#` comment syntax instead of `//` with the same license text.

## Terminal Command Rules

- **DO NOT use timeout flags** with terminal commands (e.g., avoid `-timeout` with go test)
- Allow commands to output fully without truncation to ensure complete diagnostic information. Do not use `less`, `more`, `head`, `tail` or similar pagers to truncate output.
- If a command needs to be stopped, the user will cancel it manually.

## Terraform Best Practices

### Repository Structure

Follow this standardized module structure:

```text
/
├── examples/
│   ├── simple/                 # Minimum required variables + bare minimum dependencies
│   │   ├── locals.tf, main.tf, outputs.tf, README.md, test.tfvars, variables.tf, versions.tf
│   ├── complete/               # All optional variables and configurations
│   └── [custom-example]/       # For mutually exclusive configs (e.g., RDS vs Aurora Limitless)
│
├── tests/
│   ├── post_deploy_functional/        # DO NOT MODIFY
│   ├── post_deploy_functional_readonly/ # DO NOT MODIFY
│   └── testimpl/                      # MODIFY: test_impl.go, types.go
│
├── .gitignore, .pre-commit-config.yaml, .secrets.baseline, .terraform-docs.yaml, .tflint.hcl, .tool-versions  # DO NOT MODIFY
├── CHANGELOG.md                # Update when making changes
├── go.mod, go.sum              # Managed by Go toolchain
├── LICENSE, NOTICE, Makefile   # DO NOT MODIFY
├── locals.tf, main.tf, outputs.tf, README.md, variables.tf, versions.tf
```

#### Example Purpose Guidelines

- **simple/**: Demonstrates minimal viable usage. Include only bare minimum dependent resources (e.g., VPC for certificate module, but not EC2/RDS). All values exposed via variables; variables without safe defaults defined in test.tfvars.
- **complete/**: Attempts to set every configurable parameter to show full feature set.
- **[custom-example]/**: Create when mutually exclusive configuration branches exist that cannot coexist in "complete".

### Coding Standards

- Maintain consistency with existing patterns across all files.
- Always validate configurations with `terraform validate` before planning or applying.
- When building multiple examples ensure that they use common "outputs" to deduplicate testing functions.

## Testing Guidelines

- Write tests that verify actual AWS resource creation and configuration.
- Use the AWS SDK to verify resource properties match Terraform outputs.
- **Critical**: Every setting configured in Terraform must be validated against the actual deployed AWS resource using the AWS SDK. This ensures Terraform configurations are actually applied in the environment and aren't just reflected in Terraform state/outputs.
- Test coverage must include:
  - All required parameters
  - At least one configuration using optional parameters
  - Resource naming conventions
  - Security-related settings (encryption, access controls, etc.)
  - Compliance-related configurations (tagging, logging, etc.)
  - Settings that affect resource accessibility, durability, or cost

## Documentation Standards

- Prioritize explaining **why** decisions were made (design rationale, use cases, trade-offs) over **how** to use basic syntax.
- Include **how** information through clear examples when demonstrating complex configurations.
- Write concisely—avoid redundancy and unnecessary detail.
- Update documentation whenever code functionality changes.
- Document all changes in CHANGELOG.md using [Keep a Changelog](https://keepachangelog.com/) format:
  - Group changes by type: `Added`, `Changed`, `Deprecated`, `Removed`, `Fixed`, `Security`
  - Organize by version number and date (versions are set via automated semantic versioning)
  - Include brief descriptions of what changed and why when relevant
- Do not create individual change documentation files.

### Documentation Update Requirements

**Always update when making changes:**

- **README.md** (root and example directories):
  - Add/update when adding new variables or outputs
  - Update examples when changing resource configuration
  - Update feature descriptions when adding capabilities
  - Refresh terraform-docs generated sections (runs automatically via pre-commit)

- **CHANGELOG.md** (root only):
  - Document every functional change before committing
  - Include rationale for breaking changes
  - Reference related issues or PRs when applicable

## Terraform Primitive Module Development

### Design Principles

- **Reusability**: Design modules for use across multiple projects and teams.
- **No Opinions**: Do not hardcode defaults, naming conventions, organizational policies, or best-practice configurations. All settings must be exposed as variables.
- **Single Resource**: Wrap only one resource type per module. Exception: Include required data sources that the primary resource must query to function properly.
  - Maintain a 1:1 relationship between module and resource type in primitive modules.
  - Collection modules may call nested module/primitive dependencies, but this should be avoided unless necessary as it creates complexity.
- **Root-level Code**: Place all Terraform module code (main.tf, variables.tf, outputs.tf, etc.) in the repository root.

### Variable and Output Naming Conventions

**Primitive Modules** (single resource):

- Use variable names exactly as defined in the upstream Terraform resource documentation.
- Outputs should mirror the resource attribute names without prefixes.
- Organize variables in `variables.tf` in alphabetical order.
- Variable descriptions should reference the upstream Terraform module documentation for that provider version.

**Collection Modules** (multiple resources/primitives):

- Prefix each variable and output with the full resource/primitive name to prevent naming conflicts.
- Example: `aws_acmpca_certificate_authority_certificate_certificate`, `aws_acmpca_certificate_authority_type`
- Organize variables alphabetically within each resource group, with resource groups ordered by their appearance in `main.tf`.
- Organize outputs similarly.
- Variable descriptions should reference the upstream Terraform module documentation for that provider version.

### Variable Defaults and Validation

**Default Values**:

- Provide safe, recommended default values wherever realistic to simplify usage.
- Required variables without sensible defaults should omit the `default` attribute.
- Optional variables should have defaults set to match the upstream resource's default behavior.

**Validation Blocks**:

- Keep validation minimal—rely on upstream AWS services to catch most constraint violations.
- Apply validation for known service-specific constraints documented in upstream Terraform provider or AWS CLI reference.
  - Example: AWS RDS for PostgreSQL has restrictions on username values and password complexity requirements.
- Use validation to enforce requirements that would otherwise cause cryptic deployment failures.

### Conditional Logic and Resource Iteration

**for_each vs count vs dynamic blocks**:

- Use `for_each` when creating multiple instances of a resource from a collection (e.g., multiple security groups, subnets).
- Use `count` for simple enable/disable patterns or fixed quantities.
- Use `dynamic` blocks for optional nested configuration blocks within a resource.
- Consider map-based structures when the resource's dynamic objects require key-value relationships.

**Testing with Dynamic Resources**:

- Design collection-based resources to support selective enabling/disabling via array indices for testing.
- Example: For multi-AZ resources across zones a/b/c, use indices `[0,1,2]` to create all, `[0,2]` to exclude zone b, etc.
- This allows tests to validate resource stability when others are created/destroyed.
- In Go tests, use `SetVarsAfterVarFiles = true` to override variables:

  ```go
  // CRITICAL: Set this flag to ensure -var flags come after -var-file flags
  // This allows our variable override to take precedence over test.tfvars
  terraformOptions.SetVarsAfterVarFiles = true

  terraformOptions.Vars = map[string]interface{}{
      "enabled_subnet_indices": []int{1, 2}, // Keep indices 1 and 2, remove 0
  }
  ```

### AWS-Specific Considerations

**Regional Resources**:

- Not all AWS resources require or support explicit region configuration.
- If a resource requires AWS region(s) to be specified, expose them as module variables.
- If the resource does not require explicit region configuration, rely on transitive provider configuration.
- Do not hardcode region values—allow them to be inherited from the AWS provider or passed explicitly when needed.

**Authentication**:

- Assume transitive credentials (SSO, assumed roles, environment variables) for AWS authentication.
- Do not hardcode or expose AWS credentials in module code or examples.
- Modules should work with standard AWS provider authentication mechanisms.

**Account-Specific Values**:

- Avoid hardcoding AWS account IDs, ARNs, or other account-specific identifiers.
- When account-specific values are required, expose them as variables or use data sources to retrieve them dynamically.

### Quality Control

**Pre-commit Hooks**:

- The `.pre-commit-config.yaml` file contains advanced linting checks that run automatically before commits.
- These checks validate code quality (trailing whitespace, merge conflicts, credential exposure, etc.).
- Do not modify this file—it is managed separately and should be treated as infrastructure.

### File Modification Rules

The agent may modify:

- **Root directory**: All `.tf` files (main.tf, variables.tf, outputs.tf, locals.tf, versions.tf)
- **Examples**: All files within `/examples/` subdirectories
- **Tests**: Only `/tests/testimpl/test_impl.go` and `/tests/testimpl/types.go`
- **Documentation**: README.md and CHANGELOG.md **must be updated** when:
  - Adding, removing, or modifying variables or outputs
  - Changing module behavior or functionality
  - Adding or updating test coverage
  - Modifying example implementations
  - Fixing bugs or addressing security issues

The agent must NOT modify:

- `/tests/post_deploy_functional/main_test.go`
- `/tests/post_deploy_functional_readonly/main_test.go`
- Configuration files (.gitignore, .pre-commit-config.yaml, .tflint.hcl, etc.)
- LICENSE, NOTICE, Makefile
- Do not create new test files or directories
