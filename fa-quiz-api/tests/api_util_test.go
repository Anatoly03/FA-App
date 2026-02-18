package api_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

// setupTestApp initializes a new TestApp instance.
// TODO: read what testing.TB | testing.B | testing.T difference is
func setupTestApp(t testing.TB) *tests.TestApp {
	_, filePath, _, _ := runtime.Caller(0)
	dataDir := filepath.Clean(filepath.Join(filepath.Dir(filePath), "..", "pb_data"))

	testApp, err := tests.NewTestApp(dataDir)

	if err != nil {
		t.Fatal(err)
	}

	return testApp
}

// getAuthToken is a helper function to retrieve an auth token for a given user in the specified collection.
func getAuthToken(t testing.TB, app *tests.TestApp, recordId string) string {
	record, err := app.FindRecordById("users", recordId)
	if err != nil || record == nil {
		app.Logger().Error("Failed to find auth record", "record_id", recordId, "error", err)
		t.Fatalf("Missing auth record %q", recordId)
	}

	token, err := record.NewAuthToken()
	if err != nil {
		app.Logger().Error("Failed to generate token for auth record", "record_id", recordId, "error", err)
		t.Fatalf("Failed to generate auth token for %q: %v", recordId, err)
	}

	return token
}

// getSuperUserToken is a helper function to retrieve an auth token for the superuser.
func getSuperUserToken(t testing.TB, app *tests.TestApp) string {
	superUserRecord, err := app.FindRecordById("users", "pb_record_users")
	if err != nil || superUserRecord == nil {
		app.Logger().Error("Failed to find superuser record", "error", err)
		t.Fatal("Missing superuser record")
	}

	token, err := superUserRecord.NewAuthToken()
	if err != nil {
		app.Logger().Error("Failed to generate token for superuser", "error", err)
		t.Fatalf("Failed to generate superuser token: %v", err)
	}

	return token
}
