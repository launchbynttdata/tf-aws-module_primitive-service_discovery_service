package testimpl

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	TestComposableCompleteReadonly(t, ctx)
}

func TestComposableCompleteReadonly(t *testing.T, ctx types.TestContext) {
	client := GetAWSServiceDiscoveryClient(t)
	opts := ctx.TerratestTerraformOptions()
	serviceID := terraform.Output(t, opts, "service_id")

	t.Run("TestServiceDiscoveryServiceExists", func(t *testing.T) {
		out, err := client.GetService(context.TODO(), &servicediscovery.GetServiceInput{
			Id: aws.String(serviceID),
		})
		require.NoError(t, err, "GetService should succeed")
		require.NotNil(t, out.Service, "Service should be present")
		assert.Equal(t, serviceID, *out.Service.Id, "Service ID should match Terraform output")
		assert.Regexp(t, `^arn:aws:servicediscovery:[a-z0-9-]+:[0-9]{12}:service/.+$`, *out.Service.Arn, "ARN should match expected pattern")
		assert.Equal(t, "test-service70238", *out.Service.Name, "Service name should match the complete example")
		assert.Equal(t, "MULTIVALUE", string(out.Service.DnsConfig.RoutingPolicy), "Routing policy should be MULTIVALUE")
	})
}

func GetAWSServiceDiscoveryClient(t *testing.T) *servicediscovery.Client {
	return servicediscovery.NewFromConfig(GetAWSConfig(t))
}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}
