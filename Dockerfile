FROM golang:1.16.14

WORKDIR /app
EXPOSE 8000

COPY . .

RUN go mod download && \
    go get -u github.com/slack-go/slack \
        github.com/joho/godotenv \
        google.golang.org/api/sheets/v4 \
        golang.org/x/oauth2/google

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]