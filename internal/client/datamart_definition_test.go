package client

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// Helpers

type Case struct {
	name     string
	value    interface{}
	expected interface{}
}

func testCases(t *testing.T, cases []Case) {
	for _, c := range cases {
		value := c.value
		if c.expected == nil {
			if !reflect.ValueOf(value).IsNil() {
				t.Errorf("Expected %s to be nil, got %v", c.name, value)
			}
			continue
		}
		if reflect.ValueOf(value).Kind() == reflect.Ptr {
			value = reflect.ValueOf(value).Elem().Interface()
		}
		if c.expected != value {
			t.Errorf("Expected %s to be %v, got %v", c.name, c.expected, value)
		}
	}
}

// ListDatamartDefinitions

func TestListDatamartDefinitions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cases := []Case{
			{"path", r.URL.Path, "/api/datamart_definitions"},
			{"method", r.Method, http.MethodGet},
		}
		testCases(t, cases)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := `
      {
        "items": [
          {
            "id": 1,
            "name": "Test Datamart 01",
            "description": "This is a first test datamart",
            "data_warehouse_type": "bigquery",
            "created_at": "2024-07-29T19:00:00.000+09:00",
            "updated_at": "2024-07-29T20:00:00.000+09:00"
          },
          {
            "id": 2,
            "name": "Test Datamart 02",
            "description": "This is a second test datamart",
            "data_warehouse_type": "snowflake",
            "created_at": "2024-07-29T21:00:00.000+09:00",
            "updated_at": "2024-07-29T22:00:00.000+09:00"
          }
        ]
      }
    `
		_, err := w.Write([]byte(resp))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	output, err := client.ListDatamartDefinitions(nil)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if len(output.Items) != 2 {
		t.Errorf("Expected output.Items to have length 2, got %d", len(output.Items))
	}
	cases := []Case{
		{"first item's ID", output.Items[0].ID, int64(1)},
		{"first item's Name", output.Items[0].Name, "Test Datamart 01"},
		{"first item's Description", output.Items[0].Description, "This is a first test datamart"},
		{"first item's DataWarehouseType", output.Items[0].DataWarehouseType, "bigquery"},
		{"first item's CreatedAt", output.Items[0].CreatedAt, "2024-07-29T19:00:00.000+09:00"},
		{"first item's UpdatedAt", output.Items[0].UpdatedAt, "2024-07-29T20:00:00.000+09:00"},
		{"second item's ID", output.Items[1].ID, int64(2)},
		{"second item's Name", output.Items[1].Name, "Test Datamart 02"},
		{"second item's Description", output.Items[1].Description, "This is a second test datamart"},
		{"second item's DataWarehouseType", output.Items[1].DataWarehouseType, "snowflake"},
		{"second item's CreatedAt", output.Items[1].CreatedAt, "2024-07-29T21:00:00.000+09:00"},
		{"second item's UpdatedAt", output.Items[1].UpdatedAt, "2024-07-29T22:00:00.000+09:00"},
	}
	testCases(t, cases)
}

func TestListDatamartDefinitionsLimitAndCursor(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cases := []Case{
			{"query parameter limit", r.URL.Query().Get("limit"), "1"},
			{"query parameter cursor", r.URL.Query().Get("cursor"), "test_prev_cursor"},
		}
		testCases(t, cases)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := `
      {
        "items": [],
        "next_cursor": "test_next_cursor"
      }
    `
		_, err := w.Write([]byte(resp))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	input := ListDatamartDefinitionsInput{}
	input.SetLimit(1)
	input.SetCursor("test_prev_cursor")
	output, err := client.ListDatamartDefinitions(&input)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	cases := []Case{
		{"next_cursor", *output.NextCursor, "test_next_cursor"},
	}
	testCases(t, cases)
}

// GetDatamartDefinition

