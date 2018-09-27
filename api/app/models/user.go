package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

// The user information will be pulled from Pivotal tracker
// the image URL from github and the ID's from each
// platform will be saved but a new one will be
// created for eagle. We will try look for
// similarities in the names and if
// similar the name will be
// as from the same
// person.
type User struct {
    gorm.Model
    Name string `gorm:"type:varchar(100);not null" json:"name"`
    ImageURL string `gorm:"type:varchar(200);not null" json:"imageUrl"`
    Initials string `gorm:"type:varchar(5);not null" json:"initials"`
    PT_id int `gorm:"not null" json:"ptId"`
    GH_id int `gorm:"not null" json:"ghId"`
}

