package testimpl

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/gruntwork-io/terratest/modules/terraform"
	testTypes "github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/require"
)

// TestComposableComplete is the main test function called by both functional and readonly test harnesses
// It discovers and tests all examples in the configured examples folder
func TestComposableComplete(t *testing.T, ctx testTypes.TestContext) {
	tfOpts := ctx.TerratestTerraformOptions()
	examplesDir := tfOpts.TerraformDir

	// Discover all example subdirectories
	entries, err := os.ReadDir(examplesDir)
	require.NoError(t, err, "Failed to read examples directory: %s", examplesDir)

	// Test each example subdirectory
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		exampleName := entry.Name()
		examplePath := filepath.Join(examplesDir, exampleName)

		// Check if this directory has a test.tfvars file
		testVarsPath := filepath.Join(examplePath, "test.tfvars")
		if _, err := os.Stat(testVarsPath); os.IsNotExist(err) {
			t.Logf("Skipping %s - no test.tfvars found", exampleName)
			continue
		}

		// Run test for this example
		t.Run(exampleName, func(t *testing.T) {
			// Update the terraform directory to point to the specific example
			originalDir := tfOpts.TerraformDir
			tfOpts.TerraformDir = examplePath
			defer func() { tfOpts.TerraformDir = originalDir }()

			// Detect which test to run based on example name or outputs
			switch exampleName {
			case "simple":
				TestComposableSimpleExample(t, ctx)
			case "complete":
				TestComposableCompleteExample(t, ctx)
			default:
				// Try to auto-detect based on outputs
				_, err := terraform.OutputE(t, tfOpts, "crl_bucket_name")
				if err == nil {
					TestComposableCompleteExample(t, ctx)
				} else {
					TestComposableSimpleExample(t, ctx)
				}
			}
		})
	}
}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}
