package main

import (
    "fmt"
    "net/http"
    "os"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "app/api"
    "app/webhooks"
    "app/models"
)


func main() {
    PORT := fmt.Sprintf(":%s", os.Getenv("PORT")) // ":5000"

    db, err := initializeDatabase()
    if err != nil {
        fmt.Println("Database initialization failed")
        panic(err)
    }

    mux := http.NewServeMux()
    webhooks.RegisterWebhooks(mux, db)
    api.RegisterApiRoutes(mux, db)
    fmt.Println(fmt.Sprintf("Serving the application on http://localhost%s", PORT))
    http.ListenAndServe(PORT, mux)
}

func initializeDatabase() (db *gorm.DB, err error) {
    db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL") + "?sslmode=disable")
    if err == nil {
        db.LogMode(true)
        // Commit the migrations
        db.AutoMigrate(&models.Project{}, &models.Story{}, &models.Status{}, &models.User{}, &models.Comment{}, &models.PullRequest{}, &models.Label{}, &models.Activity{})
    }

    return
}
