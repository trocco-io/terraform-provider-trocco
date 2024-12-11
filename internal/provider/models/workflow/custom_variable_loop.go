package workflow

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"

	we "terraform-provider-trocco/internal/client/entities/workflow"
	wp "terraform-provider-trocco/internal/client/parameters/workflow"
)

//
// CustomVariableLoop
//

type CustomVariableLoop struct {
	Type types.String `tfsdk:"type"`

	StringConfig    *StringCustomVariableLoopConfig    `tfsdk:"string_config"`
	PeriodConfig    *PeriodCustomVariableLoopConfig    `tfsdk:"period_config"`
	BigqueryConfig  *BigqueryCustomVariableLoopConfig  `tfsdk:"bigquery_config"`
	SnowflakeConfig *SnowflakeCustomVariableLoopConfig `tfsdk:"snowflake_config"`
	RedshiftConfig  *RedshiftCustomVariableLoopConfig  `tfsdk:"redshift_config"`
}

func NewCustomVariableLoop(en *we.CustomVariableLoop) *CustomVariableLoop {
	if en == nil {
		return nil
	}

	md := &CustomVariableLoop{
		Type: types.StringValue(en.Type),
	}

	if en.StringConfig != nil {
		md.StringConfig = NewStringCustomVariableLoopConfig(en.StringConfig)
	}
	if en.PeriodConfig != nil {
		md.PeriodConfig = NewPeriodCustomVariableLoopConfig(en.PeriodConfig)
	}
	if en.BigqueryConfig != nil {
		md.BigqueryConfig = NewBigqueryCustomVariableLoopConfig(en.BigqueryConfig)
	}
	if en.SnowflakeConfig != nil {
		md.SnowflakeConfig = NewSnowflakeCustomVariableLoopConfig(en.SnowflakeConfig)
	}
	if en.RedshiftConfig != nil {
		md.RedshiftConfig = NewRedshiftCustomVariableLoopConfig(en.RedshiftConfig)
	}

	return md
}

func (c *CustomVariableLoop) ToInput() wp.CustomVariableLoop {
	i := wp.CustomVariableLoop{
		Type: c.Type.ValueString(),
	}

	if c.StringConfig != nil {
		i.StringConfig = lo.ToPtr(c.StringConfig.ToInput())
	}
	if c.PeriodConfig != nil {
		i.PeriodConfig = lo.ToPtr(c.PeriodConfig.ToInput())
	}
	if c.BigqueryConfig != nil {
		i.BigqueryConfig = lo.ToPtr(c.BigqueryConfig.ToInput())
	}
	if c.SnowflakeConfig != nil {
		i.SnowflakeConfig = lo.ToPtr(c.SnowflakeConfig.ToInput())
	}
	if c.RedshiftConfig != nil {
		i.RedshiftConfig = lo.ToPtr(c.RedshiftConfig.ToInput())
	}

	return i
}

//
// StringCustomVariableLoopConfig
//

type StringCustomVariableLoopConfig struct {
	Variables []StringCustomVariableLoopVariable `tfsdk:"variables"`
}

func NewStringCustomVariableLoopConfig(en *we.StringCustomVariableLoopConfig) *StringCustomVariableLoopConfig {
	if en == nil {
		return nil
	}

	variables := []StringCustomVariableLoopVariable{}
	for _, variable := range en.Variables {
		variables = append(variables, NewStringCustomVariableLoopVariable(variable))
	}

	return &StringCustomVariableLoopConfig{
		Variables: variables,
	}
}

func (c *StringCustomVariableLoopConfig) ToInput() wp.StringCustomVariableLoopConfig {
	vs := []wp.StringCustomVariableLoopVariable{}
	for _, v := range c.Variables {
		vs = append(vs, v.ToInput())
	}

	return wp.StringCustomVariableLoopConfig{
		Variables: vs,
	}
}

type StringCustomVariableLoopVariable struct {
	Name   types.String   `tfsdk:"name"`
	Values []types.String `tfsdk:"values"`
}

func NewStringCustomVariableLoopVariable(en we.StringCustomVariableLoopVariable) StringCustomVariableLoopVariable {
	values := []types.String{}
	for _, val := range en.Values {
		values = append(values, types.StringValue(val))
	}

	return StringCustomVariableLoopVariable{
		Name:   types.StringValue(en.Name),
		Values: values,
	}
}

func (v *StringCustomVariableLoopVariable) ToInput() wp.StringCustomVariableLoopVariable {
	values := []string{}
	for _, val := range v.Values {
		values = append(values, val.ValueString())
	}

	return wp.StringCustomVariableLoopVariable{
		Name:   v.Name.ValueString(),
		Values: values,
	}
}

