package pipeline_definition

type Workflow struct {
	ID               int64             `json:"id"`
	Name             *string           `json:"name"`
	Description      *string           `json:"description"`
	Labels           []string          `json:"labels"`
	Notifications    []Notification    `json:"notifications"`
	Schedules        []Schedule        `json:"schedules"`
	Tasks            []*Task           `json:"tasks"`
	TaskDependencies []*TaskDependency `json:"task_dependencies"`
}
