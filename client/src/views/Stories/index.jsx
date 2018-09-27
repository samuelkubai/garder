import React, { Component, Fragment } from 'react';
import NavBar from '../../components/NavBar';
import Story from '../../components/Story';
import './index.css';

class Stories extends Component {
  state = {
    error: null,
    isLoaded: false,
    stories: [],
  }
  
  componentDidMount() {
    fetch(process.env.REACT_APP_API_URL + "/stories")
      .then(response => response.json())
      .then(data => {
        this.setState({
          isLoaded: true,
          stories: data  
        })
      }) 
  }
  
  renderStories () {
    const { stories } = this.state;
    console.log(stories.map(story => story));

    return (
        <Fragment>
            {stories.map(story => <Story story={story}></Story>)}
        </Fragment>
    );
  }

  render() {
    return (
        <Fragment>
            <NavBar></NavBar>
            
            <div className="stories">
                {this.renderStories()}
            </div>
        </Fragment>
    );
  }
}

export default Stories;
