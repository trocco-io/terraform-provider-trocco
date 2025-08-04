package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"terraform-provider-trocco/internal/client/entity"
	jobDefinitionEntities "terraform-provider-trocco/internal/client/entity/job_definition"
	"terraform-provider-trocco/internal/client/entity/job_definition/filter"
	"terraform-provider-trocco/internal/client/parameter"
	filterParameters "terraform-provider-trocco/internal/client/parameter/job_definition/filter"
	"testing"

	"github.com/samber/lo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteJobDefinition(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check request method and path are correct.
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/api/job_definitions/8", r.URL.Path)
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)

	err := c.DeleteJobDefinition(8)

	assert.NoError(t, err)
}

func TestCreateJobDefinition(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/job_definitions", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")

		jobDefinition := JobDefinition{
			ID:                        8,
			Name:                      "name",
			Description:               lo.ToPtr("description"),
			ResourceGroupID:           lo.ToPtr(int64(9)),
			IsRunnableConcurrently:    lo.ToPtr(true),
			RetryLimit:                10,
			ResourceEnhancement:       lo.ToPtr("medium"),
			FilterColumns:             []filter.FilterColumn{},
			FilterRows:                lo.ToPtr(filter.FilterRows{}),
			FilterMasks:               []filter.FilterMask{},
			FilterAddTime:             lo.ToPtr(filter.FilterAddTime{}),
			FilterGsub:                []filter.FilterGsub{},
			FilterStringTransforms:    []filter.FilterStringTransform{},
			FilterHashes:              []filter.FilterHash{},
			FilterUnixTimeConversions: []filter.FilterUnixTimeConversion{},
			InputOptionType:           "gcs",
			InputOption:               InputOption{},
			OutputOptionType:          "mysql",
			OutputOption:              OutputOption{},
			Labels:                    nil,
			Schedules:                 nil,
			Notifications:             nil,
		}
		if err := json.NewEncoder(w).Encode(jobDefinition); err != nil {
			panic(err)
		}
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)
	out, err := c.CreateJobDefinition(&CreateJobDefinitionInput{
		Name:                   "name",
		Description:            lo.ToPtr(parameter.NullableString{Value: "description", Valid: true}),
		ResourceGroupID:        lo.ToPtr(parameter.NullableInt64{Value: 9, Valid: true}),
		IsRunnableConcurrently: true,
		RetryLimit:             10,
		ResourceEnhancement:    lo.ToPtr("medium"),
		FilterColumns:          []filterParameters.FilterColumnInput{},
		FilterRows: lo.ToPtr(parameter.NullableObject[filterParameters.FilterRowsInput]{
			Valid: true,
			Value: &filterParameters.FilterRowsInput{
				Condition:           "or",
				FilterRowConditions: make([]filterParameters.FilterRowConditionInput, 0),
			},
		}),
		FilterMasks: []filterParameters.FilterMaskInput{},
		FilterAddTime: lo.ToPtr(parameter.NullableObject[filterParameters.FilterAddTimeInput]{
			Valid: true,
			Value: &filterParameters.FilterAddTimeInput{
				ColumnName:      "col_name",
				Type:            "string",
				TimestampFormat: nil,
				TimeZone:        nil,
			},
		}),
		FilterGsub:                []filterParameters.FilterGsubInput{},
		FilterStringTransforms:    []filterParameters.FilterStringTransformInput{},
		FilterHashes:              []filterParameters.FilterHashInput{},
		FilterUnixTimeConversions: []filterParameters.FilterUnixTimeConversionInput{},
		InputOptionType:           "gcs",
		InputOption:               InputOptionInput{},
		OutputOptionType:          "mysql",
		OutputOption:              OutputOptionInput{},
		Labels:                    nil,
		Schedules:                 nil,
		Notifications:             nil,
	})

	require.NoError(t, err)
	assert.Equal(t, int64(8), out.ID)
	assert.Equal(t, "name", out.Name)
	assert.Equal(t, "description", *out.Description)
	assert.Equal(t, int64(9), *out.ResourceGroupID)
	assert.True(t, *out.IsRunnableConcurrently)
	assert.Equal(t, int64(10), out.RetryLimit)
	assert.Equal(t, "medium", *out.ResourceEnhancement)
	assert.Equal(t, []filter.FilterColumn{}, out.FilterColumns)
	assert.Equal(t, filter.FilterRows{}, *out.FilterRows)
	assert.Equal(t, []filter.FilterMask{}, out.FilterMasks)
	assert.Equal(t, filter.FilterAddTime{}, *out.FilterAddTime)
	assert.Equal(t, []filter.FilterGsub{}, out.FilterGsub)
	assert.Equal(t, []filter.FilterStringTransform{}, out.FilterStringTransforms)
	assert.Equal(t, []filter.FilterHash{}, out.FilterHashes)
	assert.Equal(t, []filter.FilterUnixTimeConversion{}, out.FilterUnixTimeConversions)
	assert.Equal(t, "gcs", out.InputOptionType)
	assert.Equal(t, "mysql", out.OutputOptionType)
	assert.Equal(t, InputOption{}, out.InputOption)
	assert.Equal(t, OutputOption{}, out.OutputOption)
	assert.Equal(t, []entity.Label(nil), out.Labels)
	assert.Equal(t, []entity.Schedule(nil), out.Schedules)
	assert.Equal(t, []jobDefinitionEntities.JobDefinitionNotification(nil), out.Notifications)
}

