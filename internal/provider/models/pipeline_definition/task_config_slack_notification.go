package workflow

import (
	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SlackNotificationTaskConfig struct {
	Name         types.String `tfsdk:"name"`
	ConnectionID types.Int64  `tfsdk:"connection_id"`
	Message      types.String `tfsdk:"message"`
}

func NewSlackNotificationTaskConfig(c *we.SlackNotificationTaskConfig) *SlackNotificationTaskConfig {
	if c == nil {
		return nil
	}

	return &SlackNotificationTaskConfig{
		Name:         types.StringValue(c.Name),
		ConnectionID: types.Int64Value(c.ConnectionID),
		Message:      types.StringValue(c.Message),
	}
}

func (c *SlackNotificationTaskConfig) ToInput() *wp.SlackNotificationTaskConfig {
	return &wp.SlackNotificationTaskConfig{
		Name:         c.Name.ValueString(),
		ConnectionID: c.ConnectionID.ValueInt64(),
		Message:      c.Message.ValueString(),
	}
}