package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCustomConnectorInputResource(t *testing.T) {
	resourceName := "trocco_custom_connector_input.test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + LoadTextFile("testdata/custom_connector_input/basic_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "custom-connector-input-test"),
					resource.TestCheckResourceAttr(resourceName, "description", "test custom connector input"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://example.com"),
					resource.TestCheckResourceAttr(resourceName, "auth_type", "api_key"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.name", "list_users"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.path", "/users"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.success_codes", "200"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.query_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.query_parameters.0.name", "limit"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.request_body_parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.request_body_parameters.0.name", "status"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.request_body_parameters.0.default_value", "active"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.paginator.inject_into", "query"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.paginator.page_increment_strategy.page_size", "100"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.request_timeout_sec", "60"),
					// list_minimal sets only `name`/`path`; every other endpoint
					// attribute falls back to the schema default, which mirrors
					// the server's.
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.name", "list_minimal"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.method", "GET"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.jsonpath_root", "$.*"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.success_codes", "200"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.not_retryable_codes", "400,401,403,404"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.request_timeout_sec", "30"),
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
				Config: providerConfig + LoadTextFile("testdata/custom_connector_input/basic_update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "updated description"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.success_codes", "200,201"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.headers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.headers.0.name", "X-Api-Version"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.request_body_parameters.0.default_value", "inactive"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.request_body_parameters.0.is_editable", "false"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.paginator.offset_increment_strategy.page_size", "50"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.request_timeout_sec", "90"),
					// The defaults are re-sent on PATCH rather than nulled out.
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.method", "GET"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.jsonpath_root", "$.*"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.request_timeout_sec", "30"),
				),
			},
		},
	})
}

func TestAccCustomConnectorInputResourceInvalidConfig(t *testing.T) {
	testCases := []struct {
		name        string
		configFile  string
		expectError string
	}{
		{
			name:        "multiple_pagination_strategies",
			configFile:  "testdata/custom_connector_input/invalid_multiple_pagination_strategies.tf",
			expectError: `attributes specified when one`,
		},
		{
			name:        "field_option_missing_default",
			configFile:  "testdata/custom_connector_input/invalid_field_option_missing_default.tf",
			expectError: `default_value is required when`,
		},
		{
			name:        "duplicate_endpoint_name",
			configFile:  "testdata/custom_connector_input/invalid_duplicate_endpoint_name.tf",
			expectError: `Endpoint name .* is duplicated`,
		},
		{
			name:        "path_parameter_not_in_path",
			configFile:  "testdata/custom_connector_input/invalid_path_parameter_not_in_path.tf",
			expectError: `must appear in the endpoint path`,
		},
		{
			name:        "oauth2_missing_access_token_uri",
			configFile:  "testdata/custom_connector_input/invalid_oauth2_missing_access_token_uri.tf",
			expectError: `access_token_uri.*is required when`,
		},
		{
			name:        "last_page_size_missing_max_request_count",
			configFile:  "testdata/custom_connector_input/invalid_last_page_size_missing_max_request_count.tf",
			expectError: `must be specified when`,
		},
		{
			name:        "paginator_request_body_missing_page_token_option",
			configFile:  "testdata/custom_connector_input/invalid_paginator_request_body_missing_page_token_option.tf",
			expectError: `page_token_option.*is required when`,
		},
		{
			name:        "paginator_field_name_liquid_format",
			configFile:  "testdata/custom_connector_input/invalid_paginator_field_name_liquid_format.tf",
			expectError: `must be a valid`,
		},
		{
			name:        "paginator_duplicate_field_name",
			configFile:  "testdata/custom_connector_input/invalid_paginator_duplicate_field_name.tf",
			expectError: `must not be the`,
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
