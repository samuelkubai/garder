package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

// A story belongs to a single project but can have multiple
// owners assigned to the story to work on but our
// expectation is that one story will have a
// single owner.
type Story struct {
    gorm.Model
    Name string `gorm:"type:varchar(100);not null;" json:"name"`
    Type string `gorm:"type:varchar(20);not null;" json:"type"`
    Kind string `gorm:"type:varchar(20);not null;" json:"kind"`
    ProjectID int `json:"projectId"`
    Project Project `json:"project"`
    Activities []Activity `json:"activities,omitempty"`
    PullRequests []PullRequest `json:"pullRequests,omitempty"`
    Statuses []*Status `gorm:"many2many:story_statuses;" json:"statuses,omitempty"`
    Comments []Comment `json"comments,omitempty"`
    Activity []Activity `gorm:"polymorphic:Actor;"`
}

