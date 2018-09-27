package api

import (
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "app/models"
    "app/internal"
)

type StoriesController struct {
    DB *gorm.DB
}

func (ctrl StoriesController) GetStories(w http.ResponseWriter, r *http.Request) {
   enableCors(&w)

   var respond internal.Response
   var stories []models.Story
   var story models.Story

   params := r.URL.Query()
   if storyID, ok := params["id"]; ok {
      ctrl.DB.Preload("Activities", func(db *gorm.DB) *gorm.DB {
        return db.Order("activities.updated_at DESC")
      }).Preload("Activity").Preload("Project").Preload("PullRequests").Preload("Comments").Where("id = ?", storyID).First(&story)
      respond.AsJson(w, story)
   } else {
      ctrl.DB.Preload("Project").Preload("PullRequests").Preload("Comments").Find(&stories)
      respond.AsJson(w, stories)
   }
}
