// API Testing for Guild Roles
// https://pocketbase.io/docs/go-testing/

package api_test

import (
	// "net/http"
	// "strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

// TestPermissionAdmin tests that only admins should be able to perform the
// following actions:
// - create new roles
// - delete existing roles
// - assign roles to users
func TestPermissionAdmin(t *testing.T) {
	t.Parallel()

	// superUserToken := getSuperUserToken(setupTestApp(t))

	scenarios := []tests.ApiScenario{
		// {
		// 	Name:   "create guild to populate test data",
		// 	Method: http.MethodPost,
		// 	URL:    "/api/collections/guilds/records",
		// 	Headers: map[string]string{
		// 		"Authorization": superUserToken,
		// 	},
		// 	Body: strings.NewReader(`{
		// 		"name": "Test Guild"
		// 	}`),
		// 	DisableTestAppCleanup: true,
		// 	ExpectedStatus: 200,
		// 	ExpectedContent: []string{
		// 		`"name":"Test Guild"`,
		// 	},
		// },
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
