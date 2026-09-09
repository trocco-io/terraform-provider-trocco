package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"terraform-provider-trocco/internal/client"
	notificationDestinationParameters "terraform-provider-trocco/internal/client/parameter/notification_destination"
	"terraform-provider-trocco/internal/provider/model/notification_destination"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &notificationDestinationResource{}
	_ resource.ResourceWithConfigure   = &notificationDestinationResource{}
	_ resource.ResourceWithImportState = &notificationDestinationResource{}
)

func NewNotificationDestinationResource() resource.Resource {
	return &notificationDestinationResource{}
}

type notificationDestinationResource struct {
	client *client.TroccoClient
}

type notificationDestinationResourceModel struct {
	Type               types.String                                 `tfsdk:"type"`
	ID                 types.Int64                                  `tfsdk:"id"`
	EmailConfig        *notification_destination.EmailConfig        `tfsdk:"email_config"`
	SlackChannelConfig *notification_destination.SlackChannelConfig `tfsdk:"slack_channel_config"`
	HTTPConfig         *notification_destination.HTTPConfig         `tfsdk:"http_config"`
}

func (m *notificationDestinationResourceModel) ToCreateNotificationDestinationInput() *client.CreateNotificationDestinationInput {
	input := &client.CreateNotificationDestinationInput{}

	switch m.Type.ValueString() {
	case "email":
		if m.EmailConfig != nil {
			input.EmailConfig = &notificationDestinationParameters.EmailConfigInput{
				Email: m.EmailConfig.Email.ValueStringPointer(),
			}
		}
	case "slack_channel":
		if m.SlackChannelConfig != nil {
			input.SlackChannelConfig = &notificationDestinationParameters.SlackChannelConfigInput{
				Channel:    m.SlackChannelConfig.Channel.ValueStringPointer(),
				WebhookURL: m.SlackChannelConfig.WebhookURL.ValueStringPointer(),
			}
		}
	case "http":
		input.HTTPConfig = toHTTPConfigInput(m.HTTPConfig)
	}

	return input
}

func (m *notificationDestinationResourceModel) ToUpdateNotificationDestinationInput() *client.UpdateNotificationDestinationInput {
	input := &client.UpdateNotificationDestinationInput{}

	switch m.Type.ValueString() {
	case "email":
		if m.EmailConfig != nil {
			input.EmailConfig = &notificationDestinationParameters.EmailConfigInput{
				Email: m.EmailConfig.Email.ValueStringPointer(),
			}
		}
	case "slack_channel":
		if m.SlackChannelConfig != nil {
			input.SlackChannelConfig = &notificationDestinationParameters.SlackChannelConfigInput{
				Channel:    m.SlackChannelConfig.Channel.ValueStringPointer(),
				WebhookURL: m.SlackChannelConfig.WebhookURL.ValueStringPointer(),
			}
		}
	case "http":
		input.HTTPConfig = toHTTPConfigInput(m.HTTPConfig)
	}

	return input
}

// toHTTPConfigInput converts the http_config block into the API input.
// Headers and query parameters are always sent (an empty list when omitted) so that
// the API replaces the whole list and the resource does not drift from the configuration.
func toHTTPConfigInput(c *notification_destination.HTTPConfig) *notificationDestinationParameters.HTTPConfigInput {
	if c == nil {
		return nil
	}
	description := c.Description.ValueString()
	return &notificationDestinationParameters.HTTPConfigInput{
		Name:        c.Name.ValueStringPointer(),
		URL:         c.URL.ValueStringPointer(),
		Description: &description,
		Headers:     toHTTPKeyValueInputs(c.Headers),
		QueryParams: toHTTPKeyValueInputs(c.QueryParams),
	}
}

func toHTTPKeyValueInputs(items []notification_destination.HTTPKeyValue) *[]notificationDestinationParameters.HTTPKeyValueInput {
	result := make([]notificationDestinationParameters.HTTPKeyValueInput, 0, len(items))
	for _, item := range items {
		result = append(result, notificationDestinationParameters.HTTPKeyValueInput{
			Key:     item.Key.ValueString(),
			Value:   item.Value.ValueString(),
			Masking: item.Masking.ValueBool(),
		})
	}
	return &result
}

// httpConfigFromPlan builds the state of http_config after Create/Update. Name, URL and
// description come from the API response; headers and query parameters come from the plan
// because the API returns an empty value for masked entries.
func httpConfigFromPlan(plan *notification_destination.HTTPConfig, api *client.NotificationDestination) *notification_destination.HTTPConfig {
	if plan == nil {
		return nil
	}
	return &notification_destination.HTTPConfig{
		Name:        types.StringPointerValue(api.Name),
		URL:         types.StringPointerValue(api.URL),
		Description: httpDescriptionValue(api.Description),
		Headers:     plan.Headers,
		QueryParams: plan.QueryParams,
	}
}

