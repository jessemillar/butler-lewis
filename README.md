## Setup
You need to set three environment variables:
```
TRELLO_KEY
TRELLO_TOKEN
DUNN_SECRET
```

## Running
```
docker build -t dunn .
docker run -d -p 9999:9999 -e DUNN_SECRET="blah" -e TRELLO_KEY="12345" -e TRELLO_TOKEN="67890" --restart=always --name dunn dunn
```

## Usage
```
HTTP PUT
localhost:9999/v1/dunn?name=Test test&secret=blah
```
