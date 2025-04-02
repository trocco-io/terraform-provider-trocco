package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"terraform-provider-trocco/internal/client"
	notification_parameter "terraform-provider-trocco/internal/client/parameter/notification_destination"
	"terraform-provider-trocco/internal/provider/model/notification_destination"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
}

func (m *notificationDestinationResourceModel) ToCreateNotificationDestinationInput() *client.CreateNotificationDestinationInput {
	input := &client.CreateNotificationDestinationInput{}

	switch m.Type.ValueString() {
	case "email":
		if m.EmailConfig != nil {
			input.EmailConfig = &notification_parameter.EmailConfigInput{
				Email: m.EmailConfig.Email.ValueStringPointer(),
			}
		}
	case "slack_channel":
		if m.SlackChannelConfig != nil {
			input.SlackChannelConfig = &notification_parameter.SlackChannelConfigInput{
				Channel:    m.SlackChannelConfig.Channel.ValueStringPointer(),
				WebhookURL: m.SlackChannelConfig.WebhookURL.ValueStringPointer(),
			}
		}
	}

	return input
}

func (m *notificationDestinationResourceModel) ToUpdateNotificationDestinationInput() *client.UpdateNotificationDestinationInput {
	input := &client.UpdateNotificationDestinationInput{}

	switch m.Type.ValueString() {
	case "email":
		if m.EmailConfig != nil {
			input.EmailConfig = &notification_parameter.EmailConfigInput{
				Email: m.EmailConfig.Email.ValueStringPointer(),
			}
		}
	case "slack_channel":
		if m.SlackChannelConfig != nil {
			input.SlackChannelConfig = &notification_parameter.SlackChannelConfigInput{
				Channel:    m.SlackChannelConfig.Channel.ValueStringPointer(),
				WebhookURL: m.SlackChannelConfig.WebhookURL.ValueStringPointer(),
			}
		}
	}

	return input
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
				MarkdownDescription: `The type of the notification destination. Must be either "email" or "slack_channel".`,
				Validators: []validator.String{
					stringvalidator.OneOf("email", "slack_channel"),
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
	default:
		resp.Diagnostics.AddError("type", `"type" must be either "email" or "slack_channel".`)
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
