package models

import (
	"github.com/jinzhu/gorm"
	//"github.com/thalesmonteiro/measurementApi/internal/config"
	"api/internal/config"
)

var db *gorm.DB
//user
type User struct {
	UserID   int    `json:"user_id" gorm:"primaryKey" gorm:"autoIncrement"`
	Username string `json:"username"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUser() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("user_id = ?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(ID int64) User {
	var user User
	db.Where("user_id = ?", ID).Delete(user)
	return user
}

func GetUserByUsername(username string) User {
	var user User
	db.Where("username = ?", username).Find(&user)
	return user
}

func GetUsersHasMeasure() []User {
	var user []User
	db.Joins("JOIN value_types ON value_types.user_id = users.user_id").Joins("JOIN measures ON value_types.type_id = measures.type_id").Group("user_id").Find(&user)
	return user
}
