package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

// Projects will be pulled from Pivotal tracker (PT)
// and all the stories will be linked to a single
// project. This will allow for automatic
// discovery as soon as you add eagle
// as a webhook on your PT project.
type Project struct {
    gorm.Model
    Name string `gorm:"type:varchar(100);not null" json:"name"`
    Stories []Story `json:"stories,omitempty"`
}
