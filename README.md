
[![CircleCI](https://circleci.com/gh/jessemillar/butler-lewis.svg?style=svg)](https://circleci.com/gh/jessemillar/butler-lewis)

## Overview
`butler-lewis` is a small REST API for interacting with Trello boards. There are two endpoints which are documented below. `butler-lewis` also serves as a personal playground for trying out new server, build, and deployment technologies.

## Setup
You need to set some environment variables:
```
PORT
TRELLO_KEY
TRELLO_TOKEN
BUTLER_LEWIS_SECRET
```

## Running
### Locally/Manually
```
docker build -t butler-lewis .
docker run -d -p 9999:9999 -e BUTLER_LEWIS_SECRET="blah" -e TRELLO_KEY="12345" -e TRELLO_TOKEN="67890" -e PORT="9999" --restart=always --name butler-lewis butler-lewis
```

### Heroku
#### Setup
```
heroku create
heroku stack:set container
```

#### Deploy
```
git push heroku master
```

## Usage
```
HTTP PUT
localhost:9999/v1/butler-lewis?name=Test test&secret=blah
```
