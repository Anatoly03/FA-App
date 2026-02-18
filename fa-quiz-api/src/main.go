package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	// _ "fa-quiz-api/migrations"
)

func main() {
	// Check if the program is in development mode.
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())
	isAutoMigrate := isGoRun || os.Getenv("AUTOMIGRATE") == "1"

	_ = godotenv.Load(".env")
	app := pocketbase.New()

	// Register the migration command.
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: isAutoMigrate,
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
