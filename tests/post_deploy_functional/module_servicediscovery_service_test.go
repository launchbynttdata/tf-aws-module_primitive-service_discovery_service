package test

import (
	"context"
	"io/fs"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"

	"github.com/stretchr/testify/assert"
)

const (
	base            = "../../examples/"
	testVarFileName = "/test.tfvars"
)

func TestServiceDiscoveryService(t *testing.T) {
	t.Parallel()
	examplesFolders, err := os.ReadDir(base)
	if err != nil {
		assert.Error(t, err)
	}
	forEveryExampleRunTests(t, examplesFolders)
}

func forEveryExampleRunTests(t *testing.T, examplesFolders []fs.DirEntry) {
	stage := test_structure.RunTestStage
	for _, file := range examplesFolders {
		dir := base + file.Name()
		if file.IsDir() {
			defer stage(t, "teardown_serviceDiscovery_service", func() { tearDownServiceDiscoveryService(t, dir) })
			stage(t, "setup_serviceDiscovery_service", func() { setupServiceDiscoveryServiceTest(t, dir) })
			stage(t, "test_serviceDiscovery_service", func() { testServiceDiscoveryService(t, dir) })
		}
	}
}

func setupServiceDiscoveryServiceTest(t *testing.T, dir string) {
	terraformOptions := &terraform.Options{
		TerraformDir: dir,
		VarFiles:     []string{dir + testVarFileName},
		NoColor:      true,
		Logger:       logger.Discard,
	}
	test_structure.SaveTerraformOptions(t, dir, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

func testServiceDiscoveryService(t *testing.T, dir string) {
	terraformOptions := test_structure.LoadTerraformOptions(t, dir)
	terraformOptions.Logger = logger.Discard

	actualServiceId := terraform.Output(t, terraformOptions, "service_id")
	assert.NotEmpty(t, actualServiceId, "Service Discovery Service Id is empty")

	actualServiceNamespaceARN := terraform.Output(t, terraformOptions, "arn")
	assert.NotEmpty(t, actualServiceNamespaceARN, "Service Namespace ARN is empty")

	actualServiceNamespaceId := terraform.Output(t, terraformOptions, "id")
	serviceDiscoveryService := getDeployedServiceDiscoveryService(t, actualServiceNamespaceId)

	expectedPatternARN := "^arn:aws:servicediscovery:[a-z0-9-]+:[0-9]{12}:service/.+$"
	assert.Regexp(t, expectedPatternARN, *(serviceDiscoveryService.Arn), "ARN does not match expected pattern")

	expectedName, err := terraform.GetVariableAsStringFromVarFileE(t, dir+testVarFileName, "service_name")
	assert.NoError(t, err)
	actualName := *(serviceDiscoveryService.Name)
	assert.Equal(t, expectedName, actualName, "Service Name does not match with expected Name.")

	assert.Equal(t, []types.RoutingPolicy{"MULTIVALUE", "WEIGHTED"}, serviceDiscoveryService.DnsConfig.RoutingPolicy.Values())
}

func getDeployedServiceDiscoveryService(t *testing.T, actualServiceNamespaceId string) types.ServiceSummary {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile(os.Getenv("AWS_PROFILE")),
	)
	assert.NoError(t, err, "can't connect to aws")
	client := servicediscovery.NewFromConfig(cfg)

	filterStruct := types.ServiceFilter{
		Name:      "NAMESPACE_ID",
		Values:    []string{actualServiceNamespaceId},
		Condition: "EQ",
	}
	input := &servicediscovery.ListServicesInput{
		Filters: []types.ServiceFilter{
			filterStruct,
		},
	}
	result, err := client.ListServices(context.TODO(), input)
	assert.NoError(t, err, "The expected service discovery service was not found")

	return result.Services[0]
}

func tearDownServiceDiscoveryService(t *testing.T, dir string) {
	terraformOptions := test_structure.LoadTerraformOptions(t, dir)
	terraformOptions.Logger = logger.Discard
	terraform.Destroy(t, terraformOptions)
}
