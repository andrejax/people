FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go get -u github.com/go-swagger/go-swagger/cmd/swagger

WORKDIR /source
COPY ./people .
RUN swagger generate spec -o swagger.json
WORKDIR swagger
RUN swagger generate server --spec=../swagger.json

COPY ./swagger/go.mod .
RUN go mod download

WORKDIR cmd/awesome-server

RUN go build -o main .

EXPOSE 2020

# Command to run when starting the container
CMD ["/source/swagger/cmd/awesome-server/main", "--port=2020", "--host=0.0.0.0"]