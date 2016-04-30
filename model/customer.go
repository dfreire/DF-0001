package model

type Customer struct {
	ID    string `gorm:"primary_key"`
	Name  string
	Email string
}
