FROM node:carbon

ENV TERM=xterm-256color

WORKDIR /usr/app

# In order to allow this command to use the cache if
# the package.json or yarn.lock file has not 
# changed
COPY client/package.json /usr/app
COPY client/yarn.lock /usr/app

RUN yarn install
