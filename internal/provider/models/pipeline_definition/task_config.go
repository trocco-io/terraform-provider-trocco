package workflow

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"

	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"
)

//
// TroccoTransferConfig
//

//
// TroccoTransferBulkTaskConfig
//

type TroccoTransferBulkTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`
}

func NewTroccoTransferBulkTaskConfig(c *we.TroccoTransferBulkTaskConfig) *TroccoTransferBulkTaskConfig {
	if c == nil {
		return nil
	}

	return &TroccoTransferBulkTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),
	}
}

func (c *TroccoTransferBulkTaskConfig) ToInput() *wp.TroccoTransferBulkTaskConfig {
	return &wp.TroccoTransferBulkTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}
}

//
// DBTTaskConfig
//

type DBTTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`
}

func NewDBTTaskConfig(c *we.DBTTaskConfig) *DBTTaskConfig {
	if c == nil {
		return nil
	}

	return &DBTTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),
	}
}

func (c *DBTTaskConfig) ToInput() *wp.DBTTaskConfig {
	return &wp.DBTTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}
}

//
// TroccoAgentTaskConfig
//

type TroccoAgentTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`
}

func NewTroccoAgentTaskConfig(c *we.TroccoAgentTaskConfig) *TroccoAgentTaskConfig {
	if c == nil {
		return nil
	}

	return &TroccoAgentTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),
	}
}

func (c *TroccoAgentTaskConfig) ToInput() *wp.TroccoAgentTaskConfig {
	return &wp.TroccoAgentTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}
}

//
// TroccoBigQueryDatamartTaskConfig
//

type TroccoBigQueryDatamartTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`

	CustomVariableLoop *CustomVariableLoop `tfsdk:"custom_variable_loop"`
}

func NewTroccoBigQueryDatamartTaskConfig(c *we.TroccoBigQueryDatamartTaskConfig) *TroccoBigQueryDatamartTaskConfig {
	if c == nil {
		return nil
	}

	return &TroccoBigQueryDatamartTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),

		CustomVariableLoop: NewCustomVariableLoop(c.CustomVariableLoop),
	}
}

func (c *TroccoBigQueryDatamartTaskConfig) ToInput() *wp.TroccoBigQueryDatamartTaskConfig {
	in := &wp.TroccoBigQueryDatamartTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}

	if c.CustomVariableLoop != nil {
		in.CustomVariableLoop = lo.ToPtr(c.CustomVariableLoop.ToInput())
	}

	return in
}

//
// TroccoRedshiftDatamartTaskConfig
//

type TroccoRedshiftDatamartTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`

	CustomVariableLoop *CustomVariableLoop `tfsdk:"custom_variable_loop"`
}

func NewTroccoRedshiftDatamartTaskConfig(c *we.TroccoRedshiftDatamartTaskConfig) *TroccoRedshiftDatamartTaskConfig {
	if c == nil {
		return nil
	}

	return &TroccoRedshiftDatamartTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),

		CustomVariableLoop: NewCustomVariableLoop(c.CustomVariableLoop),
	}
}

func (c *TroccoRedshiftDatamartTaskConfig) ToInput() *wp.TroccoRedshiftDatamartTaskConfig {
	in := &wp.TroccoRedshiftDatamartTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}

	if c.CustomVariableLoop != nil {
		in.CustomVariableLoop = lo.ToPtr(c.CustomVariableLoop.ToInput())
	}

	return in
}

//
// TroccoSnowflakeDatamartTaskConfig
//

type TroccoSnowflakeDatamartTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`

	CustomVariableLoop *CustomVariableLoop `tfsdk:"custom_variable_loop"`
}

func NewTroccoSnowflakeDatamartTaskConfig(c *we.TroccoSnowflakeDatamartTaskConfig) *TroccoSnowflakeDatamartTaskConfig {
	if c == nil {
		return nil
	}

	return &TroccoSnowflakeDatamartTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),

		CustomVariableLoop: NewCustomVariableLoop(c.CustomVariableLoop),
	}
}

func (c *TroccoSnowflakeDatamartTaskConfig) ToInput() *wp.TroccoSnowflakeDatamartTaskConfig {
	in := &wp.TroccoSnowflakeDatamartTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}

	if c.CustomVariableLoop != nil {
		in.CustomVariableLoop = lo.ToPtr(c.CustomVariableLoop.ToInput())
	}

	return in
}

//
// WorkflowTaskConfig
//

type WorkflowTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`

	CustomVariableLoop *CustomVariableLoop `tfsdk:"custom_variable_loop"`
}

func NewWorkflowTaskConfig(c *we.WorkflowTaskConfig) *WorkflowTaskConfig {
	if c == nil {
		return nil
	}

	return &WorkflowTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),

		CustomVariableLoop: NewCustomVariableLoop(c.CustomVariableLoop),
	}
}

func (c *WorkflowTaskConfig) ToInput() *wp.WorkflowTaskConfig {
	in := &wp.WorkflowTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}

	if c.CustomVariableLoop != nil {
		in.CustomVariableLoop = lo.ToPtr(c.CustomVariableLoop.ToInput())
	}

	return in
}

//
// TableauDataExtractionTaskConfig
//

type TableauDataExtractionTaskConfig struct {
	Name         types.String `tfsdk:"name"`
	ConnectionID types.Int64  `tfsdk:"connection_id"`
	TaskID       types.String `tfsdk:"task_id"`
}

func NewTableauDataExtractionTaskConfig(c *we.TableauDataExtractionTaskConfig) *TableauDataExtractionTaskConfig {
	if c == nil {
		return nil
	}

	return &TableauDataExtractionTaskConfig{
		Name:         types.StringValue(c.Name),
		ConnectionID: types.Int64Value(c.ConnectionID),
		TaskID:       types.StringValue(c.TaskID),
	}
}

func (c *TableauDataExtractionTaskConfig) ToInput() *wp.TableauDataExtractionTaskConfig {
	return &wp.TableauDataExtractionTaskConfig{
		Name:         c.Name.ValueString(),
		ConnectionID: c.ConnectionID.ValueInt64(),
		TaskID:       c.TaskID.ValueString(),
	}
}
