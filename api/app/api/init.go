package api

import (
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func RegisterApiRoutes(mux *http.ServeMux, db *gorm.DB) {
    // Initialize controllers
    storiesCtrl := StoriesController{ DB: db }
    commentsCtrl := CommentsController{ DB: db }
    pullRequestsCtrl := PullRequestsController{ DB: db }
    activitiesCtrl := ActivitiesController{ DB: db }

    // Register the routes
    mux.HandleFunc("/comments", commentsCtrl.GetComments)
    mux.HandleFunc("/stories", storiesCtrl.GetStories)
    mux.HandleFunc("/pull_requests", pullRequestsCtrl.GetPullRequests)
    mux.HandleFunc("/activity", activitiesCtrl.GetActivities)
}

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}
