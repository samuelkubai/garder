import React, { Component } from 'react';
import { withRouter } from 'react-router-dom';
import './index.css';

class Story extends Component {
    render() {
        const { story, history } = this.props;

        return (
            <div  className="single-story">
                <div className="single-story__info">
                    <div className="single-story__indicator"></div>
                    <div onClick={() => { history.push(`/stories/${story.ID}`)}} className="single-story__title">
                        <span className="single-story__id">
                            #{story.ID}
                        </span>
                    
                        {story.name} 
                    </div>

                    <div className="single-story__metadata">
                        <div className="single-story__last-updated">
                            Last updated 8 days ago
                        </div>

                        <ul className="single-story__owners">
                            <li className="single-story__owners-pic">
                                <img src="images/1.jpg"/>
                            </li>
                            <li className="single-story__owners-pic">
                                <img src="images/2.jpg"/>
                            </li>
                        </ul>
                    </div>
                </div>


                <div className="single-story__statistics">
                    <div className="single-story-statistic">
                        <div className="single-story-statistic__value">
                            45%
                        </div>

                        <div className="single-story-statistic__name">
                            Confidence
                        </div>
                    </div>

                    <div className="single-story-statistic">
                        <div className="single-story-statistic__value">
                            {story.Comments.length}
                        </div>

                        <div className="single-story-statistic__name">
                            Comments
                        </div>
                    </div>

                    <div className="single-story-statistic">
                        <div className="single-story-statistic__value">
                            {!story.pullRequests ? 0 : story.pullRequests.length}
                        </div>

                        <div className="single-story-statistic__name">
                            Pull requests
                        </div>
                    </div>

                    <div className="single-story-statistic">
                        <div className="single-story-statistic__value">
                            3
                        </div>

                        <div className="single-story-statistic__name">
                            Complexity
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

const DecoratedStory = withRouter(Story)

export default DecoratedStory;
