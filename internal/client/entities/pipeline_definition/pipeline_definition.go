package pipeline_definition

type Workflow struct {
	ID                           int64             `json:"id"`
	ResourceGroupID              *int64            `json:"resource_group_id"`
	Name                         *string           `json:"name"`
	Description                  *string           `json:"description"`
	MaxTaskParallelism           *int64            `json:"max_task_parallelism"`
	ExecutionTimeout             *int64            `json:"execution_timeout"`
	MaxRetries                   *int64            `json:"max_retries"`
	MinRetryInterval             *int64            `json:"min_retry_interval"`
	IsConcurrentExecutionSkipped *bool             `json:"is_concurrent_execution_skipped"`
	IsStoppedOnErrors            *bool             `json:"is_stopped_on_errors"`
	Labels                       []string          `json:"labels"`
	Notifications                []Notification    `json:"notifications"`
	Schedules                    []Schedule        `json:"schedules"`
	Tasks                        []*Task           `json:"tasks"`
	TaskDependencies             []*TaskDependency `json:"task_dependencies"`
}
