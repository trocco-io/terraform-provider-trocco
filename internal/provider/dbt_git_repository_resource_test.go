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
					resource.TestCheckResourceAttr(resourceName, "ref_type", "branch"),
					resource.TestCheckResourceAttr(resourceName, "branch", "main"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr("trocco_dbt_git_repository.test_with_subdirectory", "subdirectory", "dbt/"),
					resource.TestCheckResourceAttr("trocco_dbt_git_repository.test_with_tag", "ref_type", "tag"),
					resource.TestCheckResourceAttr("trocco_dbt_git_repository.test_with_tag", "tag", "v1.0.0"),
					resource.TestCheckResourceAttr("trocco_dbt_git_repository.test_with_commit_hash", "ref_type", "commit_hash"),
					resource.TestCheckResourceAttr("trocco_dbt_git_repository.test_with_commit_hash", "commit_hash", "0123456789abcdef0123456789abcdef01234567"),
				),
			},
			// ImportState testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing (branch -> tag)
			{
				Config: providerConfig + LoadTextFile("testdata/dbt_git_repository/basic_update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test_repo_renamed"),
					resource.TestCheckResourceAttr(resourceName, "description", "updated description"),
					resource.TestCheckResourceAttr(resourceName, "dbt_version", "1.10"),
					resource.TestCheckResourceAttr(resourceName, "url", "git@github.com:example/repo-new.git"),
					resource.TestCheckResourceAttr(resourceName, "ref_type", "tag"),
					resource.TestCheckResourceAttr(resourceName, "tag", "v2.0.0"),
					resource.TestCheckResourceAttr(resourceName, "subdirectory", "dbt/"),
				),
			},
		},
	})
}

func TestAccDbtGitRepositoryResourceInvalidConfig(t *testing.T) {
	testCases := []struct {
		name        string
		configFile  string
		expectError string
	}{
		{
			name:        "invalid_adapter_type",
			configFile:  "testdata/dbt_git_repository/invalid_adapter_type.tf",
			expectError: `Attribute adapter_type value must be one of`,
		},
		{
			name:        "ref_type_conflict",
			configFile:  "testdata/dbt_git_repository/invalid_ref_type_conflict.tf",
			expectError: `must not be set when .ref_type. is`,
		},
		{
			name:        "ref_type_missing_value",
			configFile:  "testdata/dbt_git_repository/invalid_ref_type_missing.tf",
			expectError: `is required when .ref_type. is`,
		},
		{
			name:        "invalid_commit_hash",
			configFile:  "testdata/dbt_git_repository/invalid_commit_hash.tf",
			expectError: `must be a 40-character lowercase hexadecimal Git commit hash`,
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
