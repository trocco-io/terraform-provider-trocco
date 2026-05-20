package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDbtGitRepositoryResource(t *testing.T) {
	resourceName := "trocco_dbt_git_repository.test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + LoadTextFile("testdata/dbt_git_repository/basic_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test_repo"),
					resource.TestCheckResourceAttr(resourceName, "description", "test description"),
					resource.TestCheckResourceAttr(resourceName, "adapter_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "dbt_version", "1.11"),
					resource.TestCheckResourceAttr(resourceName, "url", "git@github.com:example/repo.git"),
					resource.TestCheckResourceAttr(resourceName, "branch", "main"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr("trocco_dbt_git_repository.test_with_subdirectory", "subdirectory", "dbt/"),
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
				Config: providerConfig + LoadTextFile("testdata/dbt_git_repository/basic_update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test_repo_renamed"),
					resource.TestCheckResourceAttr(resourceName, "description", "updated description"),
					resource.TestCheckResourceAttr(resourceName, "dbt_version", "1.10"),
					resource.TestCheckResourceAttr(resourceName, "url", "git@github.com:example/repo-new.git"),
					resource.TestCheckResourceAttr(resourceName, "branch", "develop"),
					resource.TestCheckResourceAttr(resourceName, "subdirectory", "dbt/"),
				),
			},
		},
	})
}

func TestAccDbtGitRepositoryResourceInvalidAdapterType(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/dbt_git_repository/invalid_adapter_type.tf"),
				ExpectError: regexp.MustCompile(`Attribute adapter_type value must be one of`),
			},
		},
	})
}
