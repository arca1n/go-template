package models

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

// User model for a user
type User struct {
	gorm.Model
	Username string  `gorm:"type:varchar(100);unique_index"`
	Email    string  `gorm:"type:varchar(100);unique_index"`
	Password string  `gorm:"size:255"`
	Profile  Profile `gorm:"FOREIGNKEY"`
}

// Profile is a user's profile
type Profile struct {
	gorm.Model
	Name     string
	Age      sql.NullInt64
	Birthday *time.Time

	Role         string  `gorm:"size:255"`        // set field size to 255
	MemberNumber *string `gorm:"unique;not null"` // set member number to unique and not null
	Address      string  `gorm:"index:addr"`      // create index with name `addr` for address
}
