package main

import (
	"context"
	"flag"
	"log"
	version2 "terraform-provider-trocco/version"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"terraform-provider-trocco/internal/provider"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate -provider-name trocco

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/trocco-io/trocco",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version2.ProviderVersion), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
