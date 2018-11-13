# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/app/yast

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/auth

RUN dep ensure

RUN go install 

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/auth

# Document that the service listens on port 8080.
EXPOSE 8080