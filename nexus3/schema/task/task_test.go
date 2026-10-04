package task

import (
	"encoding/json"
	"os"
	"testing"
)

// The fixtures are real GET /v1/tasks/{id} responses from Nexus 3.94.1.
func TestTaskUnmarshalFlatSchedule(t *testing.T) {
	for _, tc := range []struct {
		file, schedule, cron string
	}{
		{"testdata/get_task_cron.json", "advanced", "0 15 4 * * ?"},
		{"testdata/get_task_daily.json", "daily", ""},
	} {
		raw, err := os.ReadFile(tc.file)
		if err != nil {
			t.Fatal(err)
		}
		var got Task
		if err := json.Unmarshal(raw, &got); err != nil {
			t.Fatalf("%s: %v", tc.file, err)
		}
		if got.Schedule != tc.schedule || got.CronExpression != tc.cron {
			t.Errorf("%s: schedule=%q cron=%q, want %q %q", tc.file, got.Schedule, got.CronExpression, tc.schedule, tc.cron)
		}
		if got.Frequency != nil {
			t.Errorf("%s: GET has no nested frequency, got %+v", tc.file, got.Frequency)
		}
		if got.StartDate == "" {
			t.Errorf("%s: startDate should be the server-set ISO string", tc.file)
		}
		if got.Properties["blobstoreName"] != "(All Blob Stores)" {
			t.Errorf("%s: properties not decoded: %v", tc.file, got.Properties)
		}
	}
}
