package jobs

import (
    
)

type StoriesJob struct {
    interval int
}

func (stories StoriesJob) Run() {
    // TODO: Fetch the stories from pivotal tracker
    // with a time frame of last updated 
    // {interval} seconds ago


    // Create the records that don't exist.
    // Update the records that exist.
}
