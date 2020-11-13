package models

import (
	"api/internal/config"
)

type Measure struct {
	MeasureID int64  `json:"measureId" gorm:"primaryKey" gorm:"autoIncrement"`
	TypeID    int    `json:"typeId"`
	Value     string `json:"value"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Measure{})
}

func (m *Measure) CreateMeasure() *Measure {
	db.NewRecord(m)
	db.Create(&m)
	return m
}

func GetAllMeasureFromUsername(username string) *[]Measure {
	var measure []Measure
	db.Joins("JOIN value_types ON value_types.type_id = measures.type_id").Joins("JOIN users ON users.user_id = value_types.user_id").Where("users.username = ?", username).Find(&measure)
	return &measure
}
