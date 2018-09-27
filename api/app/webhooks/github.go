package webhooks

import (
    "strconv"
    "regexp"
    "strings"
    "encoding/json"
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "app/models"
    "app/internal"
)

type GHActivity struct {
    Action string
    Number int
    Review GHReview
    Pull_request GHPullRequest
    Sender GHUser
}

type GHPullRequest struct {
    Url string
    Id int
    Title string
    User GHUser
    Body string
    State string
    Merged bool
    Mergeable bool
    Labels []GHLabel
}

type GHReview struct {
    Id int
    Body string
    State string
    User GHUser
}

type GHUser struct {
    Id int
    Avatar_url string
}

type GHLabel struct {
    Id int
    Name string
    Color string
}

type Github struct {
    DB *gorm.DB
}

func (gh *Github) ExtractStoryID(activity GHActivity) int {
    // Attempt to match the pull request to the appropriate
    // Pivotal tracker story for better tracking and
    // analytics.
    var storyID int

    re := regexp.MustCompile("#(\\d*)")
    if possibleIDs := re.FindAllString(activity.Pull_request.Title, -1); len(possibleIDs) > 0 {
        storyID, _ = strconv.Atoi(strings.Replace(possibleIDs[0], "#", "", -1))
    }
    
    return storyID
}

func (gh *Github) Comment(activity GHActivity) *models.Comment {
    storyID := gh.ExtractStoryID(activity)

    comment := &models.Comment{
        Message: activity.Review.Body,
        State: activity.Review.State,
        PullRequestID: activity.Pull_request.Id,
        StoryID: storyID,
    }
    
    gh.DB.Save(comment)

    LogCommentActivity(comment, "comment-pull-request", storyID, gh.DB)
    return comment
}

func (gh *Github) Save(activity GHActivity) *models.PullRequest {
    // Prepare the pull request labels
    var labels []*models.Label

    for _, label := range activity.Pull_request.Labels {
        labels = append(labels, &models.Label{
            Model: gorm.Model{ID: uint(label.Id)},
            Name: label.Name,
            Color: label.Color,
        })
    }

    storyID := gh.ExtractStoryID(activity)

    // Create or update a pull request based
    // on whether the pull request ID 
    // exists or not already.
    pull_request := &models.PullRequest{
        Model: gorm.Model{ID: uint(activity.Pull_request.Id)},
        Title: activity.Pull_request.Title,
        Body: activity.Pull_request.Body,
        State: activity.Pull_request.State,
        Merged: activity.Pull_request.Merged,
        Mergeable: activity.Pull_request.Mergeable,
        StoryID: storyID,
        Labels: labels,
    }
    
    gh.DB.Save(pull_request)

    LogPullRequestActivity(pull_request, activity.Action + "-pull-request", storyID, gh.DB)
    return pull_request
}

func (gh *Github) HandlePullRequestsActivity(w http.ResponseWriter, r *http.Request) {
    var respond internal.Response
    var response GHActivity


    err := json.NewDecoder(r.Body).Decode(&response)
    if err != nil {
        panic(err)
    }
    
    switch action := response.Action; action {
    case "submitted":
        respond.AsJson(w, gh.Comment(response))
    default:
        respond.AsJson(w, gh.Save(response))
    }
}
