package task

type FrequencyXO struct {
	Schedule       string        `json:"schedule"`
	StartDate      int           `json:"startDate,omitempty"`
	TimeZoneOffset string        `json:"timeZoneOffset,omitempty"`
	RecurringDays  []interface{} `json:"recurringDays,omitempty"`
	CronExpression string        `json:"cronExpression,omitempty"`
}

type Task struct {
	ID                    string                 `json:"id"`
	Type                  string                 `json:"type"`
	Name                  string                 `json:"name"`
	Message               string                 `json:"message,omitempty"`
	CurrentState          string                 `json:"currentState,omitempty"`
	LastRunResult         string                 `json:"lastRunResult,omitempty"`
	Enabled               bool                   `json:"enabled"`
	AlertEmail            string                 `json:"alertEmail,omitempty"`
	NotificationCondition string                 `json:"notificationCondition,omitempty"`
	Frequency             *FrequencyXO           `json:"frequency,omitempty"`
	NextRun               string                 `json:"nextRun,omitempty"`
	LastRun               string                 `json:"lastRun,omitempty"`
	Properties            map[string]interface{} `json:"properties,omitempty"`

	// GET /v1/tasks/{id} returns the schedule flat on the task, not nested
	// under "frequency" as on create (verified on Nexus 3.94.1). A task
	// created with schedule "cron" reads back as "advanced", and the server
	// fills in startDate (an ISO-8601 string) and timeZoneOffset itself.
	Schedule       string `json:"schedule,omitempty"`
	CronExpression string `json:"cronExpression,omitempty"`
	StartDate      string `json:"startDate,omitempty"`
	TimeZoneOffset string `json:"timeZoneOffset,omitempty"`
	RecurringDays  []int  `json:"recurringDays,omitempty"`
}

type TaskCreateStruct struct {
	Type                  string                 `json:"type,omitempty"`
	Name                  string                 `json:"name"`
	Enabled               bool                   `json:"enabled"`
	AlertEmail            string                 `json:"alertEmail,omitempty"`
	NotificationCondition string                 `json:"notificationCondition"`
	Frequency             *FrequencyXO           `json:"frequency,omitempty"`
	Message               string                 `json:"message,omitempty"`
	Properties            map[string]interface{} `json:"properties,omitempty"`
	ConcurrentRun         bool                   `json:"concurrentRun,omitempty"`
}