func TestGetDatamartDefinitionMinimum(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cases := []Case{
			{"path", r.URL.Path, "/api/datamart_definitions/1"},
			{"method", r.Method, http.MethodGet},
		}
		testCases(t, cases)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := `
      {
        "id": 1,
        "name": "Test Datamart 01",
        "description": "",
        "data_warehouse_type": "bigquery",
        "datamart_bigquery_option": {
          "bigquery_connection_id": 1,
          "query_mode": "insert",
          "query": "SELECT * FROM table",
          "destination_dataset": "test_dataset",
          "destination_table": "test_table",
          "write_disposition": "truncate",
          "partitioning": ""
        },
        "created_at": "2024-07-29T19:00:00.000+09:00",
        "updated_at": "2024-07-29T20:00:00.000+09:00"
      }
    `
		_, err := w.Write([]byte(resp))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	output, err := client.GetDatamartDefinition(1)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	cases := []Case{
		{"id", output.ID, int64(1)},
		{"name", output.Name, "Test Datamart 01"},
		{"description", output.Description, nil},
		{"data_warehouse_type", output.DataWarehouseType, "bigquery"},
		{"datamart_bigquery_option.bigquery_connection_id", output.DatamartBigqueryOption.BigqueryConnectionID, int64(1)},
		{"datamart_bigquery_option.query_mode", output.DatamartBigqueryOption.QueryMode, "insert"},
		{"datamart_bigquery_option.query", output.DatamartBigqueryOption.Query, "SELECT * FROM table"},
		{"datamart_bigquery_option.destination_dataset", *output.DatamartBigqueryOption.DestinationDataset, "test_dataset"},
		{"datamart_bigquery_option.destination_table", *output.DatamartBigqueryOption.DestinationTable, "test_table"},
		{"datamart_bigquery_option.write_disposition", *output.DatamartBigqueryOption.WriteDisposition, "truncate"},
		{"datamart_bigquery_option.partitioning", output.DatamartBigqueryOption.Partitioning, nil},
		{"created_at", output.CreatedAt, "2024-07-29T19:00:00.000+09:00"},
		{"updated_at", output.UpdatedAt, "2024-07-29T20:00:00.000+09:00"},
	}
	testCases(t, cases)
}

func TestGetDatamartDefinitionQueryMode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := `
      {
        "id": 1,
        "name": "Test Datamart 01",
        "description": "",
        "data_warehouse_type": "bigquery",
        "datamart_bigquery_option": {
          "bigquery_connection_id": 1,
          "query_mode": "query",
          "query": "SELECT * FROM table",
          "location": "asia-northeast1"
        },
        "created_at": "2024-07-29T19:00:00.000+09:00",
        "updated_at": "2024-07-29T20:00:00.000+09:00"
      }
    `
		_, err := w.Write([]byte(resp))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	output, err := client.GetDatamartDefinition(1)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	cases := []Case{
		{"datamart_bigquery_option.query_mode", output.DatamartBigqueryOption.QueryMode, "query"},
		{"datamart_bigquery_option.location", *output.DatamartBigqueryOption.Location, "asia-northeast1"},
	}
	testCases(t, cases)
}

