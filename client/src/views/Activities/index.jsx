import React, { Component, Fragment } from 'react';
import { withRouter } from 'react-router-dom';
import StoryHeader from '../../components/StoryHeader';
import Activity from '../../components/Activity';

class Activities extends Component {
  state = {
    error: null,
    isLoaded: false,
    activities: [],
    story: {
      ID: '',
      name: '',
      Comments: [],
      pull_request: [],
    }
  }
  
  componentDidMount() {
    const { match } = this.props;
    fetch(`${process.env.REACT_APP_API_URL}/stories?id=${match.params.storyID}`)
      .then(response => response.json())
      .then(data => {
        this.setState({
          isLoaded: true,
          story: data,
          activities: data.activities ? data.activities : [],
        })
      }) 
  }

  renderStoryHeader() {
    const { story } = this.state; 
    return (<StoryHeader story={story} />);
  }

  renderActivities() {
    const { activities } = this.state;

    return (
      <div className="activities__container">
        {activities.map(activity => <Activity activity={activity}></Activity>)}
      </div>
    ); 
  }

  render() {
    return (
      <Fragment>
       {this.renderStoryHeader()}
       {this.renderActivities()}
      </Fragment>
    );
  }
}

const DecoratedActivities = withRouter(Activities);

export default DecoratedActivities;
