package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccUserResource(t *testing.T) {
	resourceName := "trocco_user.test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + LoadTextFile("testdata/user/basic_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "email", "test@example.com"),
					resource.TestCheckResourceAttr(resourceName, "role", "admin"),
					resource.TestCheckResourceAttr(resourceName, "can_use_audit_log", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_restricted_connection_modify", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
			// Update and Read testing
			{
				Config: providerConfig + LoadTextFile("testdata/user/basic_update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "email", "test@example.com"),
					resource.TestCheckResourceAttr(resourceName, "role", "member"),
					resource.TestCheckResourceAttr(resourceName, "can_use_audit_log", "true"),
					resource.TestCheckResourceAttr(resourceName, "is_restricted_connection_modify", "true"),
				),
			},
		},
	})
}

func TestAccUserResourceInvalidEmail(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/user/invalid_email.tf"),
				ExpectError: regexp.MustCompile(`invalid email address`),
			},
		},
	})
}

func TestAccUserResourceInvalidPassword(t *testing.T) {
	testCases := []struct {
		name        string
		configFile  string
		expectError string
	}{
		{
			name:        "short_password",
			configFile:  "testdata/user/invalid_password_short.tf",
			expectError: "password string length",
		},
		{
			name:        "no_letter",
			configFile:  "testdata/user/invalid_password_no_letter.tf",
			expectError: "must contain at least one letter",
		},
		{
			name:        "no_number",
			configFile:  "testdata/user/invalid_password_no_number.tf",
			expectError: "must contain at least one number",
		},
		{
			name:        "missing_password",
			configFile:  "testdata/user/missing_password.tf",
			expectError: "Missing Required Attribute",
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

func TestAccUserResourceInvalidRole(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/user/invalid_role.tf"),
				ExpectError: regexp.MustCompile(`role value must be one of`),
			},
		},
	})
}
