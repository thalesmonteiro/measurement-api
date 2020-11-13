package models

import (
	"api/internal/config"
)

type ValueTypes struct {
	TypeID      int    `json:"typeId" gorm:"primaryKey" gorm:"autoIncrement"`
	UserID      int    `json:"userId"`
	Description string `json:"description"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&ValueTypes{})
}

func (v *ValueTypes) CreateValueType() *ValueTypes {
	db.NewRecord(v)
	db.Create(&v)
	return v
}

func GetAllTypes() []ValueTypes {
	var types []ValueTypes
	db.Find(&types)
	return types
}

func GetTypeByIdAndDescription(userID int, description string) *ValueTypes {
	var valueType ValueTypes
	db.Where("user_id = ? and description = ?", userID, description).Find(&valueType)
	return &valueType
}

func GetTypeByID(typeID int) ValueTypes {
	var valueType ValueTypes
	db.Where("type_id = ?", typeID).Find(&valueType)

	return valueType
}

func GetTypesByUser(username string) []ValueTypes {
	var types []ValueTypes
	db.Joins("JOIN users ON value_types.user_id = users.user_id AND users.username = ?", username).Find(&types)

	return types
}

func GetTypeForAllUsersByDescription(description string) []ValueTypes {
	var types []ValueTypes
	db.Joins("JOIN users ON value_types.user_id = users.user_id AND value_types.description = ?", description).Find(&types)

	return types
}
