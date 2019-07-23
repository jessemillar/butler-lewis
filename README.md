# dunn
## Overview
[![CircleCI](https://circleci.com/gh/jessemillar/dunn.svg?style=svg)](https://circleci.com/gh/jessemillar/dunn)
`dunn` is a small REST API for interacting with Trello boards. There are two endpoints which are documented below. `dunn` also serves as a personal playground for trying out new server, build, and deployment technologies.

## Setup
You need to set some environment variables:
```
PORT
TRELLO_KEY
TRELLO_TOKEN
DUNN_SECRET
```

## Running
### Locally/Manually
```
docker build -t dunn .
docker run -d -p 9999:9999 -e DUNN_SECRET="blah" -e TRELLO_KEY="12345" -e TRELLO_TOKEN="67890" -e PORT="9999" --restart=always --name dunn dunn
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
localhost:9999/v1/dunn?name=Test test&secret=blah
```
