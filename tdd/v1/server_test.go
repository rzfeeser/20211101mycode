/* Alta3 Research | RZFeeser
   TDD - Building tests to drive feature development (refactored)

   This test creates a map to store key/values in (string: int). This test data
   may then be used to test our application. */

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
)

// stub up a place to store key/values we wish to pass into our program
type StubPlayerStore struct {
	scores map[string]int   // easy way to pass key/value pairs
}

// create a method for our structure
// returns the player score
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

// create a store for our test
func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}


// new function added to help our testing function (added during refactoring)
// looks up players by name and returns their score
func newGetScoreRequest(name string) *http.Request {
        req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
        return req
}

// new function added to help our testing function (added during refactoring)
// ensures the body contains the value we expect
func assertResponseBody(t testing.TB, got, want string) {
        t.Helper()   // Helper marks the calling function as a test helper function
                     // When printing file and line information, that function will be skipped.
        if got != want {
                t.Errorf("response body is wrong, got %q want %q", got, want)
        }
}
