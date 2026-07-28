package entity

type DbtGitRepository struct {
	ID              int64   `json:"id"`
	Name            string  `json:"name"`
	Description     *string `json:"description"`
	AdapterType     string  `json:"adapter_type"`
	DbtVersion      string  `json:"dbt_version"`
	URL             string  `json:"url"`
	RefType         string  `json:"ref_type"`
	Branch          *string `json:"branch"`
	Tag             *string `json:"tag"`
	CommitHash      *string `json:"commit_hash"`
	Subdirectory    *string `json:"subdirectory"`
	ResourceGroupID *int64  `json:"resource_group_id"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}
