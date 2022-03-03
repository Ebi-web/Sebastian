FROM golang:1.16.14

WORKDIR /app

COPY . .

RUN go mod download && \
    go get -u github.com/slack-go/slack \
        github.com/go-sql-driver/mysql


RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]