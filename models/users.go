package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name     string `gorm:"size:100" json: "nama" form: "nama"`
	Email    string `json: "email" form: "email"`
	Password string `json: "password" form: "password"`
}
