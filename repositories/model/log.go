package model

type Log struct {
	LogEvent LogEvent `json:"after"`
}

type LogEvent struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	ResourceID   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
	Published    bool   `json:"published"`	
}

