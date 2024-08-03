package client

import (
	"fmt"
	"net/http"
	"net/url"
)

// List of datamart_definitions
// ref: https://documents.trocco.io/apidocs/get-datamart-definitions

type ListDatamartDefinitionsInput struct {
	Limit  *int
	Cursor *string
}

func (input *ListDatamartDefinitionsInput) SetLimit(limit int) {
	input.Limit = &limit
}

func (input *ListDatamartDefinitionsInput) SetCursor(cursor string) {
	input.Cursor = &cursor
}

type ListDatamartDefinitionsOutput struct {
	Items      []ListDatamartDefinitionsItem `json:"items"`
	NextCursor *string                       `json:"next_cursor"`
}

type ListDatamartDefinitionsItem struct {
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	DataWarehouseType string `json:"data_warehouse_type"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

const MaxListDatamartDefinitionsLimit = 200

func (client *TroccoClient) ListDatamartDefinitions(input *ListDatamartDefinitionsInput) (*ListDatamartDefinitionsOutput, error) {
	params := url.Values{}
	if input != nil && input.Limit != nil {
		limit := *input.Limit
		if limit > MaxListDatamartDefinitionsLimit || limit < 1 {
			return nil, fmt.Errorf("limit must be between 1 and %d", MaxListDatamartDefinitionsLimit)
		}
		params.Add("limit", fmt.Sprintf("%d", limit))
	}
	if input != nil && input.Cursor != nil {
		params.Add("cursor", *input.Cursor)
	}
	path := fmt.Sprintf("/api/datamart_definitions?%s", params.Encode())
	output := new(ListDatamartDefinitionsOutput)
	err := client.Do(http.MethodGet, path, nil, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Get a datamart_definition
// ref: https://documents.trocco.io/apidocs/get-datamart-definition

type GetDatamartDefinitionOutput struct {
	DatamartDefinition
}

type DatamartDefinition struct {
	ID                     int64                   `json:"id"`
	Name                   string                  `json:"name"`
	Description            string                  `json:"description"`
	DataWarehouseType      string                  `json:"data_warehouse_type"`
	IsRunnableConcurrently bool                    `json:"is_runnable_concurrently"`
	ResourceGroup          *ResourceGroup          `json:"resource_group"`
	CustomVariableSettings []CustomVariableSetting `json:"custom_variable_settings"`
	DatamartBigqueryOption *DatamartBigqueryOption `json:"datamart_bigquery_option"`
	CreatedAt              string                  `json:"created_at"`
	UpdatedAt              string                  `json:"updated_at"`
	Notifications          []DatamartNotification  `json:"notifications"`
	Schedules              []Schedule              `json:"schedules"`
	Labels                 []Label                 `json:"labels"`
}

type DatamartBigqueryOption struct {
	BigqueryConnectionID int64    `json:"bigquery_connection_id"`
	QueryMode            string   `json:"query_mode"`
	Query                string   `json:"query"`
	DestinationDataset   *string  `json:"destination_dataset"`
	DestinationTable     *string  `json:"destination_table"`
	WriteDisposition     *string  `json:"write_disposition"`
	BeforeLoad           *string  `json:"before_load"`
	Partitioning         *string  `json:"partitioning"`
	PartitioningTime     *string  `json:"partitioning_time"`
	PartitioningField    *string  `json:"partitioning_field"`
	ClusteringFields     []string `json:"clustering_fields"`
	Location             *string  `json:"location"`
}

type ResourceGroup struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CustomVariableSetting struct {
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Value     *string `json:"value"`
	Quantity  *int    `json:"quantity"`
	Unit      *string `json:"unit"`
	Direction *string `json:"direction"`
	Format    *string `json:"format"`
	TimeZone  *string `json:"time_zone"`
}

type DatamartNotification struct {
	DestinationType  string  `json:"destination_type"`
	SlackChannelID   *int64  `json:"slack_channel_id"`
	EmailID          *int64  `json:"email_id"`
	NotificationType string  `json:"notification_type"`
	NotifyWhen       *string `json:"notify_when"`
	RecordCount      *int64  `json:"record_count"`
	RecordOperator   *string `json:"record_operator"`
	Message          string  `json:"message"`
}

type Schedule struct {
	Frequency string `json:"frequency"`
	Minute    int    `json:"minute"`
	Hour      *int   `json:"hour"`
	Day       *int   `json:"day"`
	DayOfWeek *int   `json:"day_of_week"`
	TimeZone  string `json:"time_zone"`
}

type Label struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (client *TroccoClient) GetDatamartDefinition(id int64) (*GetDatamartDefinitionOutput, error) {
	path := fmt.Sprintf("/api/datamart_definitions/%d", id)
	output := new(GetDatamartDefinitionOutput)
	err := client.Do(http.MethodGet, path, nil, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Create a datamart_definition
// ref: https://documents.trocco.io/apidocs/post-datamart-definition

type CreateDatamartDefinitionInput struct {
	Name                   string                             `json:"name"`
	DatawarehouseType      string                             `json:"data_warehouse_type"`
	Description            *string                            `json:"description,omitempty"`
	IsRunnableConcurrently bool                               `json:"is_runnable_concurrently"`
	ResourceGroupID        *int64                             `json:"resource_group_id,omitempty"`
	CustomVariableSettings *[]CustomVariableSettingInput      `json:"custom_variable_settings,omitempty"`
	DatamartBigqueryOption *CreateDatamartBigqueryOptionInput `json:"datamart_bigquery_option,omitempty"`
}

func (input *CreateDatamartDefinitionInput) SetDescription(description string) {
	input.Description = &description
}

func (input *CreateDatamartDefinitionInput) SetResourceGroupID(resourceGroupID int64) {
	input.ResourceGroupID = &resourceGroupID
}

func (input *CreateDatamartDefinitionInput) SetCustomVariableSettings(customVariableSettings []CustomVariableSettingInput) {
	input.CustomVariableSettings = &customVariableSettings
}

func (input *CreateDatamartDefinitionInput) SetDatamartBigqueryOption(datamartBigqueryOption CreateDatamartBigqueryOptionInput) {
	input.DatamartBigqueryOption = &datamartBigqueryOption
}

type CustomVariableSettingInput struct {
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Value     *string `json:"value,omitempty"`
	Quantity  *int    `json:"quantity,omitempty"`
	Unit      *string `json:"unit,omitempty"`
	Direction *string `json:"direction,omitempty"`
	Format    *string `json:"format,omitempty"`
	TimeZone  *string `json:"time_zone,omitempty"`
}

func NewStringTypeCustomVariableSettingInput(
	name string,
	value string,
) CustomVariableSettingInput {
	return CustomVariableSettingInput{
		Name:  name,
		Type:  "string",
		Value: &value,
	}
}

func NewTimestampTypeCustomVariableSettingInput(
	name string,
	typ string,
	quantity int,
	unit string,
	direction string,
	format string,
	timeZone string,
) CustomVariableSettingInput {
	return CustomVariableSettingInput{
		Name:      name,
		Type:      typ,
		Quantity:  &quantity,
		Unit:      &unit,
		Direction: &direction,
		Format:    &format,
		TimeZone:  &timeZone,
	}
}

type CreateDatamartBigqueryOptionInput struct {
	BigqueryConnectionID int64     `json:"bigquery_connection_id"`
	QueryMode            string    `json:"query_mode"`
	Query                string    `json:"query"`
	DestinationDataset   *string   `json:"destination_dataset,omitempty"`
	DestinationTable     *string   `json:"destination_table,omitempty"`
	WriteDisposition     *string   `json:"write_disposition,omitempty"`
	BeforeLoad           *string   `json:"before_load,omitempty"`
	Partitioning         *string   `json:"partitioning,omitempty"`
	PartitioningTime     *string   `json:"partitioning_time,omitempty"`
	PartitioningField    *string   `json:"partitioning_field,omitempty"`
	ClusteringFields     *[]string `json:"clustering_fields,omitempty"`
	Location             *string   `json:"location,omitempty"`
}

func NewInsertModeCreateDatamartBigqueryOptionInput(
	bigqueryConnectionID int64,
	query string,
	destinationDataset string,
	destinationTable string,
	writeDisposition string,
) CreateDatamartBigqueryOptionInput {
	return CreateDatamartBigqueryOptionInput{
		BigqueryConnectionID: bigqueryConnectionID,
		QueryMode:            "insert",
		Query:                query,
		DestinationDataset:   &destinationDataset,
		DestinationTable:     &destinationTable,
		WriteDisposition:     &writeDisposition,
	}
}

func NewQueryModeCreateDatamartBigqueryOptionInput(
	bigqueryConnectionID int64,
	query string,
	location string,
) CreateDatamartBigqueryOptionInput {
	return CreateDatamartBigqueryOptionInput{
		BigqueryConnectionID: bigqueryConnectionID,
		QueryMode:            "query",
		Query:                query,
		Location:             &location,
	}
}

func (datamartBigqueryOption *CreateDatamartBigqueryOptionInput) SetBeforeLoad(beforeLoad string) {
	datamartBigqueryOption.BeforeLoad = &beforeLoad
}

func (datamartBigqueryOption *CreateDatamartBigqueryOptionInput) SetPartitioning(partitioning string) {
	datamartBigqueryOption.Partitioning = &partitioning
}

func (datamartBigqueryOption *CreateDatamartBigqueryOptionInput) SetPartitioningTime(partitioningTime string) {
	datamartBigqueryOption.PartitioningTime = &partitioningTime
}

func (datamartBigqueryOption *CreateDatamartBigqueryOptionInput) SetPartitioningField(partitioningField string) {
	datamartBigqueryOption.PartitioningField = &partitioningField
}

func (datamartBigqueryOption *CreateDatamartBigqueryOptionInput) SetClusteringFields(clusteringFields []string) {
	datamartBigqueryOption.ClusteringFields = &clusteringFields
}

type CreateDatamartDefinitionOutput struct {
	ID int64 `json:"id"`
}

func (client *TroccoClient) CreateDatamartDefinition(input *CreateDatamartDefinitionInput) (*CreateDatamartDefinitionOutput, error) {
	path := "/api/datamart_definitions"
	output := new(CreateDatamartDefinitionOutput)
	err := client.Do(http.MethodPost, path, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Update a datamart_definition
// ref: https://documents.trocco.io/apidocs/patch-datamart-definition

type UpdateDatamartDefinitionInput struct {
	Name                   *string                            `json:"name,omitempty"`
	Description            *string                            `json:"description,omitempty"`
	IsRunnableConcurrently *bool                              `json:"is_runnable_concurrently,omitempty"`
	ResourceGroupID        *int64                             `json:"resource_group_id,omitempty"`
	CustomVariableSettings *[]CustomVariableSettingInput      `json:"custom_variable_settings,omitempty"`
	DatamartBigqueryOption *UpdateDatamartBigqueryOptionInput `json:"datamart_bigquery_option,omitempty"`
	Schedules              *[]ScheduleInput                   `json:"schedules,omitempty"`
	Notifications          *[]DatamartNotificationInput       `json:"notifications,omitempty"`
	Labels                 *[]string                          `json:"labels,omitempty"`
}

func (input *UpdateDatamartDefinitionInput) SetName(name string) {
	input.Name = &name
}

func (input *UpdateDatamartDefinitionInput) SetDescription(description string) {
	input.Description = &description
}

func (input *UpdateDatamartDefinitionInput) SetIsRunnableConcurrently(isRunnableConcurrently bool) {
	input.IsRunnableConcurrently = &isRunnableConcurrently
}

func (input *UpdateDatamartDefinitionInput) SetResourceGroupID(resourceGroupID int64) {
	input.ResourceGroupID = &resourceGroupID
}

func (input *UpdateDatamartDefinitionInput) SetCustomVariableSettings(customVariableSettings []CustomVariableSettingInput) {
	input.CustomVariableSettings = &customVariableSettings
}

func (input *UpdateDatamartDefinitionInput) SetDatamartBigqueryOption(datamartBigqueryOption UpdateDatamartBigqueryOptionInput) {
	input.DatamartBigqueryOption = &datamartBigqueryOption
}

func (input *UpdateDatamartDefinitionInput) SetSchedules(schedules []ScheduleInput) {
	input.Schedules = &schedules
}

func (input *UpdateDatamartDefinitionInput) SetNotifications(notifications []DatamartNotificationInput) {
	input.Notifications = &notifications
}

func (input *UpdateDatamartDefinitionInput) SetLabels(labels []string) {
	input.Labels = &labels
}

type UpdateDatamartBigqueryOptionInput struct {
	BigqueryConnectionID *int64    `json:"bigquery_connection_id,omitempty"`
	QueryMode            *string   `json:"query_mode,omitempty"`
	Query                *string   `json:"query,omitempty"`
	DestinationDataset   *string   `json:"destination_dataset,omitempty"`
	DestinationTable     *string   `json:"destination_table,omitempty"`
	WriteDisposition     *string   `json:"write_disposition,omitempty"`
	BeforeLoad           *string   `json:"before_load,omitempty"`
	Partitioning         *string   `json:"partitioning,omitempty"`
	PartitioningTime     *string   `json:"partitioning_time,omitempty"`
	PartitioningField    *string   `json:"partitioning_field,omitempty"`
	ClusteringFields     *[]string `json:"clustering_fields,omitempty"`
	Location             *string   `json:"location,omitempty"`
}

func (datamartBigqueryOption *UpdateDatamartBigqueryOptionInput) SetBigqueryConnectionID(bigqueryConnectionID int64) {
	datamartBigqueryOption.BigqueryConnectionID = &bigqueryConnectionID
}

func (datamartBigqueryOption *UpdateDatamartBigqueryOptionInput) SetQueryMode(queryMode string) {
	datamartBigqueryOption.QueryMode = &queryMode
}

func (datamartBigqueryOption *UpdateDatamartBigqueryOptionInput) SetQuery(query string) {
	datamartBigqueryOption.Query = &query
}

func (datamartBigqueryOption *UpdateDatamartBigqueryOptionInput) SetDestinationDataset(destinationDataset string) {
	datamartBigqueryOption.DestinationDataset = &destinationDataset
}

func (datamartBigqueryOption *UpdateDatamartBigqueryOptionInput) SetDestinationTable(destinationTable string) {
	datamartBigqueryOption.DestinationTable = &destinationTable
}

func (datamartBigqueryOption *UpdateDatamartBigqueryOptionInput) SetWriteDisposition(writeDisposition string) {
	datamartBigqueryOption.WriteDisposition = &writeDisposition
}

func (datamartBigqueryOption *UpdateDatamartBigqueryOptionInput) SetBeforeLoad(beforeLoad string) {
	datamartBigqueryOption.BeforeLoad = &beforeLoad
}

func (datamartBigqueryOption *UpdateDatamartBigqueryOptionInput) SetPartitioning(partitioning string) {
	datamartBigqueryOption.Partitioning = &partitioning
}

func (datamartBigqueryOption *UpdateDatamartBigqueryOptionInput) SetPartitioningTime(partitioningTime string) {
	datamartBigqueryOption.PartitioningTime = &partitioningTime
}

func (datamartBigqueryOption *UpdateDatamartBigqueryOptionInput) SetPartitioningField(partitioningField string) {
	datamartBigqueryOption.PartitioningField = &partitioningField
}

func (datamartBigqueryOption *UpdateDatamartBigqueryOptionInput) SetClusteringFields(clusteringFields []string) {
	datamartBigqueryOption.ClusteringFields = &clusteringFields
}

func (datamartBigqueryOption *UpdateDatamartBigqueryOptionInput) SetLocation(location string) {
	datamartBigqueryOption.Location = &location
}

type ScheduleInput struct {
	Frequency string `json:"frequency"`
	Minute    int    `json:"minute"`
	Hour      *int   `json:"hour,omitempty"`
	Day       *int   `json:"day,omitempty"`
	DayOfWeek *int   `json:"day_of_week,omitempty"`
	TimeZone  string `json:"time_zone"`
}

func NewHourlyScheduleInput(
	minute int,
	timeZone string,
) ScheduleInput {
	return ScheduleInput{
		Frequency: "hourly",
		Minute:    minute,
		TimeZone:  timeZone,
	}
}

func NewDailyScheduleInput(
	hour int,
	minute int,
	timeZone string,
) ScheduleInput {
	return ScheduleInput{
		Frequency: "daily",
		Hour:      &hour,
		Minute:    minute,
		TimeZone:  timeZone,
	}
}

func NewWeeklyScheduleInput(
	dayOfWeek int,
	hour int,
	minute int,
	timeZone string,
) ScheduleInput {
	return ScheduleInput{
		Frequency: "weekly",
		DayOfWeek: &dayOfWeek,
		Hour:      &hour,
		Minute:    minute,
		TimeZone:  timeZone,
	}
}

func NewMonthlyScheduleInput(
	day int,
	hour int,
	minute int,
	timeZone string,
) ScheduleInput {
	return ScheduleInput{
		Frequency: "monthly",
		Day:       &day,
		Hour:      &hour,
		Minute:    minute,
		TimeZone:  timeZone,
	}
}

type DatamartNotificationInput struct {
	DestinationType  string  `json:"destination_type"`
	SlackChannelID   *int64  `json:"slack_channel_id,omitempty"`
	EmailID          *int64  `json:"email_id,omitempty"`
	NotificationType string  `json:"notification_type"`
	NotifyWhen       *string `json:"notify_when,omitempty"`
	RecordCount      *int64  `json:"record_count,omitempty"`
	RecordOperator   *string `json:"record_operator,omitempty"`
	Message          string  `json:"message"`
}

func NewSlackJobDatamartNotificationInput(
	slackChannelID int64,
	notifyWhen string,
	message string,
) DatamartNotificationInput {
	return DatamartNotificationInput{
		DestinationType:  "slack",
		SlackChannelID:   &slackChannelID,
		NotificationType: "job",
		NotifyWhen:       &notifyWhen,
		Message:          message,
	}
}

func NewEmailJobDatamartNotificationInput(
	emailID int64,
	notifyWhen string,
	message string,
) DatamartNotificationInput {
	return DatamartNotificationInput{
		DestinationType:  "email",
		EmailID:          &emailID,
		NotificationType: "job",
		NotifyWhen:       &notifyWhen,
		Message:          message,
	}
}

func NewSlackRecordDatamartNotificationInput(
	slackChannelID int64,
	recordCount int64,
	recordOperator string,
	message string,
) DatamartNotificationInput {
	return DatamartNotificationInput{
		DestinationType:  "slack",
		SlackChannelID:   &slackChannelID,
		NotificationType: "record",
		RecordCount:      &recordCount,
		RecordOperator:   &recordOperator,
		Message:          message,
	}
}

func NewEmailRecordDatamartNotificationInput(
	emailID int64,
	recordCount int64,
	recordOperator string,
	message string,
) DatamartNotificationInput {
	return DatamartNotificationInput{
		DestinationType:  "email",
		EmailID:          &emailID,
		NotificationType: "record",
		RecordCount:      &recordCount,
		RecordOperator:   &recordOperator,
		Message:          message,
	}
}

func (client *TroccoClient) UpdateDatamartDefinition(id int64, input *UpdateDatamartDefinitionInput) error {
	path := fmt.Sprintf("/api/datamart_definitions/%d", id)
	err := client.Do(http.MethodPatch, path, input, nil)
	if err != nil {
		return err
	}
	return nil
}

// Delete a datamart_definition
// ref: https://documents.trocco.io/apidocs/delete-datamart-definition

func (client *TroccoClient) DeleteDatamartDefinition(id int64) error {
	path := fmt.Sprintf("/api/datamart_definitions/%d", id)
	err := client.Do(http.MethodDelete, path, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
