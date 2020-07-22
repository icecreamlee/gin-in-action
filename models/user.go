package models

import "gin_in_action/core"

type User struct {
	ID        uint        `gorm:"primary_key" json:"id"`
	Name      string      `json:"name"`
	Password  string      `json:"-"`
	Passsalt  string      `json:"-"`
	IsDelete  int         `json:"is_delete"`
	CreatedAt core.MTime  `json:"created_at"`
	UpdatedAt core.MTime  `json:"updated_at"`
	DeletedAt *core.MTime `json:"-"`
}

func (u User) TableName() string {
	return "users"
}

func GetUser(id uint) User {
	var user User
	DB.Where("id = ? ", id).First(&user)
	return user
}

func GetUserByName(name string) User {
	var user User
	DB.Where("name = ? ", name).First(&user)
	return user
}
