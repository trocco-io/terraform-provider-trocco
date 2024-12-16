package client

import (
	"encoding/json"
	"github.com/samber/lo"
	"net/http"
	"net/http/httptest"
	"terraform-provider-trocco/internal/client/entities"
	"terraform-provider-trocco/internal/client/entities/job_definitions"
	"terraform-provider-trocco/internal/client/entities/job_definitions/filter"
	filter2 "terraform-provider-trocco/internal/client/parameters/job_definitions/filter"
	"testing"

	"github.com/stretchr/testify/assert"
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
		Name:                      "name",
		Description:               lo.ToPtr("description"),
		ResourceGroupID:           lo.ToPtr(int64(9)),
		IsRunnableConcurrently:    lo.ToPtr(true),
		RetryLimit:                10,
		ResourceEnhancement:       lo.ToPtr("medium"),
		FilterColumns:             []filter2.FilterColumnInput{},
		FilterRows:                lo.ToPtr(filter2.FilterRowsInput{}),
		FilterMasks:               []filter2.FilterMaskInput{},
		FilterAddTime:             lo.ToPtr(filter2.FilterAddTimeInput{}),
		FilterGsub:                []filter2.FilterGsubInput{},
		FilterStringTransforms:    []filter2.FilterStringTransformInput{},
		FilterHashes:              []filter2.FilterHashInput{},
		FilterUnixTimeConversions: []filter2.FilterUnixTimeConversionInput{},
		InputOptionType:           "gcs",
		InputOption:               InputOptionInput{},
		OutputOptionType:          "mysql",
		OutputOption:              OutputOptionInput{},
		Labels:                    nil,
		Schedules:                 nil,
		Notifications:             nil,
	})

	assert.NoError(t, err)
	assert.Equal(t, int64(8), out.ID)
	assert.Equal(t, "name", out.Name)
	assert.Equal(t, "description", *out.Description)
	assert.Equal(t, int64(9), *out.ResourceGroupID)
	assert.Equal(t, true, *out.IsRunnableConcurrently)
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
	assert.Equal(t, []entities.Label(nil), out.Labels)
	assert.Equal(t, []entities.Schedule(nil), out.Schedules)
	assert.Equal(t, []job_definitions.JobDefinitionNotification(nil), out.Notifications)
}

//
//func TestUpdateJobDefinition(t *testing.T) {
//	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		assert.Equal(t, http.MethodPatch, r.Method)
//		assert.Equal(t, "/api/job_definitions/8", r.URL.Path)
//
//		w.Header().Set("Content-Type", "application/json")
//
//		jobDefinition := JobDefinition{
//			ID:                        8,
//			Name:                      "name",
//			Description:               lo.ToPtr("description"),
//			ResourceGroupID:           lo.ToPtr(int64(9)),
//			IsRunnableConcurrently:    lo.ToPtr(true),
//			RetryLimit:                10,
//			ResourceEnhancement:       lo.ToPtr("medium"),
//			FilterColumns:             []filter.FilterColumn{},
//			FilterRows:                lo.ToPtr(filter.FilterRows{}),
//			FilterMasks:               []filter.FilterMask{},
//			FilterAddTime:             lo.ToPtr(filter.FilterAddTime{}),
//			FilterGsub:                []filter.FilterGsub{},
//			FilterStringTransforms:    []filter.FilterStringTransform{},
//			FilterHashes:              []filter.FilterHash{},
//			FilterUnixTimeConversions: []filter.FilterUnixTimeConversion{},
//			InputOptionType:           "gcs",
//			InputOption:               InputOption{},
//			OutputOptionType:          "mysql",
//			OutputOption:              OutputOption{},
//			Labels:                    nil,
//			Schedules:                 nil,
//			Notifications:             nil,
//		}
//		if err := json.NewEncoder(w).Encode(jobDefinition); err != nil {
//			panic(err)
//		}
//	}))
//	defer s.Close()
//
//	c := NewDevTroccoClient("1234567890", s.URL)
//
//	out, err := c.UpdateConnection("bigquery", 8, &UpdateConnectionInput{
//		Name:        lo.ToPtr("Foo"),
//		Description: lo.ToPtr("The quick brown fox jumps over the lazy dog."),
//		ResourceGroupID: lo.ToPtr(NullableInt64{
//			Valid: true,
//			Value: int64(42),
//		}),
//	})
//
//	assert.NoError(t, err)
//
//	assert.Equal(t, int64(8), out.ID)
//	assert.Equal(t, "Foo", *out.Name)
//	assert.Equal(t, "The quick brown fox jumps over the lazy dog.", *out.Description)
//	assert.Equal(t, int64(42), *out.ResourceGroupID)
//}
