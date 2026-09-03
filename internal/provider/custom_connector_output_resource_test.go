package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCustomConnectorOutputResource(t *testing.T) {
	resourceName := "trocco_custom_connector_output.test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + LoadTextFile("testdata/custom_connector_output/basic_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "custom-connector-output-test"),
					resource.TestCheckResourceAttr(resourceName, "description", "test custom connector output"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://example.com"),
					resource.TestCheckResourceAttr(resourceName, "auth_type", "api_key"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.name", "create_user"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.request_type", "single"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.batch_size", "1"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.method", "POST"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.operation", "create"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.path", "/users"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.payload_type", "json"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.success_codes", "200,201,204"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.request_timeout_sec", "45"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.template", `{"name": "{{ name }}"}`),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.headers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.headers.0.name", "X-Api-Version"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.query_parameters.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "endpoints.0.id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: providerConfig + LoadTextFile("testdata/custom_connector_output/basic_update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "updated description"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.name", "create_user"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.request_type", "multiple"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.batch_size", "50"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.success_codes", "200,201"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.request_timeout_sec", "60"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.headers.0.default_value", "3"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.query_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.query_parameters.0.name", "dry_run"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.name", "update_user"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.method", "PATCH"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.operation", "update"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.path", "/users/{{user_id}}"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.path_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.path_parameters.0.value", "user_id"),
				),
			},
		},
	})
}

func TestAccCustomConnectorOutputResourceInvalidConfig(t *testing.T) {
	testCases := []struct {
		name        string
		configFile  string
		expectError string
	}{
		{
			name:        "duplicate_endpoint_name",
			configFile:  "testdata/custom_connector_output/invalid_duplicate_endpoint_name.tf",
			expectError: `Endpoint name .* is duplicated`,
		},
		{
			name:        "empty_endpoint_name",
			configFile:  "testdata/custom_connector_output/invalid_empty_endpoint_name.tf",
			expectError: `string length must be at least 1`,
		},
		{
			name:        "batch_size_for_single_request",
			configFile:  "testdata/custom_connector_output/invalid_batch_size_for_single_request.tf",
			expectError: `batch_size.*must be 1 when`,
		},
		{
			name:        "path_parameter_not_in_path",
			configFile:  "testdata/custom_connector_output/invalid_path_parameter_not_in_path.tf",
			expectError: `must appear in the endpoint path`,
		},
		{
			name:        "field_option_missing_default",
			configFile:  "testdata/custom_connector_output/invalid_field_option_missing_default.tf",
			expectError: `default_value is required when`,
		},
		{
			name:        "oauth2_missing_access_token_uri",
			configFile:  "testdata/custom_connector_output/invalid_oauth2_missing_access_token_uri.tf",
			expectError: `access_token_uri.*is required when`,
		},
		{
			name:        "api_key_with_oauth2_attributes",
			configFile:  "testdata/custom_connector_output/invalid_api_key_with_oauth2_attributes.tf",
			expectError: `can only be set when`,
		},
		{
			name:        "authorization_code_missing_auth_uri",
			configFile:  "testdata/custom_connector_output/invalid_authorization_code_missing_auth_uri.tf",
			expectError: `auth_uri.*is required when`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config:      providerConfig + LoadTextFile(tc.configFile),
						ExpectError: regexp.MustCompile(tc.expectError),
					},
				},
			})
		})
	}
}
