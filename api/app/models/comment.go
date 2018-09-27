package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

// Comments will be shared between Github (GH) and Pivotal
// tracker (PT) we will however be rendering formatted
// comments with all their appropriate HTML tags.
type Comment struct {
    gorm.Model
    Message string `gorm:"type:varchar(500);not null" json:"message"`
    State string `gorm:"type:varchar(50);not null;default:'commented';" json:"state"`
    PullRequestID int `json:"pullRequestId"`
    PullRequest PullRequest `gorm:"PRELOAD:false" json:"pullRequest"`
    StoryID int `json:"storyID"`
    Story Story `gorm:"PRELOAD:false" json:"story"`
    UserID int `json:"userId"`
    User User `gorm:"PRELOAD:false" json:"user"`
    Activity []Activity `gorm:"polymorphic:Actor;"`
}
