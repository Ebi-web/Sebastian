package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"google.golang.org/api/option"
	sheets "google.golang.org/api/sheets/v4"
)

func main() {
	err := godotenv.Load("./env/.dev.env")
	spreadsheetID, botToken := os.Getenv("SPREADSHEET_ID"), os.Getenv("BOT_TOKEN")
	c := slack.New(botToken)

	credential := option.WithCredentialsFile("./secret.json")

	srv, err := sheets.NewService(context.TODO(), credential)
	if err != nil {
		log.Fatal(err) // zaprに置換だよ
	}

	readRange := "A:C"

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).MajorDimension("COLUMNS").Do()
	if err != nil {
		log.Fatalln(err)
	}
	if len(resp.Values) == 0 {
		log.Fatalln("data not found")
	}
fmt.Print(resp.Values)
	_, _, err = c.PostMessage("テストチャンネル", slack.MsgOptionText("Hello World", true))
	if err != nil {
		panic(err)
	}
}
