package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Label struct {
    gorm.Model
    Name string `json:"name"`
    Color string `json:"color"`
    PullRequests []*PullRequest `gorm:"many2many:pull_request_labels" json:"pullRequests,omitempty"`
}