// httpConfigFromAPI builds the state of http_config on Read. The API response is the source
// of truth (its lists are ordered by id). A masked value is returned as an empty string by the
// API, so it is restored from the previous state only when the entry at the same index has the
// same key. Other entries keep the API value, so changes made outside Terraform show up as drift.
func httpConfigFromAPI(state *notification_destination.HTTPConfig, api *client.NotificationDestination) *notification_destination.HTTPConfig {
	var stateHeaders, stateQueryParams []notification_destination.HTTPKeyValue
	if state != nil {
		stateHeaders = state.Headers
		stateQueryParams = state.QueryParams
	}
	return &notification_destination.HTTPConfig{
		Name:        types.StringPointerValue(api.Name),
		URL:         types.StringPointerValue(api.URL),
		Description: httpDescriptionValue(api.Description),
		Headers:     mergeHTTPKeyValues(api.Headers, stateHeaders),
		QueryParams: mergeHTTPKeyValues(api.QueryParams, stateQueryParams),
	}
}

func mergeHTTPKeyValues(apiItems []client.HTTPKeyValue, stateItems []notification_destination.HTTPKeyValue) []notification_destination.HTTPKeyValue {
	if len(apiItems) == 0 {
		if stateItems != nil {
			// keep an explicit empty list distinct from an omitted attribute
			return []notification_destination.HTTPKeyValue{}
		}
		return nil
	}
	result := make([]notification_destination.HTTPKeyValue, 0, len(apiItems))
	for i, item := range apiItems {
		value := item.Value
		if item.Masking && value == "" && i < len(stateItems) && stateItems[i].Key.ValueString() == item.Key {
			value = stateItems[i].Value.ValueString()
		}
		result = append(result, notification_destination.HTTPKeyValue{
			Key:     types.StringValue(item.Key),
			Value:   types.StringValue(value),
			Masking: types.BoolValue(item.Masking),
		})
	}
	return result
}

// httpDescriptionValue maps an empty description from the API to null so that an omitted
// description in the configuration does not produce a diff.
func httpDescriptionValue(description *string) types.String {
	if description == nil || *description == "" {
		return types.StringNull()
	}
	return types.StringValue(*description)
}

func (r *notificationDestinationResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = fmt.Sprintf("%s_notification_destination", req.ProviderTypeName)
}

func (r *notificationDestinationResource) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}

	c, ok := req.ProviderData.(*client.TroccoClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = c
}

func (r *notificationDestinationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO notification destination resource.",
		Attributes: map[string]schema.Attribute{
			"type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: `The type of the notification destination. Must be one of "email", "slack_channel" or "http".`,
				Validators: []validator.String{
					stringvalidator.OneOf("email", "slack_channel", "http"),
				},
			},
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The ID of the notification destination.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"email_config": schema.SingleNestedAttribute{
				Optional: true,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("slack_channel_config")),
					objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("http_config")),
				},
				Attributes: map[string]schema.Attribute{
					"email": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`),
								"invalid email address",
							),
						},
						MarkdownDescription: "The email address to notify.",
					},
				},
			},
			"slack_channel_config": schema.SingleNestedAttribute{
				Optional: true,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("email_config")),
					objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("http_config")),
				},
				Attributes: map[string]schema.Attribute{
					"channel": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.UTF8LengthAtLeast(1),
						},
						MarkdownDescription: "The name of the Slack channel to notify.",
					},
					"webhook_url": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.UTF8LengthAtLeast(1),
						},
						MarkdownDescription: "The webhook URL of the Slack channel.",
					},
				},
			},
			"http_config": schema.SingleNestedAttribute{
				Optional: true,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("email_config")),
					objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("slack_channel_config")),
				},
				MarkdownDescription: "Configuration of an HTTP (webhook) notification destination. Header and query parameter values are stored in plain text in the Terraform state; mark secrets with `masking = true` so that they are masked in the TROCCO console and API responses.",
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.UTF8LengthAtLeast(1),
						},
						MarkdownDescription: "The name of the notification destination. Must be unique within the account.",
					},
					"url": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^https?://`), "must start with http:// or https://"),
						},
						MarkdownDescription: "The URL that receives the notification request.",
					},
					"description": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The description of the notification destination.",
					},
					"headers": schema.ListNestedAttribute{
						Optional:            true,
						MarkdownDescription: "HTTP headers sent with the notification request. The order is preserved.",
						NestedObject:        httpKeyValueNestedObject("header"),
					},
					"query_params": schema.ListNestedAttribute{
						Optional:            true,
						MarkdownDescription: "Query parameters appended to the URL of the notification request. The order is preserved.",
						NestedObject:        httpKeyValueNestedObject("query parameter"),
					},
				},
			},
		},
	}
}

func httpKeyValueNestedObject(kind string) schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"key": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: fmt.Sprintf("The name of the %s.", kind),
			},
			"value": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: fmt.Sprintf("The value of the %s. When `masking` is `true`, the API never returns this value; after import it has to be set manually.", kind),
			},
			"masking": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: fmt.Sprintf("Whether to mask the value of the %s in the TROCCO console and API responses. Defaults to `false`.", kind),
			},
		},
	}
}

