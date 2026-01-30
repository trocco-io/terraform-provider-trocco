package output_option

type HubspotOutputOption struct {
	HubspotConnectionID int64                       `json:"hubspot_connection_id"`
	ObjectType          string                      `json:"object_type"`
	Mode                string                      `json:"mode"`
	UpsertKey           *string                     `json:"upsert_key"`
	NumberOfParallels   int64                       `json:"number_of_parallels"`
	Associations        *[]HubspotOutputAssociation `json:"associations"`
}

type HubspotOutputAssociation struct {
	ToObjectType  string `json:"to_object_type"`
	FromObjectKey string `json:"from_object_key"`
	ToObjectKey   string `json:"to_object_key"`
}
