package provider

import (
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var (
	providerConfig = fmt.Sprintf(`
		provider "trocco" {
		  dev_base_url = "%s"
		}
	`, os.Getenv("TROCCO_TEST_URL"))

	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"trocco": providerserver.NewProtocol6WithError(New("test")()),
	}
)

// LoadTextFile reads a file from the filesystem and returns its content as a string.
func LoadTextFile(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic("Error loading file: " + err.Error())
	}
	return string(content)
}
