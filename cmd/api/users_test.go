package main

import (
	"log"
	"net/http"
	"testing"
)

func TestGetUser(t *testing.T) {
	app := newTestApplication(t)
	mux := app.mount()

	testToken, err := app.authenticator.GenerateToken(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("should not allow unathenticated requests", func(t *testing.T) {
		// check for 401 status code
		req, err := http.NewRequest(http.MethodGet, "/v1/users/2", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := exeuteRequest(req, mux)

		// if rr.Code != http.StatusUnauthorized {
		// 	t.Errorf("expected the response code to be %d and we got %d", http.StatusUnauthorized, rr.Code)
		// }
		checkResponseCode(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("should allow authenticated requests", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/users/2", nil)
		if err != nil {
			t.Fatal(err)
		}

		// rr := exeuteRequest(req, mux)
		req.Header.Set("Authorization", "Bearer "+testToken)
		rr := exeuteRequest(req, mux)

		// if rr.Code != http.StatusUnauthorized {
		// 	t.Errorf("expected the response code to be %d and we got %d", http.StatusUnauthorized, rr.Code)
		// }
		checkResponseCode(t, http.StatusOK, rr.Code)

		log.Println(rr.Body)

	})
}
