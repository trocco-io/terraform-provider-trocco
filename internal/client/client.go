package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TroccoClient struct {
	Endpoint   string
	Token      string
	httpClient *http.Client
}

func NewTroccoClient(endpoint, token string, local bool) *TroccoClient {
	httpClient := &http.Client{}
	if local {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient = &http.Client{Transport: tr}
	}
	return &TroccoClient{
		Endpoint:   endpoint,
		Token:      token,
		httpClient: httpClient,
	}
}

func (client *TroccoClient) NewRequest(method, url string, input interface{}) (*http.Request, error) {
	var body io.Reader
	if input != nil {
		b, err := json.Marshal(input)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(b)
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Token "+client.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (client *TroccoClient) Do(req *http.Request, output interface{}) error {
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if output == nil {
		return nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, output)
	if err != nil {
		return err
	}
	return nil
}

type ListDatamartDefinitionsOutput struct {
	Items      []ListDatamartDefinitionsItem `json:"items"`
	NextCursor string                        `json:"next_cursor,omitempty"`
}
type ListDatamartDefinitionsItem struct {
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	DataWarehouseType string `json:"data_warehouse_type"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

const MAX_LIST_LIMIT = 200

func (client *TroccoClient) ListDatamartDefinitions(limit int, cursor string) (*ListDatamartDefinitionsOutput, error) {
	if limit > MAX_LIST_LIMIT || limit < 1 {
		return nil, fmt.Errorf("limit must be between 1 and %d", MAX_LIST_LIMIT)
	}
	var cursorStr string
	if cursor != "" {
		cursorStr = fmt.Sprintf("&cursor=%s", cursor)
	}
	url := fmt.Sprintf("%s/api/datamart_definitions?limit=%d%s", client.Endpoint, limit, cursorStr)
	req, err := client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	output := new(ListDatamartDefinitionsOutput)
	err = client.Do(req, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (client *TroccoClient) ListDatamartDefinitionsAll() ([]ListDatamartDefinitionsItem, error) {
	items := make([]ListDatamartDefinitionsItem, 0)
	var cursor string
	for {
		response, err := client.ListDatamartDefinitions(MAX_LIST_LIMIT, cursor)
		if err != nil {
			return nil, err
		}
		items = append(items, response.Items...)
		if response.NextCursor == "" {
			break
		}
		cursor = response.NextCursor
	}
	return items, nil
}

type GetDatamartDefinitionOutput struct {
	ID                     int64                      `json:"id"`
	Name                   string                     `json:"name"`
	Description            string                     `json:"description"`
	DataWarehouseType      string                     `json:"data_warehouse_type"`
	IsRunnableConcurrently bool                       `json:"is_runnable_concurrently"`
	ResourceGroup          GetResourceGroup           `json:"resource_group,omitempty"`
	CustomVariableSettings []GetCustomVariableSetting `json:"custom_variable_settings,omitempty"`
	DatamartBigqueryOption GetDatamartBigqueryOption  `json:"datamart_bigquery_option,omitempty"`
	CreatedAt              string                     `json:"created_at"`
	UpdatedAt              string                     `json:"updated_at"`
}

type GetDatamartBigqueryOption struct {
	BigqueryConnectionID int64    `json:"bigquery_connection_id"`
	QueryMode            string   `json:"query_mode"`
	Query                string   `json:"query"`
	DestinationDataset   string   `json:"destination_dataset,omitempty"`
	DestinationTable     string   `json:"destination_table,omitempty"`
	WriteDisposition     string   `json:"write_disposition,omitempty"`
	BeforeLoadQuery      string   `json:"before_load_query,omitempty"`
	Partitioning         string   `json:"partitioning,omitempty"`
	PartitioningTime     string   `json:"partitioning_time,omitempty"`
	PartitioningField    string   `json:"partitioning_field,omitempty"`
	ClusteringFields     []string `json:"clustering_fields,omitempty"`
	Location             string   `json:"location,omitempty"`
}

type GetResourceGroup struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type GetCustomVariableSetting struct {
	Name string `json:"name"`
	Type string `json:"type"`
	GetStringCustomVariableSetting
	GetTimestampCustomVariableSetting
}

type GetStringCustomVariableSetting struct {
	Value string `json:"value,omitempty"`
}
type GetTimestampCustomVariableSetting struct {
	Quantity  int    `json:"quantity,omitempty"`
	Unit      string `json:"unit,omitempty"`
	Direction string `json:"direction,omitempty"`
	Format    string `json:"format,omitempty"`
	TimeZone  string `json:"time_zone,omitempty"`
}

func (client *TroccoClient) GetDatamartDefinition(id int64) (*GetDatamartDefinitionOutput, error) {
	url := fmt.Sprintf("%s/api/datamart_definitions/%d", client.Endpoint, id)
	req, err := client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	output := new(GetDatamartDefinitionOutput)
	err = client.Do(req, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type CreateDatamartDefinitionsInput struct {
	Name                   string                        `json:"name"`
	DataWarehouseType      string                        `json:"data_warehouse_type"`
	Description            *string                       `json:"description,omitempty"`
	IsRunnableConcurrently bool                          `json:"is_runnable_concurrently"`
	ResourceGroupID        *int64                        `json:"resource_group_id,omitempty"`
	CustomVariableSettings []CreateCustomVariableSetting `json:"custom_variable_settings,omitempty"`
	DatamartBigqueryOption CreateDatamartBigqueryOption  `json:"datamart_bigquery_option,omitempty"`
}
type CreateCustomVariableSetting struct {
	Name string `json:"name"`
	Type string `json:"type"`
	CreateStringCustomVariableSetting
	CreateTimestampCustomVariableSetting
}
type CreateStringCustomVariableSetting struct {
	Value *string `json:"value,omitempty"`
}
type CreateTimestampCustomVariableSetting struct {
	Quantity  *int    `json:"quantity,omitempty"`
	Unit      *string `json:"unit,omitempty"`
	Direction *string `json:"direction,omitempty"`
	Format    *string `json:"format,omitempty"`
	TimeZone  *string `json:"time_zone,omitempty"`
}
type CreateDatamartBigqueryOption struct {
	BigqueryConnectionID int64    `json:"bigquery_connection_id"`
	QueryMode            string   `json:"query_mode"`
	Query                string   `json:"query"`
	DestinationDataset   *string  `json:"destination_dataset,omitempty"`
	DestinationTable     *string  `json:"destination_table,omitempty"`
	WriteDisposition     *string  `json:"write_disposition,omitempty"`
	BeforeLoadQuery      *string  `json:"before_load_query,omitempty"`
	Partitioning         *string  `json:"partitioning,omitempty"`
	PartitioningTime     *string  `json:"partitioning_time,omitempty"`
	PartitioningField    *string  `json:"partitioning_field,omitempty"`
	ClusteringFields     []string `json:"clustering_fields,omitempty"`
	Location             *string  `json:"location,omitempty"`
}

type CreateDatamartDefinitionOutput struct {
	ID int64 `json:"id"`
}

func (client *TroccoClient) CreateDatamartDefinition(input *CreateDatamartDefinitionsInput) (*CreateDatamartDefinitionOutput, error) {
	url := fmt.Sprintf("%s/api/datamart_definitions", client.Endpoint)
	req, err := client.NewRequest("POST", url, input)
	if err != nil {
		return nil, err
	}
	output := new(CreateDatamartDefinitionOutput)
	err = client.Do(req, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type UpdateDatamartDefinitionsInput struct {
	Name                   *string                       `json:"name,omitempty"`
	Description            *string                       `json:"description,omitempty,omitempty"`
	IsRunnableConcurrently *bool                         `json:"is_runnable_concurrently,omitempty"`
	ResourceGroupID        *int64                        `json:"resource_group_id,omitempty"`
	CustomVariableSettings []UpdateCustomVariableSetting `json:"custom_variable_settings,omitempty"`
	DatamartBigqueryOption UpdateDatamartBigqueryOption  `json:"datamart_bigquery_option,omitempty"`
}
type UpdateCustomVariableSetting struct {
	Name string `json:"name"`
	Type string `json:"type"`
	CreateStringCustomVariableSetting
	CreateTimestampCustomVariableSetting
}
type UpdateStringCustomVariableSetting struct {
	Value *string `json:"value,omitempty"`
}
type UpdateTimestampCustomVariableSetting struct {
	Quantity  *int    `json:"quantity,omitempty"`
	Unit      *string `json:"unit,omitempty"`
	Direction *string `json:"direction,omitempty"`
	Format    *string `json:"format,omitempty"`
	TimeZone  *string `json:"time_zone,omitempty"`
}
type UpdateDatamartBigqueryOption struct {
	BigqueryConnectionID *int64   `json:"bigquery_connection_id,omitempty"`
	QueryMode            *string  `json:"query_mode,omitempty"`
	Query                *string  `json:"query,omitempty"`
	DestinationDataset   *string  `json:"destination_dataset,omitempty"`
	DestinationTable     *string  `json:"destination_table,omitempty"`
	WriteDisposition     *string  `json:"write_disposition,omitempty"`
	BeforeLoadQuery      *string  `json:"before_load_query,omitempty"`
	Partitioning         *string  `json:"partitioning,omitempty"`
	PartitioningTime     *string  `json:"partitioning_time,omitempty"`
	PartitioningField    *string  `json:"partitioning_field,omitempty"`
	ClusteringFields     []string `json:"clustering_fields,omitempty"`
	Location             *string  `json:"location,omitempty"`
}

func (client *TroccoClient) UpdateDatamartDefinition(id int64, input *UpdateDatamartDefinitionsInput) error {
	url := fmt.Sprintf("%s/api/datamart_definitions/%d", client.Endpoint, id)
	req, err := client.NewRequest("PATCH", url, input)
	if err != nil {
		return err
	}
	err = client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (client *TroccoClient) DeleteDatamartDefinition(id int64) error {
	url := fmt.Sprintf("%s/api/datamart_definitions/%d", client.Endpoint, id)
	req, err := client.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	err = client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}
