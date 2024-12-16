package workflow

type TaskDependency struct {
	Source      string `json:"source,omitempty"`
	Destination string `json:"destination,omitempty"`
}
