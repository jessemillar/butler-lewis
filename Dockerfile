FROM golang:1.12.7-alpine AS builder

# Download and install dep
ADD https://github.com/golang/dep/releases/download/v0.5.4/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
# Install Git so dep works
RUN apk add --no-cache git
# Install the Certificate-Authority certificates to enable HTTPS
RUN apk add --no-cache ca-certificates

# Set the environment variables for the go command:
# * CGO_ENABLED=0 to build a statically-linked executable
ENV CGO_ENABLED=0

WORKDIR $GOPATH/src/github.com/jessemillar/dunn
COPY ./ .
RUN dep ensure
RUN go build -installsuffix 'static' -o /app .

FROM scratch
# Import the Certificate-Authority certificates for enabling HTTPS
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app /app
CMD ["/app"]
