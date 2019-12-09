#Adding base image
FROM golang:alpine

# Setting my GOPATH and GOROOT from system environment
ARG app_env
ENV APP_ENV $app_env

COPY ./ /go/src/github.com/login-microservice
WORKDIR /go/src/github.com/login-microservice
RUN go get ./
RUN go build -o /app

 
CMD ["/app"]