func TestGetDatamartDefinitionFull(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := `
      {
        "id": 1,
        "name": "Test Datamart 01",
        "description": "This is a first test datamart",
        "data_warehouse_type": "bigquery",
        "resource_group": {
          "id": 1,
          "name": "test_resource_group",
          "description": "This is a test resource group",
          "created_at": "2024-07-29T19:00:00.000+09:00",
          "updated_at": "2024-07-29T20:00:00.000+09:00"
        },
        "custom_variable_settings": [
          {
            "name": "$string$",
            "type": "string",
            "value": "foo"
          },
          {
            "name": "$timestamp$",
            "type": "timestamp",
            "quantity": 1,
            "unit": "hour",
            "direction": "ago",
            "format": "%Y-%m-%d %H:%M:%S",
            "time_zone": "Asia/Tokyo"
          }
        ],
        "datamart_bigquery_option": {
          "bigquery_connection_id": 1,
          "query_mode": "insert",
          "query": "SELECT * FROM table",
          "destination_dataset": "test_dataset",
          "destination_table": "test_table",
          "write_disposition": "truncate",
          "before_load": "DELETE FROM table WHERE id = 1",
          "partitioning": "time_unit_column",
          "partitioning_time": "HOUR",
          "partitioning_field": "created_at",
          "clustering_fields": ["id", "name"]
        },
        "created_at": "2024-07-29T19:00:00.000+09:00",
        "updated_at": "2024-07-29T20:00:00.000+09:00",
        "notifications": [
          {
            "destination_type": "slack",
            "slack_channel_id": 1,
            "notification_type": "job",
            "notify_when": "finished",
            "message": "foo"
          },
          {
            "destination_type": "email",
            "email_id": 1,
            "notification_type": "record",
            "record_count": 100,
            "record_operator": "below",
            "message": "bar"
          }
        ],
        "schedules": [
          {
            "frequency": "hourly",
            "minute": 1,
            "time_zone": "Asia/Tokyo"
          },
          {
            "frequency": "daily",
            "minute": 1,
            "hour": 2,
            "time_zone": "Asia/Tokyo"
          },
          {
            "frequency": "weekly",
            "minute": 1,
            "hour": 2,
            "day_of_week": 3,
            "time_zone": "Asia/Tokyo"
          },
          {
            "frequency": "monthly",
            "minute": 1,
            "hour": 2,
            "day": 4,
            "time_zone": "Asia/Tokyo"
          }
        ],
        "labels": [
          {
            "id": 1,
            "name": "test_label",
            "description": "This is a test label",
            "color": "#FF3B1D",
            "created_at": "2024-07-29T19:00:00.000+09:00",
            "updated_at": "2024-07-29T20:00:00.000+09:00"
          }
        ]
      }
    `
		_, err := w.Write([]byte(resp))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	output, err := client.GetDatamartDefinition(1)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	cases := []Case{
		{"description", *output.Description, "This is a first test datamart"},
		{"datamart_bigquery_option.before_load", *output.DatamartBigqueryOption.BeforeLoad, "DELETE FROM table WHERE id = 1"},
		{"datamart_bigquery_option.partitioning", *output.DatamartBigqueryOption.Partitioning, "time_unit_column"},
		{"datamart_bigquery_option.partitioning_time", *output.DatamartBigqueryOption.PartitioningTime, "HOUR"},
		{"datamart_bigquery_option.partitioning_field", *output.DatamartBigqueryOption.PartitioningField, "created_at"},
		{"first datamart_bigquery_option's clustering_fields", output.DatamartBigqueryOption.ClusteringFields[0], "id"},
		{"second datamart_bigquery_option's clustering_fields", output.DatamartBigqueryOption.ClusteringFields[1], "name"},

		{"resource_group's id", output.ResourceGroup.ID, int64(1)},
		{"resource_group's name", output.ResourceGroup.Name, "test_resource_group"},
		{"resource_group's description", output.ResourceGroup.Description, "This is a test resource group"},
		{"resource_group's created_at", output.ResourceGroup.CreatedAt, "2024-07-29T19:00:00.000+09:00"},
		{"resource_group's updated_at", output.ResourceGroup.UpdatedAt, "2024-07-29T20:00:00.000+09:00"},

		{"first custom_variable_settings's name", output.CustomVariableSettings[0].Name, "$string$"},
		{"first custom_variable_settings's type", output.CustomVariableSettings[0].Type, "string"},
		{"first custom_variable_settings's value", *output.CustomVariableSettings[0].Value, "foo"},
		{"second custom_variable_settings's name", output.CustomVariableSettings[1].Name, "$timestamp$"},
		{"second custom_variable_settings's type", output.CustomVariableSettings[1].Type, "timestamp"},
		{"second custom_variable_settings's quantity", *output.CustomVariableSettings[1].Quantity, 1},
		{"second custom_variable_settings's unit", *output.CustomVariableSettings[1].Unit, "hour"},
		{"second custom_variable_settings's direction", *output.CustomVariableSettings[1].Direction, "ago"},
		{"second custom_variable_settings's format", *output.CustomVariableSettings[1].Format, "%Y-%m-%d %H:%M:%S"},
		{"second custom_variable_settings's time_zone", *output.CustomVariableSettings[1].TimeZone, "Asia/Tokyo"},

		{"first schedules's frequency", output.Schedules[0].Frequency, "hourly"},
		{"first schedules's minute", output.Schedules[0].Minute, 1},
		{"first schedules's time_zone", output.Schedules[0].TimeZone, "Asia/Tokyo"},
		{"second schedules's frequency", output.Schedules[1].Frequency, "daily"},
		{"second schedules's minute", output.Schedules[1].Minute, 1},
		{"second schedules's hour", output.Schedules[1].Hour, 2},
		{"second schedules's time_zone", output.Schedules[1].TimeZone, "Asia/Tokyo"},
		{"third schedules's frequency", output.Schedules[2].Frequency, "weekly"},
		{"third schedules's minute", output.Schedules[2].Minute, 1},
		{"third schedules's hour", output.Schedules[2].Hour, 2},
		{"third schedules's day_of_week", output.Schedules[2].DayOfWeek, 3},
		{"third schedules's time_zone", output.Schedules[2].TimeZone, "Asia/Tokyo"},
		{"fourth schedules's frequency", output.Schedules[3].Frequency, "monthly"},
		{"fourth schedules's minute", output.Schedules[3].Minute, 1},
		{"fourth schedules's hour", output.Schedules[3].Hour, 2},
		{"fourth schedules's day", output.Schedules[3].Day, 4},
		{"fourth schedules's time_zone", output.Schedules[3].TimeZone, "Asia/Tokyo"},

		{"first notifications's destination_type", output.Notifications[0].DestinationType, "slack"},
		{"first notifications's slack_channel_id", output.Notifications[0].SlackChannelID, int64(1)},
		{"first notifications's notification_type", output.Notifications[0].NotificationType, "job"},
		{"first notifications's notify_when", output.Notifications[0].NotifyWhen, "finished"},
		{"first notifications's message", output.Notifications[0].Message, "foo"},
		{"second notifications's destination_type", output.Notifications[1].DestinationType, "email"},
		{"second notifications's email_id", output.Notifications[1].EmailID, int64(1)},
		{"second notifications's notification_type", output.Notifications[1].NotificationType, "record"},
		{"second notifications's record_count", *output.Notifications[1].RecordCount, int64(100)},
		{"second notifications's record_operator", *output.Notifications[1].RecordOperator, "below"},
		{"second notifications's message", output.Notifications[1].Message, "bar"},

		{"first labels's id", output.Labels[0].ID, int64(1)},
		{"first labels's name", output.Labels[0].Name, "test_label"},
		{"first labels's description", output.Labels[0].Description, "This is a test label"},
		{"first labels's color", output.Labels[0].Color, "#FF3B1D"},
		{"first labels's created_at", output.Labels[0].CreatedAt, "2024-07-29T19:00:00.000+09:00"},
		{"first labels's updated_at", output.Labels[0].UpdatedAt, "2024-07-29T20:00:00.000+09:00"},
	}
	testCases(t, cases)
}

