package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDbtJobDefinitionResource(t *testing.T) {
	resourceName := "trocco_dbt_job_definition.test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + LoadTextFile("testdata/dbt_job_definition/basic_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "dbt-job-test"),
					resource.TestCheckResourceAttr(resourceName, "description", "test dbt job"),
					resource.TestCheckResourceAttr(resourceName, "adapter_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "threads", "4"),
					resource.TestCheckResourceAttr(resourceName, "target", "prod"),
					resource.TestCheckResourceAttr(resourceName, "bigquery_setting.dataset", "analytics"),
					resource.TestCheckResourceAttr(resourceName, "bigquery_setting.location", "asia-northeast1"),
					resource.TestCheckResourceAttr(resourceName, "commands.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "commands.0.command", "run"),
					resource.TestCheckResourceAttr(resourceName, "commands.0.options.0.key", "--vars"),
					resource.TestCheckResourceAttr(resourceName, "commands.1.command", "test"),
					resource.TestCheckResourceAttr(resourceName, "custom_variable_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "custom_variable_settings.0.name", "$ds$"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				Config: providerConfig + LoadTextFile("testdata/dbt_job_definition/basic_update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "dbt-job-test-renamed"),
					resource.TestCheckResourceAttr(resourceName, "description", "updated description"),
					resource.TestCheckResourceAttr(resourceName, "threads", "8"),
					resource.TestCheckResourceAttr(resourceName, "target", "dev"),
					resource.TestCheckResourceAttr(resourceName, "bigquery_setting.dataset", "analytics_v2"),
					resource.TestCheckResourceAttr(resourceName, "commands.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "commands.0.command", "seed"),
					resource.TestCheckResourceAttr(resourceName, "commands.1.command", "snapshot"),
					resource.TestCheckResourceAttr(resourceName, "custom_variable_settings.#", "0"),
				),
			},
		},
	})
}

func TestAccDbtJobDefinitionResourceInvalidConfig(t *testing.T) {
	testCases := []struct {
		name        string
		configFile  string
		expectError string
	}{
		{
			name:        "multiple_settings",
			configFile:  "testdata/dbt_job_definition/invalid_multiple_settings.tf",
			expectError: `Attribute .* cannot be specified when`,
		},
		{
			name:        "threads_out_of_range",
			configFile:  "testdata/dbt_job_definition/invalid_threads.tf",
			expectError: `Attribute threads value must be between 1 and 16`,
		},
		{
			name:        "command_enum",
			configFile:  "testdata/dbt_job_definition/invalid_command.tf",
			expectError: `Attribute commands\[\d+\].command value must be one of`,
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
