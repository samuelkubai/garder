import React, { Component } from 'react';
import './index.css';

const icons = {
  "create-story": "thumb_up",
  "update-story": "edit",
  "story-status-started": "skip_next",
  "comment-story": "chat_bubble_outline",
  "opened-pull-request": "play_circle_outline",
  "synchronize-pull-request": "replay",
  "edited-pull-request": "edit",
  "labeled-pull-request": "label",
  "comment-pull-request": "chat_bubble_outline",
};

const title = {
  "create-story": (story) => `Story created`,
  "update-story": (story) => `Story updated`,
  "story-status-started": (story) => `Story started`,
  "comment-story": (comment) => `New comment on <a href="#">PT</a>`,
  "opened-pull-request":  (pull_request) => `Pull request opened`,
  "synchronize-pull-request":  (pr) => `Pull request synced`,
  "edited-pull-request": (pr) => `Pull request edited`,
  "labeled-pull-request": (pr) => `Labels updated`,
  "comment-pull-request": (comment) => `New comment on <a href="#">Github</a>`,
};

const body = {
  "create-story": (story) => `The story ${story.name} (#${story.ID}) has been created by John Doe`,
  "update-story": (story) => `#${story.ID} has been updated on pivotal tracker`,
  "story-status-started": (story) => `#${story.ID} has been started on pivotal tracker by John Doe`,
  "comment-story": (comment) => `${comment.message}`,
  "opened-pull-request":  (pr) => `#${pr.storyID} has a new pull request titled: <b>${pr.title}</b>`,
  "synchronize-pull-request":  (pr) => `${pr.title} has been synced from Github`,
  "edited-pull-request": (pr) => `${pr.title} has been edited on Github`,
  "labeled-pull-request": (pr) => `${pr.title} label's have been updated on Github`,
  "comment-pull-request": (pr) => `${pr.message}`,
};

class Activity extends Component {
  state = {
    actor: {}
  };

  componentDidMount() {
    const { activity } = this.props;

    fetch(`${process.env.REACT_APP_API_URL}/${activity.ActorType}?id=${activity.ActorID}`)
      .then(response => response.json())
      .then(data => {
        this.setState({
          isLoaded: true,
          actor: data  
        })
      });
  }

  renderIcon() {
    const { activity } = this.props;

    if (activity.type in icons) {
      return icons[activity.type];
    }

    return "thumbs-up";
  }

  renderTitle() {
    const { actor } = this.state;
    const { activity } = this.props;

    if (activity.type in title) {
      return title[activity.type](actor);
    }

    return "Loading title...";
  }

  renderBody() {
    const { actor } = this.state;
    const { activity } = this.props;

    if (activity.type in body) {
      return body[activity.type](actor);
    }

    return "Loading body...";
  }

  render() {
    return (
      <div className="activity">
        <div className="activity__icon">
          <i className="material-icons">
            {this.renderIcon()}
          </i>
        </div>

        <div className="activity__container">
          <div className="activity__title" dangerouslySetInnerHTML={{__html: this.renderTitle()}}></div>
          <div className="activity__body" dangerouslySetInnerHTML={{__html: this.renderBody()}}></div>
        </div>
      </div>
    );
  }
}

export default Activity;
