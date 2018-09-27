package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

// A status is the current/previous state of a story
// the status can be derived from pivotal tracker
// or computed by eagle.
type Status struct {
    gorm.Model
    Name string `gorm:"type:varchar(100);not null;" json:"name"`
    Slug string `gorm:"type:varchar(100);not null;" json:"slug"`
    Stories []*Story `gorm:"many2many:story_statuses;" json:"stories,omitempty"`
}