// CreateDatamartDefinition

func TestCreateDatamartDefinitionMinimum(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cases := []Case{
			{"path", r.URL.Path, "/api/datamart_definitions"},
			{"method", r.Method, http.MethodPost},
		}
		testCases(t, cases)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if string(body) != `{"name":"Test Datamart 01","data_warehouse_type":"bigquery","is_runnable_concurrently":false,"datamart_bigquery_option":{"bigquery_connection_id":1,"query_mode":"query","query":"SELECT * FROM table","location":"asia-northeast1"}}` {
			t.Errorf("Not expected request body: %s", string(body))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte(`{"id":1}`))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	input := NewCreateDatamartDefinitionInput(
		"Test Datamart 01",
		"bigquery",
		false,
	)
	bigqueryOption := NewQueryModeCreateDatamartBigqueryOptionInput(
		1,
		"SELECT * FROM table",
	)
	bigqueryOption.SetLocation("asia-northeast1")
	input.SetDatamartBigqueryOption(bigqueryOption)
	output, err := client.CreateDatamartDefinition(&input)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if output.ID != 1 {
		t.Errorf("Expected output.ID to be 1, got %d", output.ID)
	}
}

func TestCreateDatamartDefinitionFull(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if string(body) != `{"name":"Test Datamart 01","data_warehouse_type":"bigquery","description":"This is a first test datamart","is_runnable_concurrently":false,"resource_group_id":1,"custom_variable_settings":[{"name":"$string$","type":"string","value":"foo"},{"name":"$timestamp$","type":"timestamp","quantity":1,"unit":"hour","direction":"ago","format":"%Y-%m-%d %H:%M:%S","time_zone":"Asia/Tokyo"}],"datamart_bigquery_option":{"bigquery_connection_id":1,"query_mode":"insert","query":"SELECT * FROM table","destination_dataset":"test_dataset","destination_table":"test_table","write_disposition":"truncate","before_load":"DELETE FROM table WHERE id = 1","partitioning":"time_unit_column","partitioning_time":"HOUR","partitioning_field":"created_at","clustering_fields":["id","name"]}}` {
			t.Errorf("Not expected request body: %s", string(body))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte(`{"id":1}`))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	input := NewCreateDatamartDefinitionInput(
		"Test Datamart 01",
		"bigquery",
		false,
	)
	input.SetDescription("This is a first test datamart")
	input.SetResourceGroupID(1)
	input.SetCustomVariableSettings([]CustomVariableSettingInput{
		NewStringTypeCustomVariableSettingInput(
			"$string$",
			"foo",
		),
		NewTimestampTypeCustomVariableSettingInput(
			"$timestamp$",
			"timestamp",
			1,
			"hour",
			"ago",
			"%Y-%m-%d %H:%M:%S",
			"Asia/Tokyo",
		),
	})
	bigqueryOption := NewInsertModeCreateDatamartBigqueryOptionInput(
		1,
		"SELECT * FROM table",
		"test_dataset",
		"test_table",
		"truncate",
	)
	bigqueryOption.SetBeforeLoad("DELETE FROM table WHERE id = 1")
	bigqueryOption.SetPartitioning("time_unit_column")
	bigqueryOption.SetPartitioningTime("HOUR")
	bigqueryOption.SetPartitioningField("created_at")
	bigqueryOption.SetClusteringFields([]string{"id", "name"})
	input.SetDatamartBigqueryOption(bigqueryOption)

	_, err := client.CreateDatamartDefinition(&input)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}

