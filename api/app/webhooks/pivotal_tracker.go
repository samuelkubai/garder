package webhooks

import (
    "encoding/json"
    "fmt"
    "strings"
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "app/models"
    "app/internal"
)

type PTActivity struct {
    Kind string
    Guid string
    Project_version int
    Message string
    Highlight string
    Changes []PTChange
    Primary_resources []PTStory
    Project PTProject
    Performed_by PTActor
}

type PTProject struct {
    Kind string
    Id int
    Name string
}

type PTStory struct {
    Kind string
    Id int
    Name string
    Story_type string
    URL string
}
type PTChange struct {
    Kind string
    Change_type string
    ID int
    New_values PTStoryValues
}

type PTStoryValues struct {
    Id int
    Name string
    Description string
    Story_type string
    Current_state string
    Owner_ids []int
    Label_ids []int
    Created_at int
    Updated_at int
    Labels []string

    // Fields for comments
    Text string
}

type PTActor struct {
    Kind string
    Id int
    Name string
    Initials string
}

type PivotalTracker struct {
    DB *gorm.DB
}

func (pt *PivotalTracker) Save(activity PTActivity) *models.Story {
    // All the story's associations are created or fetched or updated 
    // implicitly during the saving of the story. We are using
    // save to be able to reuse the same function for both
    // create actions and update.
    story := &models.Story{
        Model: gorm.Model{ID: uint(activity.Primary_resources[0].Id)},
        Name: activity.Primary_resources[0].Name,
        Kind: activity.Primary_resources[0].Kind,
        Type: activity.Primary_resources[0].Story_type,
        Project: models.Project{
            Model: gorm.Model{ID: uint(activity.Project.Id)},
            Name: activity.Project.Name,
        },
        Statuses: []*models.Status{
            &models.Status{
                Name: activity.Changes[0].New_values.Current_state,
                Slug: strings.Replace(strings.ToLower(activity.Changes[0].New_values.Current_state), " ", "-", -1),
            },
        },
    }

    pt.DB.Save(story)
    
    return story
}

func (pt *PivotalTracker) create(activity PTActivity) *models.Story {
    story := pt.Save(activity)
    LogStoryActivity(story, "create-story", pt.DB)

    return story
}

func (pt *PivotalTracker) update(activity PTActivity) *models.Story {
  story := pt.Save(activity)
  LogStoryActivity(story, "update-story", pt.DB)

  return story
}

func (pt *PivotalTracker) comment(activity PTActivity) *models.Comment {
    // Fetch the pull request.
    var pull_request models.PullRequest
    pt.DB.Where("story_id = ?", activity.Primary_resources[0].Id).First(&pull_request)

    comment := &models.Comment{
        Message: activity.Changes[0].New_values.Text,
        State: "commented",
        PullRequestID: int(pull_request.ID),
        StoryID: activity.Primary_resources[0].Id,
    }
    
    pt.DB.Save(comment)
    LogCommentActivity(comment, "comment-story", activity.Primary_resources[0].Id, pt.DB) 

    return comment
}

func (pt *PivotalTracker) markAs(status string, activity PTActivity) *models.Story {
    // All the story's associations are created or fetched or updated 
    // implicitly during the saving of the story. We are using
    // save to reuse the same function for both creation
    // and updating.
    story := &models.Story{
        Model: gorm.Model{ID: uint(activity.Primary_resources[0].Id)},
        Name: activity.Primary_resources[0].Name,
        Kind: activity.Primary_resources[0].Kind,
        Type: activity.Primary_resources[0].Story_type,
        Project: models.Project{
            Model: gorm.Model{ID: uint(activity.Project.Id)},
            Name: activity.Project.Name,
        },
        Statuses: []*models.Status{
            &models.Status{
                Name: status,
                Slug: strings.Replace(strings.ToLower(status), " ", "-", -1),
            },
        },
    }

    pt.DB.Save(story)
    LogStoryActivity(story, "story-status-" + strings.ToLower(status), pt.DB)

    return story
}

// Recieves all the activity from PT and sorts
// it and performs the various actions 
// required.
func (pt *PivotalTracker) HandleActivity(w http.ResponseWriter, r *http.Request) {
    var respond internal.Response
    var response PTActivity

    err := json.NewDecoder(r.Body).Decode(&response)
    if err != nil {
        panic(err)
    }
    
    switch kind := response.Highlight; kind {
    case "added":
        respond.AsJson(w, pt.create(response))
    case "edited":
        respond.AsJson(w, pt.update(response))
    case "started":
        respond.AsJson(w, pt.markAs(kind, response))
    case "finished":
        respond.AsJson(w, pt.markAs(kind, response))
    case "delivered":
        respond.AsJson(w, pt.markAs(kind, response))
    case "accepted":
        respond.AsJson(w, pt.markAs(kind, response))
    case "rejected":
        respond.AsJson(w, pt.markAs(kind, response))
    case "added comment:":
        respond.AsJson(w, pt.comment(response))
    default:
        fmt.Fprintf(w, "Haven't handled the activity: %s", kind)
    }
}
