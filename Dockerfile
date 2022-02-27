FROM golang:1.18rc1-alpine3.15

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go get -u github.com/slack-go/slack

COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]