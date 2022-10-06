package model

type Cat struct {
	ID           string  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name         string  `gorm:"column:name;size:100"`
	Color        string  `gorm:"column:color;size:100"`
	Weight       float64 `gorm:"column:weight"`
	Intelligence int     `gorm:"column:intelligence"`
	Laziness     int     `gorm:"column:laziness"`
	Curiosity    int     `gorm:"column:curiosity"`
	Sociability  int     `gorm:"column:sociability"`
	Egoism       int     `gorm:"column:egoism"`
	MiauPower    int     `gorm:"column:miau_power"`
	Attack       int     `gorm:"column:attack"`
}

func (Cat) TableName() string {
	return "cats"
}
