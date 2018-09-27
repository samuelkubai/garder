package api

import (
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "app/models"
    "app/internal"
)

type CommentsController struct {
    DB *gorm.DB
}

func (ctrl CommentsController) GetComments(w http.ResponseWriter, r *http.Request) {
   enableCors(&w)

   var respond internal.Response
   var comments []models.Comment
   var comment models.Comment

   params := r.URL.Query()
   if commentID, ok := params["id"]; ok {
      ctrl.DB.Preload("PullRequest").Preload("Story").Preload("Activity").Where("id = ?", commentID).First(&comment)
      respond.AsJson(w, comment)
   } else {
      ctrl.DB.Preload("PullRequest").Preload("Story").Preload("Activity").Find(&comments)
      respond.AsJson(w, comments)
   }
}
