FROM golang:alpine
LABEL MAINTAINER=lwpimis@foxmail.com
WORKDIR $GOPATH/src/go_game
COPY . $GOPATH/src/go_game
EXPOSE 12001
CMD go run ./src/boostrap.go