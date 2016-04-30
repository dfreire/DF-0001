package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type CustomerRole struct {
	ID string `gorm:"primary_key"`
}

type Customer struct {
	ID                   string `gorm:"primary_key"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Name                 string
	Email                string       `gorm:"not null;unique"`
	Role                 CustomerRole `gorm:"ForeignKey:RoleId"`
	RoleId               string       `sql:"type:varchar(255) NOT NULL REFERENCES customer_role(id)"`
	SignedUpToNewsletter bool
	InNewsletter         bool
}

type WineComment struct {
	ID         string `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Customer   Customer `gorm:"ForeignKey:CustomerId"`
	CustomerId string   `sql:"type:varchar(255) NOT NULL REFERENCES customer(id),index:idx_wine_comment"`
	WineId     string   `gorm:"not null;unique_index:idx_wine_comment"`
	WineYear   int      `gorm:"not null;unique_index:idx_wine_comment"`
	Comment    string   `gorm:"not null;size:5000"`
}

func Initialize(db *gorm.DB) {
	db.AutoMigrate(CustomerRole{})
	db.AutoMigrate(Customer{})
	db.AutoMigrate(WineComment{})

	for _, id := range getCustomerRoleIds() {
		role := CustomerRole{ID: id}
		db.Where(role).FirstOrCreate(&role)
	}
}

func getCustomerRoleIds() []string {
	return []string{
		"sommelier",
		"restaurant",
		"wine_distribution",
		"wine_shop",
		"wine_lover",
		"other",
	}
}
