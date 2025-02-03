package provider

import (
	"context"
	"os"

	"terraform-provider-trocco/internal/client"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &TroccoProvider{}

type TroccoProvider struct {
	version string
}

const DefaultRegion = "japan"

type TroccoProviderModel struct {
	APIKey     types.String `tfsdk:"api_key"`
	Region     types.String `tfsdk:"region"`
	DevBaseURL types.String `tfsdk:"dev_base_url"`
}

func (p *TroccoProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "trocco"
	resp.Version = p.version
}

func (p *TroccoProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `
The Terraform Provider for TROCCO enables the management of TROCCO resources using the TROCCO API feature, which is available only with our paid plans.
    `,
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				MarkdownDescription: "Your TROCCO API key. This can also be set using the `TROCCO_API_KEY` environment variable.",
			},
			"region": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The region of TROCCO. This can also be set using the `TROCCO_REGION` environment variable. The following regions are available: `japan`, `india`, `korea`.",
				Validators: []validator.String{
					stringvalidator.OneOf("japan", "india", "korea"),
				},
			},
			"dev_base_url": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The base URL of API. This is used for only development purposes.",
			},
		},
	}
}

func (p *TroccoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config TroccoProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.APIKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Unknown api_key",
			"The provider cannot create the TROCCO client as there is an unknown configuration value for the api_key. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TROCCO_API_KEY environment variable.",
		)
	}
	if config.Region.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("region"),
			"Unknown region",
			"The provider cannot create the TROCCO client as there is an unknown configuration value for the region. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TROCCO_REGION environment variable.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	api_key := os.Getenv("TROCCO_API_KEY")
	region := os.Getenv("TROCCO_REGION")

	if !config.APIKey.IsNull() {
		api_key = config.APIKey.ValueString()
	}
	if !config.Region.IsNull() {
		region = config.Region.ValueString()
	}

	if api_key == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing api key",
			"The provider cannot create the TROCCO client as there is a missing or empty value for the api_key. "+
				"Set the api_key value in the configuration or use the TROCCO_API_KEY environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
	if region == "" {
		region = DefaultRegion
	}

	c, err := client.NewTroccoClientWithRegion(api_key, region)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Trocco client",
			err.Error(),
		)
		return
	}
	if !config.DevBaseURL.IsNull() {
		c = client.NewDevTroccoClient(api_key, config.DevBaseURL.ValueString())
	}
	resp.DataSourceData = c
	resp.ResourceData = c
}

func (p *TroccoProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewBigqueryDatamartDefinitionResource,
		NewConnectionResource,
		NewUserResource,
		NewPipelineDefinitionResource,
		NewJobDefinitionResource,
		NewTeamResource,
		NewResourceGroupResource,
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
