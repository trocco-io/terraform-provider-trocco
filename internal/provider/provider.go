// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client"
)

const DEFAULT_ENDPOINT = "https://trocco.io"

var _ provider.Provider = &TroccoProvider{}

type TroccoProvider struct {
	version string
}

type TroccoProviderModel struct {
	APIKey   types.String `tfsdk:"api_key"`
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *TroccoProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "trocco"
	resp.Version = p.version
}

func (p *TroccoProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"endpoint": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *TroccoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config TroccoProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.APIKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Unknown TROCCO api key",
			"The provider cannot create the TROCCO api client as there is an unknown configuration value for the TROCCO api key. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TROCCO_API_KEY environment variable.",
		)
	}
	if config.Endpoint.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Unknown TROCCO endpoint",
			"The provider cannot create the TROCCO api client as there is an unknown configuration value for the TROCCO endpoint. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TROCCO_ENDPOINT environment variable.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	api_key := os.Getenv("TROCCO_API_KEY")
	endpoint := os.Getenv("TROCCO_ENDPOINT")

	if !config.APIKey.IsNull() {
		api_key = config.APIKey.ValueString()
	}
	if !config.Endpoint.IsNull() {
		endpoint = config.Endpoint.ValueString()
	}

	if api_key == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing TROCCO api key",
			"The provider cannot create the TROCCO api client as there is a missing or empty value for the TROCCO api key. "+
				"Set the api_key value in the configuration or use the TROCCO_API_KEY environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
	if endpoint == "" {
		endpoint = DEFAULT_ENDPOINT
	}

	client := client.NewTroccoClient(endpoint, api_key, true)
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *TroccoProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewDatamartDefinitionResource,
	}
}

func (p *TroccoProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *TroccoProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &TroccoProvider{
			version: version,
		}
	}
}