// UpdateDatamartDefinition

func TestUpdateDatamartDefinitionWithBasicValues(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cases := []Case{
			{"path", r.URL.Path, "/api/datamart_definitions/1"},
			{"method", r.Method, http.MethodPatch},
		}
		testCases(t, cases)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if string(body) != `{"name":"Test Datamart 01","description":"This is a first test datamart","is_runnable_concurrently":false,"resource_group_id":1}` {
			t.Errorf("Not expected request body: %s", string(body))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("{}"))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	input := UpdateDatamartDefinitionInput{}
	input.SetName("Test Datamart 01")
	input.SetDescription("This is a first test datamart")
	input.SetIsRunnableConcurrently(false)
	input.SetResourceGroupID(1)
	err := client.UpdateDatamartDefinition(1, &input)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}

func TestUpdateDatamartDefinitionWithCustomVariableSettings(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if string(body) != `{"custom_variable_settings":[{"name":"$string$","type":"string","value":"foo"},{"name":"$timestamp$","type":"timestamp","quantity":1,"unit":"hour","direction":"ago","format":"%Y-%m-%d %H:%M:%S","time_zone":"Asia/Tokyo"}]}` {
			t.Errorf("Not expected request body: %s", string(body))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("{}"))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	input := UpdateDatamartDefinitionInput{}
	input.SetCustomVariableSettings([]CustomVariableSettingInput{
		NewStringTypeCustomVariableSettingInput(
			"$string$",
			"foo",
		),
		NewTimestampTypeCustomVariableSettingInput(
			"$timestamp$",
			"timestamp",
			1,
			"hour",
			"ago",
			"%Y-%m-%d %H:%M:%S",
			"Asia/Tokyo",
		),
	})
	err := client.UpdateDatamartDefinition(1, &input)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}

func TestUpdateDatamartDefinitionWithDatamartBigqueryOptionQueryMode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if string(body) != `{"datamart_bigquery_option":{"bigquery_connection_id":1,"query_mode":"query","query":"SELECT * FROM table","location":"asia-northeast1"}}` {
			t.Errorf("Not expected request body: %s", string(body))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("{}"))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	input := UpdateDatamartDefinitionInput{}
	bigqueryOption := UpdateDatamartBigqueryOptionInput{}
	bigqueryOption.SetBigqueryConnectionID(1)
	bigqueryOption.SetQueryMode("query")
	bigqueryOption.SetQuery("SELECT * FROM table")
	bigqueryOption.SetLocation("asia-northeast1")
	input.SetDatamartBigqueryOption(bigqueryOption)
	err := client.UpdateDatamartDefinition(1, &input)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}

