package api

import (
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "app/models"
    "app/internal"
)

type PullRequestsController struct {
    DB *gorm.DB
}

func (ctrl PullRequestsController) GetPullRequests(w http.ResponseWriter, r *http.Request) {
   enableCors(&w)

   var respond internal.Response
   var pullRequests []models.PullRequest
   var pullRequest models.PullRequest

   params := r.URL.Query()
   if pullRequestID, ok := params["id"]; ok {
      ctrl.DB.Preload("Activity").Where("id = ?", pullRequestID).First(&pullRequest)
      respond.AsJson(w, pullRequest)
   } else {
      ctrl.DB.Preload("Activity").Find(&pullRequests)
      respond.AsJson(w, pullRequests)
   }
}
