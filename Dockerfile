FROM golang:1.6.2-alpine

RUN apk update && apk upgrade && apk add git

RUN mkdir -p /go/src/github.com/jessemillar
ADD . /go/src/github.com/jessemillar/dunn

WORKDIR /go/src/github.com/jessemillar/dunn
RUN go get -d -v
RUN go install -v

CMD ["/go/bin/dunn"]

EXPOSE 9999
