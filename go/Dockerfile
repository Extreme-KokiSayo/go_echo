FROM golang:1.17.7-alpine3.15

ENV GO111MODULE=on

WORKDIR /src
COPY ./src /src

#RUN go mod init main \
#    && go get github.com/labstack/echo/v4 \
#    && go mod tidy \
#    && go build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

EXPOSE 8080

#CMD ["go", "run", "main.go"]