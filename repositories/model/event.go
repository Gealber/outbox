package model

type Event struct {
	ID           string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Type         string `gorm:"column:type"`
	ResourceID   string `gorm:"column:resource_id"`
	ResourceType string `gorm:"column:resource_type"`
	Published    bool   `gorm:"column:published"`
}

func (Event) TableName() string {
	return "events"
}