func TestUpdateDatamartDefinitionWithDatamartBigqueryOptionInsertMode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if string(body) != `{"datamart_bigquery_option":{"destination_dataset":"test_dataset","destination_table":"test_table","write_disposition":"truncate","before_load":"DELETE FROM table WHERE id = 1","partitioning":"time_unit_column","partitioning_time":"HOUR","partitioning_field":"created_at","clustering_fields":["id","name"]}}` {
			t.Errorf("Not expected request body: %s", string(body))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("{}"))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	input := UpdateDatamartDefinitionInput{}
	bigqueryOption := UpdateDatamartBigqueryOptionInput{}
	bigqueryOption.SetDestinationDataset("test_dataset")
	bigqueryOption.SetDestinationTable("test_table")
	bigqueryOption.SetWriteDisposition("truncate")
	bigqueryOption.SetBeforeLoad("DELETE FROM table WHERE id = 1")
	bigqueryOption.SetPartitioning("time_unit_column")
	bigqueryOption.SetPartitioningTime("HOUR")
	bigqueryOption.SetPartitioningField("created_at")
	bigqueryOption.SetClusteringFields([]string{"id", "name"})
	input.SetDatamartBigqueryOption(bigqueryOption)
	err := client.UpdateDatamartDefinition(1, &input)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}

func TestUpdateDatamartDefinitionWithSchedules(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if string(body) != `{"schedules":[{"frequency":"hourly","minute":1,"time_zone":"Asia/Tokyo"},{"frequency":"daily","minute":1,"hour":2,"time_zone":"Asia/Tokyo"},{"frequency":"weekly","minute":1,"hour":2,"day_of_week":3,"time_zone":"Asia/Tokyo"},{"frequency":"monthly","minute":1,"hour":2,"day":4,"time_zone":"Asia/Tokyo"}]}` {
			t.Errorf("Not expected request body: %s", string(body))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("{}"))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	input := UpdateDatamartDefinitionInput{}
	input.SetSchedules([]ScheduleInput{
		NewHourlyScheduleInput(1, "Asia/Tokyo"),
		NewDailyScheduleInput(2, 1, "Asia/Tokyo"),
		NewWeeklyScheduleInput(3, 2, 1, "Asia/Tokyo"),
		NewMonthlyScheduleInput(4, 2, 1, "Asia/Tokyo"),
	})
	err := client.UpdateDatamartDefinition(1, &input)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}

func TestUpdateDatamartDefinitionWithNotifications(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if string(body) != `{"notifications":[{"destination_type":"slack","slack_channel_id":1,"notification_type":"job","notify_when":"finished","message":"foo"},{"destination_type":"email","email_id":1,"notification_type":"record","record_count":100,"record_operator":"below","message":"bar"}]}` {
			t.Errorf("Not expected request body: %s", string(body))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("{}"))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	input := UpdateDatamartDefinitionInput{}
	input.SetNotifications([]DatamartNotificationInput{
		NewSlackJobDatamartNotificationInput(1, "finished", "foo"),
		NewEmailRecordDatamartNotificationInput(1, 100, "below", "bar"),
	})
	err := client.UpdateDatamartDefinition(1, &input)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}

func TestUpdateDatamartDefinitionWithLabels(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if string(body) != `{"labels":["test_label_1","test_label_2"]}` {
			t.Errorf("Not expected request body: %s", string(body))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("{}"))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	input := UpdateDatamartDefinitionInput{}
	input.SetLabels([]string{
		"test_label_1",
		"test_label_2",
	})
	err := client.UpdateDatamartDefinition(1, &input)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}

func TestUpdateDatamartDefinitionWithEmptyValues(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if string(body) != `{"description":"","resource_group_id":null,"datamart_bigquery_option":{"before_load":null,"partitioning":null,"location":null}}` {
			t.Errorf("Not expected request body: %s", string(body))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("{}"))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	input := UpdateDatamartDefinitionInput{}
	input.SetDescriptionEmpty()
	input.SetResourceGroupIDEmpty()
	bigqueryOption := UpdateDatamartBigqueryOptionInput{}
	bigqueryOption.SetBeforeLoadEmpty()
	bigqueryOption.SetPartitioningEmpty()
	bigqueryOption.SetLocationEmpty()
	input.SetDatamartBigqueryOption(bigqueryOption)
	err := client.UpdateDatamartDefinition(1, &input)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}

func TestDeleteDatamartDefinition(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cases := []Case{
			{"path", r.URL.Path, "/api/datamart_definitions/1"},
			{"method", r.Method, http.MethodDelete},
		}
		testCases(t, cases)

		w.WriteHeader(http.StatusNoContent)
		_, err := w.Write([]byte(""))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	err := client.DeleteDatamartDefinition(1)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}
