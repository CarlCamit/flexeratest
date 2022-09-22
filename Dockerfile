# syntax=docker/dockerfile:1
FROM golang:1.17-alpine

WORKDIR /app

# copy Go modules and dependencies to image and download them
COPY go.mod ./
RUN go mod download

COPY . ./

# run tests
RUN go test -v ./...

# compile application
RUN go build -o /appcalc

# command to be used to execute when the image is used to start a container
CMD [ "/appcalc" ]