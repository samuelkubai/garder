package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Activity struct {
    gorm.Model
    Type string `json:"type"`
    Story Story
    StoryID int
    ActorID int
    ActorType string
}
