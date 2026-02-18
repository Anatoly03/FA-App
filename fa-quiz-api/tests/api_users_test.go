// API Testing for Users Collection
//
// https://pocketbase.io/docs/go-testing/
// https://github.com/presentator/presentator/blob/master/api_users_test.go

package api_test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

// TestUsersCreate ensures everyone can create a new unique user record
func TestUsersCreate(t *testing.T) {
	t.Parallel()

	scenarios := []tests.ApiScenario{
		{
			Name:   "everyone should be able to create a user",
			Method: http.MethodPost,
			URL:    "/api/collections/users/records",
			Body: strings.NewReader(`{
				"name":            "test",
				"email":           "test@example.com",
				"password":        "1234567890",
				"passwordConfirm": "1234567890"
			}`),
			TestAppFactory: setupTestApp,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"name":"test"`,
				`"verified":false`,
			},
			ExpectedEvents: map[string]int{
				"*":                          0,
				"OnRecordCreateRequest":      1,
				"OnModelValidate":            1,
				"OnModelCreate":              1,
				"OnModelCreateExecute":       1,
				"OnModelAfterCreateSuccess":  1,
				"OnRecordEnrich":             1,
				"OnRecordValidate":           1,
				"OnRecordCreate":             1,
				"OnRecordCreateExecute":      1,
				"OnRecordAfterCreateSuccess": 1,
			},
		},
		{
			Name:   "creating user with existing email should fail",
			Method: http.MethodPost,
			URL:    "/api/collections/users/records",
			Body: strings.NewReader(`{
				"name":            "test_new",
				"email":           "test@example.com",
				"password":        "1234567890",
				"passwordConfirm": "1234567890"
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"message":"Value must be unique."`,
			},
			ExpectedEvents: map[string]int{
				"*":                        0,
				"OnModelAfterCreateError":  1,
				"OnModelCreate":            1,
				"OnModelCreateExecute":     1,
				"OnModelValidate":          1,
				"OnRecordAfterCreateError": 1,
				"OnRecordCreate":           1,
				"OnRecordCreateExecute":    1,
				"OnRecordCreateRequest":    1,
				"OnRecordValidate":         1,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

// Ensure compliance with GDPR right to erasure: https://gdpr-info.eu/art-17-gdpr/
// https://eur-lex.europa.eu/eli/reg/2023/2854/oj/eng
func TestUsersDelete(t *testing.T) {
	t.Parallel()

	userId := "null"
	// authToken := "null"

	scenarios := []tests.ApiScenario{
		{
			Name:                  "guests trying to delete non-existant users should not break",
			Method:                http.MethodDelete,
			URL:                   "/api/collections/users/records/pb_record_users",
			TestAppFactory:        setupTestApp,
			DisableTestAppCleanup: true,
			ExpectedStatus:        404,
			ExpectedContent:       []string{`"data":{}`},
			ExpectedEvents:        map[string]int{"*": 0},
		},
		{
			Method: http.MethodPost,
			URL:    "/api/collections/users/records",
			Body: strings.NewReader(`{
				"name":            "test",
				"email":           "test@example.com",
				"password":        "1234567890",
				"passwordConfirm": "1234567890"
			}`),
			TestAppFactory:        setupTestApp,
			DisableTestAppCleanup: true,
			AfterTestFunc: func(t testing.TB, app *tests.TestApp, res *http.Response) {
				var responseData struct {
					Id string `json:"id"`
				}

				decoder := json.NewDecoder(res.Body)
				if err := decoder.Decode(&responseData); err != nil {
					t.Fatalf("Failed to decode create user response: %v", err)
				}

				userId = responseData.Id
				// authToken = getAuthToken(app, userId)
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"name":"test"`,
			},
		},
		{
			Name:                  "guests should not be able to delete existing users",
			Method:                http.MethodDelete,
			URL:                   "/api/collections/users/records/" + userId,
			DisableTestAppCleanup: true,
			ExpectedStatus:        404,
			ExpectedContent:       []string{`"data":{}`},
			ExpectedEvents:        map[string]int{"*": 0},
		},
		// TODO: enable once brave to fix
		// {
		// 	Name:   "every user should be allowed to delete themselves",
		// 	Method: http.MethodDelete,
		// 	URL:    "/api/collections/users/records/" + userId,
		// 	Headers: map[string]string{
		// 		"Authorization": authToken,
		// 	},
		// 	DisableTestAppCleanup: true,
		// 	ExpectedStatus:        204,
		// 	ExpectedContent: []string{
		// 		`null`,
		// 	},
		// 	ExpectedEvents: map[string]int{
		// 		"*":                     0,
		// 		"OnRecordDeleteRequest": 1,
		// 	},
		// },
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