//
// PeriodCustomVariableLoopConfig
//

type PeriodCustomVariableLoopConfig struct {
	Interval  types.String                       `tfsdk:"interval"`
	TimeZone  types.String                       `tfsdk:"time_zone"`
	From      PeriodCustomVariableLoopFrom       `tfsdk:"from"`
	To        PeriodCustomVariableLoopTo         `tfsdk:"to"`
	Variables []PeriodCustomVariableLoopVariable `tfsdk:"variables"`
}

func NewPeriodCustomVariableLoopConfig(en *we.PeriodCustomVariableLoopConfig) *PeriodCustomVariableLoopConfig {
	if en == nil {
		return nil
	}

	variables := []PeriodCustomVariableLoopVariable{}
	for _, variable := range en.Variables {
		variables = append(variables, NewPeriodCustomVariableLoopVariable(variable))
	}

	return &PeriodCustomVariableLoopConfig{
		Interval:  types.StringValue(en.Interval),
		TimeZone:  types.StringValue(en.TimeZone),
		From:      NewPeriodCustomVariableLoopFrom(en.From),
		To:        NewPeriodCustomVariableLoopTo(en.To),
		Variables: variables,
	}
}

func (c *PeriodCustomVariableLoopConfig) ToInput() wp.PeriodCustomVariableLoopConfig {
	vars := []wp.PeriodCustomVariableLoopVariable{}
	for _, v := range c.Variables {
		vars = append(vars, v.ToInput())
	}

	return wp.PeriodCustomVariableLoopConfig{
		Interval:  c.Interval.ValueString(),
		TimeZone:  c.TimeZone.ValueString(),
		From:      c.From.ToInput(),
		To:        c.To.ToInput(),
		Variables: vars,
	}
}

type PeriodCustomVariableLoopFrom struct {
	Value types.Int64  `tfsdk:"value"`
	Unit  types.String `tfsdk:"unit"`
}

func NewPeriodCustomVariableLoopFrom(en we.PeriodCustomVariableLoopFrom) PeriodCustomVariableLoopFrom {
	return PeriodCustomVariableLoopFrom{
		Value: types.Int64Value(en.Value),
		Unit:  types.StringValue(en.Unit),
	}
}

func (f *PeriodCustomVariableLoopFrom) ToInput() wp.PeriodCustomVariableLoopFrom {
	return wp.PeriodCustomVariableLoopFrom{
		Value: f.Value.ValueInt64Pointer(),
		Unit:  f.Unit.ValueString(),
	}
}

type PeriodCustomVariableLoopTo struct {
	Value types.Int64  `tfsdk:"value"`
	Unit  types.String `tfsdk:"unit"`
}

func NewPeriodCustomVariableLoopTo(en we.PeriodCustomVariableLoopTo) PeriodCustomVariableLoopTo {
	return PeriodCustomVariableLoopTo{
		Value: types.Int64Value(en.Value),
		Unit:  types.StringValue(en.Unit),
	}
}

func (t *PeriodCustomVariableLoopTo) ToInput() wp.PeriodCustomVariableLoopTo {
	return wp.PeriodCustomVariableLoopTo{
		Value: t.Value.ValueInt64Pointer(),
		Unit:  t.Unit.ValueString(),
	}
}

type PeriodCustomVariableLoopVariable struct {
	Name   types.String                           `tfsdk:"name"`
	Offset PeriodCustomVariableLoopVariableOffset `tfsdk:"offset"`
}

func NewPeriodCustomVariableLoopVariable(en we.PeriodCustomVariableLoopVariable) PeriodCustomVariableLoopVariable {
	return PeriodCustomVariableLoopVariable{
		Name:   types.StringValue(en.Name),
		Offset: NewStringCustomVariableLoopVariableOffset(en.Offset),
	}
}

func (v *PeriodCustomVariableLoopVariable) ToInput() wp.PeriodCustomVariableLoopVariable {
	return wp.PeriodCustomVariableLoopVariable{
		Name:   v.Name.ValueString(),
		Offset: v.Offset.ToInput(),
	}
}

type PeriodCustomVariableLoopVariableOffset struct {
	Value types.Int64  `tfsdk:"value"`
	Unit  types.String `tfsdk:"unit"`
}

func NewStringCustomVariableLoopVariableOffset(en we.PeriodCustomVariableLoopVariableOffset) PeriodCustomVariableLoopVariableOffset {
	return PeriodCustomVariableLoopVariableOffset{
		Value: types.Int64Value(en.Value),
		Unit:  types.StringValue(en.Unit),
	}
}

