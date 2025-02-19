package vismanet

import (
	"time"
)

// parseTime is a helper function to parse time from Visma
func parseTime(s string) (time.Time, error) {
	//For some reason Visma uses different time formats, so we need to try multiple formats
	formats := []string{
		time.RFC3339,
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05-07:00",
	}

	//Try to parse the time with the different formats
	var err error
	for _, format := range formats {
		var t time.Time
		t, err = time.Parse(format, s)
		if err == nil {
			return t, nil
		}
	}

	//Return the last error if we could not parse the time
	return time.Time{}, err
}
