package models

import "time"

type User struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Married   bool      `json:"married"`
	Password  string    `json:"-"`
	Login     string    `gorm:"index:idx_login,unique"                          json:"login"`
	Birthday  time.Time `json:"birthday"`
}