func (o *PeriodCustomVariableLoopVariableOffset) ToInput() wp.PeriodCustomVariableLoopVariableOffset {
	return wp.PeriodCustomVariableLoopVariableOffset{
		Value: o.Value.ValueInt64Pointer(),
		Unit:  o.Unit.ValueString(),
	}
}

//
// BigqueryCustomVariableLoopConfig
//

type BigqueryCustomVariableLoopConfig struct {
	ConnectionID types.Int64    `tfsdk:"connection_id"`
	Query        types.String   `tfsdk:"query"`
	Variables    []types.String `tfsdk:"variables"`
}

func NewBigqueryCustomVariableLoopConfig(en *we.BigqueryCustomVariableLoopConfig) *BigqueryCustomVariableLoopConfig {
	if en == nil {
		return nil
	}

	vs := []types.String{}
	for _, v := range en.Variables {
		vs = append(vs, types.StringValue(v))
	}

	return &BigqueryCustomVariableLoopConfig{
		ConnectionID: types.Int64Value(en.ConnectionID),
		Query:        types.StringValue(en.Query),
		Variables:    vs,
	}
}

func (c *BigqueryCustomVariableLoopConfig) ToInput() wp.BigqueryCustomVariableLoopConfig {
	vs := []string{}
	for _, v := range c.Variables {
		vs = append(vs, v.ValueString())
	}

	return wp.BigqueryCustomVariableLoopConfig{
		ConnectionID: c.ConnectionID.ValueInt64(),
		Query:        c.Query.ValueString(),
		Variables:    vs,
	}
}

//
// RedshiftCustomVariableLoopConfig
//

type RedshiftCustomVariableLoopConfig struct {
	ConnectionID types.Int64    `tfsdk:"connection_id"`
	Query        types.String   `tfsdk:"query"`
	Database     types.String   `tfsdk:"database"`
	Variables    []types.String `tfsdk:"variables"`
}

func NewRedshiftCustomVariableLoopConfig(en *we.RedshiftCustomVariableLoopConfig) *RedshiftCustomVariableLoopConfig {
	if en == nil {
		return nil
	}

	vs := []types.String{}
	for _, v := range en.Variables {
		vs = append(vs, types.StringValue(v))
	}

	return &RedshiftCustomVariableLoopConfig{
		ConnectionID: types.Int64Value(en.ConnectionID),
		Query:        types.StringValue(en.Query),
		Database:     types.StringValue(en.Database),
		Variables:    vs,
	}
}

func (c *RedshiftCustomVariableLoopConfig) ToInput() wp.RedshiftCustomVariableLoopConfig {
	vs := []string{}
	for _, v := range c.Variables {
		vs = append(vs, v.ValueString())
	}

	return wp.RedshiftCustomVariableLoopConfig{
		ConnectionID: c.ConnectionID.ValueInt64(),
		Query:        c.Query.ValueString(),
		Database:     c.Database.ValueString(),
		Variables:    vs,
	}
}

//
// SnowflakeCustomVariableLoopConfig
//

type SnowflakeCustomVariableLoopConfig struct {
	ConnectionID types.Int64    `tfsdk:"connection_id"`
	Query        types.String   `tfsdk:"query"`
	Warehouse    types.String   `tfsdk:"warehouse"`
	Variables    []types.String `tfsdk:"variables"`
}

func NewSnowflakeCustomVariableLoopConfig(en *we.SnowflakeCustomVariableLoopConfig) *SnowflakeCustomVariableLoopConfig {
	if en == nil {
		return nil
	}

	vs := []types.String{}
	for _, v := range en.Variables {
		vs = append(vs, types.StringValue(v))
	}

	return &SnowflakeCustomVariableLoopConfig{
		ConnectionID: types.Int64Value(en.ConnectionID),
		Query:        types.StringValue(en.Query),
		Warehouse:    types.StringValue(en.Warehouse),
		Variables:    vs,
	}
}

func (c *SnowflakeCustomVariableLoopConfig) ToInput() wp.SnowflakeCustomVariableLoopConfig {
	vs := []string{}
	for _, v := range c.Variables {
		vs = append(vs, v.ValueString())
	}

	return wp.SnowflakeCustomVariableLoopConfig{
		ConnectionID: c.ConnectionID.ValueInt64(),
		Query:        c.Query.ValueString(),
		Warehouse:    c.Warehouse.ValueString(),
		Variables:    vs,
	}
}
