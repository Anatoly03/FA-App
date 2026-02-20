package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "fa-quiz-api/migrations"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

var App *pocketbase.PocketBase

func getSyncDir() string {
	if syncDir := strings.TrimSpace(os.Getenv("PB_SYNC_DIR")); syncDir != "" {
		return syncDir
	}

	wdCandidate := filepath.Join(".", "pb_sync")
	if info, err := os.Stat(wdCandidate); err == nil && info.IsDir() {
		return wdCandidate
	}

	if executablePath, err := os.Executable(); err == nil {
		execCandidate := filepath.Join(filepath.Dir(executablePath), "pb_sync")
		if info, statErr := os.Stat(execCandidate); statErr == nil && info.IsDir() {
			return execCandidate
		}
	}

	return wdCandidate
}

// Sync all records to localstore and upload to git.
func UploadData(e *core.RecordEvent) error {
	if err := e.Next(); err != nil {
		return err
	}
	syncDir := getSyncDir()
	if mkdirErr := os.MkdirAll(syncDir, 0755); mkdirErr != nil {
		log.Printf("Error ensuring sync dir %q: %v", syncDir, mkdirErr)
		return mkdirErr
	}
	records, err := e.App.FindAllRecords(e.Record.Collection().Id)
	if err != nil {
		log.Printf("Error fetching records: %v", err)
		return err
	}

	var mappedRecords []map[string]any
	for _, record := range records {
		mappedRecord := record.FieldsData()
		delete(mappedRecord, "created")
		delete(mappedRecord, "updated")
		mappedRecords = append(mappedRecords, mappedRecord)
	}
	jsonData, err := json.MarshalIndent(mappedRecords, "", "  ")

	filePath := filepath.Join(syncDir, e.Record.Collection().Name+".json")
	if err = os.WriteFile(filePath, jsonData, 0644); err != nil {
		log.Printf("Error writing file: %v", err)
		return err
	}

	return nil
}

// Sync all records to localstore and upload to git.
func DownloadData(e *core.ServeEvent) error {
	syncDir := getSyncDir()
	log.Printf("sync source directory: %s", syncDir)

	files := []string{
		"lecture",
		"chapters",     // depends on lecture
		"mc_questions", // depends on chapters
	}

	for _, file := range files {
		filePath := filepath.Join(syncDir, file+".json")
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

		collection, err := e.App.FindCollectionByNameOrId(file)
		if err != nil {
			log.Printf("Error finding collection %q: %v", file, err)
			continue
		}

		for _, record := range records {
			recordModel := core.NewRecord(collection)
			if id, ok := record["id"].(string); ok {
				recordModel.Id = id
			}

			delete(record, "collectionId")
			delete(record, "collectionName")
			delete(record, "created")
			delete(record, "updated")

			recordModel.Load(record)
			if err = e.App.Save(recordModel); err != nil {
				log.Printf("Error saving record %q in %q: %v", recordModel.Id, collection.Name, err)
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
