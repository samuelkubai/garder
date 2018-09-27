import React from 'react';
import { Route, Switch } from 'react-router-dom';
import Stories from '../views/Stories';
import Activities from '../views/Activities';

const Routes = () => (
  <Switch>
    <Route
      path="/"
      exact
      component={Stories} />
    <Route
      path="/stories/:storyID"
      exact
      component={Activities} />
  </Switch>
)

export default Routes;
