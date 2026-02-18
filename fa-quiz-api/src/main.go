package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	_ "fa-quiz-api/migrations"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

var App *pocketbase.PocketBase

// Sync all records to localstore and upload to git.
func UploadData(e *core.RecordEvent) error {
	if err := e.Next(); err != nil {
		return err
	}

	records, err := e.App.FindAllRecords(e.Record.Collection().Id)
	if err != nil {
		log.Printf("Error fetching records: %v", err)
		return err
	}

	jsonData, err := json.MarshalIndent(records, "", "  ")

	filePath := "pb_sync/" + e.Record.Collection().Name + ".json"
	if err = os.WriteFile(filePath, jsonData, 0644); err != nil {
		log.Printf("Error writing file: %v", err)
		return err
	}

	return nil
}

// Sync all records to localstore and upload to git.
func DownloadData(e *core.ServeEvent) error {
	files := []string{
		"lecture",
		"chapters", // depends on lecture
		"mc_questions", // depends on chapters
	}

	for _, file := range files {
		filePath := "pb_sync/" + file + ".json"
		jsonData, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("Error reading file: %v", err)
			continue
		}

		var records []map[string]any
		if err = json.Unmarshal(jsonData, &records); err != nil {
			log.Printf("Error unmarshalling JSON: %v", err)
			continue
		}

		if len(records) == 0 {
			continue
		}

		collectionName, _ := records[0]["collectionName"].(string)
		collectionId, _ := records[0]["collectionId"].(string)
		collectionKey := collectionName
		if collectionKey == "" {
			collectionKey = collectionId
		}

		collection, err := e.App.FindCollectionByNameOrId(collectionKey)
		if err != nil {
			log.Printf("Error finding collection %q: %v", collectionKey, err)
			continue
		}

		for _, record := range records {
			recordModel := core.NewRecord(collection)
			if id, ok := record["id"].(string); ok {
				recordModel.Id = id
			}

			delete(record, "collectionId")
			delete(record, "collectionName")

			recordModel.Load(record)
			if err = e.App.Save(recordModel); err != nil {
				log.Printf("Error saving record: %v", err)
				continue
			}
		}
	}

	return e.Next()
}

func main() {
	// Check if the program is in development mode.
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())
	isAutoMigrate := isGoRun || os.Getenv("AUTOMIGRATE") == "1"

	_ = godotenv.Load(".env")
	App = pocketbase.New()

	// Register the migration command.
	migratecmd.MustRegister(App, App.RootCmd, migratecmd.Config{
		Automigrate: isAutoMigrate,
	})

	App.OnServe().BindFunc(DownloadData)
	App.OnRecordCreate().BindFunc(UploadData)
	App.OnRecordUpdate().BindFunc(UploadData)
	App.OnRecordDelete().BindFunc(UploadData)

	if err := App.Start(); err != nil {
		log.Fatal(err)
	}
}