func (r *notificationDestinationResource) ValidateConfig(
	ctx context.Context,
	req resource.ValidateConfigRequest,
	resp *resource.ValidateConfigResponse,
) {
	plan := &notificationDestinationResourceModel{}
	resp.Diagnostics.Append(req.Config.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	switch plan.Type.ValueString() {
	case "email":
		if plan.EmailConfig == nil {
			resp.Diagnostics.AddError(
				"Missing Email Config",
				"`email_config.email` is required when type is 'email'.",
			)
		}
	case "slack_channel":
		if plan.SlackChannelConfig == nil {
			resp.Diagnostics.AddError(
				"Missing Slack Channel Config",
				"`slack_channel_config` is required when type is 'slack_channel'.",
			)
			return
		}
	case "http":
		if plan.HTTPConfig == nil {
			resp.Diagnostics.AddError(
				"Missing HTTP Config",
				"`http_config` is required when type is 'http'.",
			)
			return
		}
	default:
		resp.Diagnostics.AddError("type", `"type" must be one of "email", "slack_channel" or "http".`)
	}
}

func (r *notificationDestinationResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	plan := &notificationDestinationResourceModel{}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	notification, err := r.client.CreateNotificationDestination(
		plan.Type.ValueString(),
		plan.ToCreateNotificationDestinationInput(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating notification_destination",
			fmt.Sprintf("Unable to create notification_destination, got error: %s", err),
		)
		return
	}

	newState := notificationDestinationResourceModel{
		Type: types.StringValue(plan.Type.ValueString()),
		ID:   types.Int64Value(notification.ID),
	}

	switch plan.Type.ValueString() {
	case "email":
		newState.EmailConfig = &notification_destination.EmailConfig{
			Email: types.StringPointerValue(notification.Email),
		}
	case "slack_channel":
		newState.SlackChannelConfig = &notification_destination.SlackChannelConfig{
			Channel:    types.StringPointerValue(notification.Channel),
			WebhookURL: plan.SlackChannelConfig.WebhookURL,
		}
	case "http":
		newState.HTTPConfig = httpConfigFromPlan(plan.HTTPConfig, notification)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *notificationDestinationResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	state := &notificationDestinationResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan := &notificationDestinationResourceModel{}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	notification, err := r.client.UpdateNotificationDestination(
		plan.Type.ValueString(),
		plan.ID.ValueInt64(),
		plan.ToUpdateNotificationDestinationInput(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating notification destination ",
			fmt.Sprintf("Unable to update notification destination, got error: %s", err),
		)
		return
	}

	newState := notificationDestinationResourceModel{
		Type: types.StringValue(plan.Type.ValueString()),
		ID:   types.Int64Value(notification.ID),
	}

	switch plan.Type.ValueString() {
	case "email":
		newState.EmailConfig = &notification_destination.EmailConfig{
			Email: types.StringPointerValue(notification.Email),
		}
	case "slack_channel":
		newState.SlackChannelConfig = &notification_destination.SlackChannelConfig{
			Channel:    types.StringPointerValue(notification.Channel),
			WebhookURL: plan.SlackChannelConfig.WebhookURL,
		}
	case "http":
		newState.HTTPConfig = httpConfigFromPlan(plan.HTTPConfig, notification)
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *notificationDestinationResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	state := &notificationDestinationResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	notification, err := r.client.GetNotificationDestination(
		state.Type.ValueString(),
		state.ID.ValueInt64(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading notification destination",
			fmt.Sprintf("Unable to read notification destination, got error: %s", err),
		)
		return
	}

	newState := notificationDestinationResourceModel{
		Type: types.StringValue(state.Type.ValueString()),
		ID:   types.Int64Value(notification.ID),
	}

	switch state.Type.ValueString() {
	case "email":
		newState.EmailConfig = &notification_destination.EmailConfig{
			Email: types.StringPointerValue(notification.Email),
		}
	case "slack_channel":
		if state.SlackChannelConfig != nil {
			newState.SlackChannelConfig = &notification_destination.SlackChannelConfig{
				Channel:    types.StringPointerValue(notification.Channel),
				WebhookURL: state.SlackChannelConfig.WebhookURL,
			}
		} else {
			newState.SlackChannelConfig = &notification_destination.SlackChannelConfig{
				Channel:    types.StringPointerValue(notification.Channel),
				WebhookURL: types.StringValue(""),
			}
		}
	case "http":
		newState.HTTPConfig = httpConfigFromAPI(state.HTTPConfig, notification)
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *notificationDestinationResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	s := &notificationDestinationResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, s)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.DeleteNotificationDestination(
		s.Type.ValueString(),
		s.ID.ValueInt64(),
	); err != nil {
		resp.Diagnostics.AddError(
			"Deleting notification destination",
			fmt.Sprintf("Unable to delete notification destination, got error: %s", err),
		)
		return
	}
}

func (r *notificationDestinationResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	idParts := strings.Split(req.ID, ",")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Importing notification destination",
			fmt.Sprintf("Expected import identifier with format: type,id. Got: %q", req.ID),
		)
		return
	}

	destType := idParts[0]

	id, err := strconv.ParseInt(idParts[1], 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing notification destination",
			fmt.Sprintf("Failed to parse ID: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("type"), destType)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