func TestUpdateJobDefinition(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		assert.Equal(t, "/api/job_definitions/8", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")

		jobDefinition := JobDefinition{
			ID:                        8,
			Name:                      "edit",
			Description:               lo.ToPtr("description edit"),
			ResourceGroupID:           lo.ToPtr(int64(10)),
			IsRunnableConcurrently:    lo.ToPtr(true),
			RetryLimit:                11,
			ResourceEnhancement:       lo.ToPtr("medium"),
			FilterColumns:             []filter.FilterColumn{},
			FilterRows:                lo.ToPtr(filter.FilterRows{}),
			FilterMasks:               []filter.FilterMask{},
			FilterAddTime:             lo.ToPtr(filter.FilterAddTime{}),
			FilterGsub:                []filter.FilterGsub{},
			FilterStringTransforms:    []filter.FilterStringTransform{},
			FilterHashes:              []filter.FilterHash{},
			FilterUnixTimeConversions: []filter.FilterUnixTimeConversion{},
			InputOptionType:           "gcs",
			InputOption:               InputOption{},
			OutputOptionType:          "mysql",
			OutputOption:              OutputOption{},
			Labels:                    nil,
			Schedules:                 nil,
			Notifications:             nil,
		}
		if err := json.NewEncoder(w).Encode(jobDefinition); err != nil {
			panic(err)
		}
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)

	out, err := c.UpdateJobDefinition(8, &UpdateJobDefinitionInput{
		Name:                   lo.ToPtr("edit"),
		Description:            lo.ToPtr(parameter.NullableString{Value: "description edit", Valid: true}),
		ResourceGroupID:        lo.ToPtr(parameter.NullableInt64{Value: 10, Valid: true}),
		IsRunnableConcurrently: lo.ToPtr(true),
		RetryLimit:             lo.ToPtr(int64(11)),
		ResourceEnhancement:    lo.ToPtr("medium"),
	})

	require.NoError(t, err)
	assert.Equal(t, int64(8), out.ID)
	assert.Equal(t, "edit", out.Name)
	assert.Equal(t, "description edit", *out.Description)
	assert.Equal(t, int64(10), *out.ResourceGroupID)
	assert.True(t, *out.IsRunnableConcurrently)
	assert.Equal(t, int64(11), out.RetryLimit)
	assert.Equal(t, "medium", *out.ResourceEnhancement)
	assert.Equal(t, []filter.FilterColumn{}, out.FilterColumns)
	assert.Equal(t, filter.FilterRows{}, *out.FilterRows)
	assert.Equal(t, []filter.FilterMask{}, out.FilterMasks)
	assert.Equal(t, filter.FilterAddTime{}, *out.FilterAddTime)
	assert.Equal(t, []filter.FilterGsub{}, out.FilterGsub)
	assert.Equal(t, []filter.FilterStringTransform{}, out.FilterStringTransforms)
	assert.Equal(t, []filter.FilterHash{}, out.FilterHashes)
	assert.Equal(t, []filter.FilterUnixTimeConversion{}, out.FilterUnixTimeConversions)
	assert.Equal(t, "gcs", out.InputOptionType)
	assert.Equal(t, "mysql", out.OutputOptionType)
	assert.Equal(t, InputOption{}, out.InputOption)
	assert.Equal(t, OutputOption{}, out.OutputOption)
	assert.Equal(t, []entity.Label(nil), out.Labels)
	assert.Equal(t, []entity.Schedule(nil), out.Schedules)
	assert.Equal(t, []jobDefinitionEntities.JobDefinitionNotification(nil), out.Notifications)
}

