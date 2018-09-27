package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type PullRequest struct {
    gorm.Model
    Title string `gorm:"type:varchar(50);not null;" json:"title"`
    Body string `gorm:"type:varchar(500);not null;" json:"body"`
    State string `gorm:"type:varchar(20);not null;" json:"state"`
    Merged bool `json:"merged"`
    Mergeable bool `gorm:"not null;" json:"mergeable"`
    StoryID int `json:"storyId"`
    User User `json:"user"`
    Story Story `json:"story"`
    Labels []*Label `gorm:"many2many:pull_request_labels" json:"label"`
    Comments []Comment `json:"comments,omitempty"`
    Activity []Activity `gorm:"polymorphic:Actor;"`
}
