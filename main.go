/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dstotijn/go-notion"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	token := flag.String("token", "", "Notion integration token")
	databaseID := flag.String("database-id", "", "Notion database ID")
	year := flag.Int("year", 0, "Year(yyyy)")

	flag.Parse()

	if *token == "" || *databaseID == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *year < 1000 || *year > 9999 {
		flag.Usage()
		os.Exit(1)
	}

	file, err := os.Open("setting.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("Error decoding JSON:", err)
		os.Exit(1)
	}

	if err := NotionDBSyukujitsuImporter(*year, *token, *databaseID, config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("success![year:%d]\n", *year)
	os.Exit(0)
}

type Config struct {
	SourceSyukujitsuURL string `json:"sourceSyukujitsuURL"`
	NotionProperties    struct {
		Title string `json:"title"`
		Date  string `json:"date"`
	} `json:"notionProperties"`
}

func NotionDBSyukujitsuImporter(year int, token string, NotionDatabaseID string, config Config) error {
	holidays, err := GetHolidays(config.SourceSyukujitsuURL)
	if err != nil {
		return err
	}
	createHolidays := make([]Holiday, 0)
	for _, holiday := range holidays {
		if holiday.Date.Year() == year {
			createHolidays = append(createHolidays, holiday)
		}
	}
	for _, holiday := range createHolidays {
		holidayName := holiday.Name
		if holidayName == "休日" {
			holidayName = "振替休日"
		}
		if err := CreateNotionDatabasePage(token, NotionDatabaseID, notion.DatabasePageProperties{
			config.NotionProperties.Title: notion.DatabasePageProperty{
				Title: []notion.RichText{
					{Text: &notion.Text{Content: holidayName}},
				},
			},
			config.NotionProperties.Date: notion.DatabasePageProperty{
				Date: &notion.Date{
					Start: notion.NewDateTime(holiday.Date, false),
				},
			},
		}); err != nil {
			return err
		}
	}
	return nil
}

type Holiday struct {
	Date time.Time
	Name string
}

func GetHolidays(sourceURL string) ([]Holiday, error) {
	// HTTP GETリクエストを送信してCSVデータを取得
	response, err := http.Get(sourceURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch CSV data: %v", err)
	}
	defer response.Body.Close()

	// HTTPレスポンスのステータスコードを確認
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch CSV data. Status: %s", response.Status)
	}

	// Shift-JISエンコーディングをUTF-8に変換してCSVデータを取得
	transformer := japanese.ShiftJIS.NewDecoder()
	reader := transform.NewReader(response.Body, transformer)

	// CSVデータをパースして構造体に変換
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1 // 不定なフィールド数を許可する
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to parse CSV data: %v", err)
	}

	// CSVデータを構造体にマッピング
	var holidays []Holiday
	for _, record := range records {
		if len(record) >= 2 {
			date, err := time.Parse("2006/1/2", record[0])
			if err != nil {
				continue
			}
			holidays = append(holidays, Holiday{
				Date: date,
				Name: record[1],
			})
		}
	}

	return holidays, nil
}

func CreateNotionDatabasePage(token string, parentID string, properties notion.DatabasePageProperties) error {
	client := notion.NewClient(token)
	if _, err := client.CreatePage(context.Background(), notion.CreatePageParams{
		ParentType:             notion.ParentTypeDatabase,
		ParentID:               parentID,
		DatabasePageProperties: &properties,
	}); err != nil {
		return err
	}
	return nil
}