func TestGetJobDefinition(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/job_definitions/8", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")

		c := JobDefinition{
			ID:                        8,
			Name:                      "new",
			Description:               lo.ToPtr("description new"),
			ResourceGroupID:           lo.ToPtr(int64(10)),
			IsRunnableConcurrently:    lo.ToPtr(true),
			RetryLimit:                11,
			ResourceEnhancement:       lo.ToPtr("medium"),
			FilterColumns:             []filter.FilterColumn{},
			FilterRows:                lo.ToPtr(filter.FilterRows{}),
			FilterMasks:               []filter.FilterMask{},
			FilterAddTime:             lo.ToPtr(filter.FilterAddTime{}),
			FilterGsub:                []filter.FilterGsub{},
			FilterStringTransforms:    []filter.FilterStringTransform{},
			FilterHashes:              []filter.FilterHash{},
			FilterUnixTimeConversions: []filter.FilterUnixTimeConversion{},
			InputOptionType:           "gcs",
			InputOption:               InputOption{},
			OutputOptionType:          "mysql",
			OutputOption:              OutputOption{},
			Labels:                    nil,
			Schedules:                 nil,
			Notifications:             nil,
		}
		if err := json.NewEncoder(w).Encode(c); err != nil {
			panic(err)
		}
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)

	out, err := c.GetJobDefinition(8)

	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.Equal(t, int64(8), out.ID)
	assert.Equal(t, "new", out.Name)
	assert.Equal(t, "description new", *out.Description)
	assert.Equal(t, int64(10), *out.ResourceGroupID)
	assert.True(t, *out.IsRunnableConcurrently)
	assert.Equal(t, int64(11), out.RetryLimit)
	assert.Equal(t, "medium", *out.ResourceEnhancement)
	assert.Equal(t, []filter.FilterColumn{}, out.FilterColumns)
	assert.Equal(t, filter.FilterRows{}, *out.FilterRows)
	assert.Equal(t, []filter.FilterMask{}, out.FilterMasks)
	assert.Equal(t, filter.FilterAddTime{}, *out.FilterAddTime)
	assert.Equal(t, []filter.FilterGsub{}, out.FilterGsub)
	assert.Equal(t, []filter.FilterStringTransform{}, out.FilterStringTransforms)
	assert.Equal(t, []filter.FilterHash{}, out.FilterHashes)
	assert.Equal(t, []filter.FilterUnixTimeConversion{}, out.FilterUnixTimeConversions)
	assert.Equal(t, "gcs", out.InputOptionType)
	assert.Equal(t, "mysql", out.OutputOptionType)
	assert.Equal(t, InputOption{}, out.InputOption)
	assert.Equal(t, OutputOption{}, out.OutputOption)
	assert.Equal(t, []entity.Label(nil), out.Labels)
	assert.Equal(t, []entity.Schedule(nil), out.Schedules)
	assert.Equal(t, []jobDefinitionEntities.JobDefinitionNotification(nil), out.Notifications)
}
