import React, { Component } from 'react';
import { withRouter } from 'react-router-dom';
import './index.css';

class StoryHeader extends Component {
    render() {
        const { story, history } = this.props;

        return (
            <div className="story story--header story--failure">
                  <div className="story__vertical-indicator"></div>
                  
                  <div onClick={() => {history.push('/')}} className="story__navigation">
                    <i className="material-icons">
                      keyboard_backspace
                    </i>
                    
                    All stories
                  </div>
                  
                  <div className="story__info-container">
                    <div className="story__info">
                      <div className="story__title">
                          <span className="story__id">
                              #{story.ID}
                          </span>
                      
                          {story.name} 
                      </div>

                      <div className="story__metadata">
                          <div className="story__last-updated">
                              Last updated 8 days ago
                          </div>

                          <ul className="story__owners">
                              <li className="story__owners-pic">
                                  <img src="/images/1.jpg"/>
                              </li>
                              <li className="story__owners-pic">
                                  <img src="/images/2.jpg"/>
                              </li>
                          </ul>
                      </div>

                      <div className="story__status">
                          <div className="label">
                            TC-Review
                          </div>

                          <div className="story__cycle-time">
                            <i className="material-icons">
                              access_time
                            </i>
                            
                            90hours
                          </div>
                      </div>
                  </div>


                  <div className="story__statistics">
                      <div className="story-statistic">
                          <div className="story-statistic__value">
                              45%
                          </div>

                          <div className="story-statistic__name">
                              Confidence
                          </div>
                      </div>

                      <div className="story-statistic">
                          <div className="story-statistic__value">
                              {story.Comments.length}
                          </div>

                          <div className="story-statistic__name">
                              Comments
                          </div>
                      </div>

                      <div className="story-statistic">
                          <div className="story-statistic__value">
                              {!story.pullRequests ? 0 : story.pullRequests.length}
                          </div>

                          <div className="story-statistic__name">
                              Pull requests
                          </div>
                      </div>

                      <div className="story-statistic">
                          <div className="story-statistic__value">
                              3
                          </div>

                          <div className="story-statistic__name">
                              Complexity
                          </div>
                      </div>
                  </div>
              </div>
            </div>
        );
    }
}

const DecoratedStoryHeader = withRouter(StoryHeader)

export default DecoratedStoryHeader;
