package output_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type HubspotOutputOptionInput struct {
	HubspotConnectionID int64                                                        `json:"hubspot_connection_id"`
	ObjectType          string                                                       `json:"object_type"`
	Mode                string                                                       `json:"mode"`
	UpsertKey           *parameter.NullableString                                    `json:"upsert_key,omitempty"`
	NumberOfParallels   int64                                                        `json:"number_of_parallels"`
	Associations        *parameter.NullableObjectList[HubspotOutputAssociationInput] `json:"associations,omitempty"`
}

type UpdateHubspotOutputOptionInput struct {
	HubspotConnectionID *int64                                                       `json:"hubspot_connection_id,omitempty"`
	ObjectType          *string                                                      `json:"object_type,omitempty"`
	Mode                *string                                                      `json:"mode,omitempty"`
	UpsertKey           *parameter.NullableString                                    `json:"upsert_key,omitempty"`
	NumberOfParallels   *int64                                                       `json:"number_of_parallels,omitempty"`
	Associations        *parameter.NullableObjectList[HubspotOutputAssociationInput] `json:"associations,omitempty"`
}

type HubspotOutputAssociationInput struct {
	ToObjectType  string `json:"to_object_type"`
	FromObjectKey string `json:"from_object_key"`
	ToObjectKey   string `json:"to_object_key"`
}
