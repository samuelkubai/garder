package api

import (
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "app/models"
    "app/internal"
)

type ActivitiesController struct {
    DB *gorm.DB
}

func (ctrl ActivitiesController) GetActivities(w http.ResponseWriter, r *http.Request) {
   enableCors(&w)

   var respond internal.Response
   var activities []models.Activity
  
   // Get the story ID and get the activities for that story.
   params := r.URL.Query()
   if storyID, ok := params["story"]; ok {
     ctrl.DB.Order("updated_at asc").Where("story_id = ?", storyID).Find(&activities)
   } else {
     ctrl.DB.Order("updated_at asc").Find(&activities)
   }

   respond.AsJson(w, activities)
}
