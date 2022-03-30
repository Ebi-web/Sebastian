package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets"
)

var spreadsheetID = "1ncpIaM8AqDLwcJOyHTdPR52SBpZKQQS3F5spX-wDqjc"

func main() {
	credential := option.WithCredentialsFile("./secret.json")

	srv, err := sheets.NewService(context.TODO(), credential)
	if err != nil {
		log.Fatal(err)
	}

	readRange := "A1:B3"

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalln(err)
	}
	if len(resp.Values) == 0 {
		log.Fatalln("data not found")
	}
	for _, row := range resp.Values {
		fmt.Printf("%s, %s\n", row[0], row[1])
	}
}
