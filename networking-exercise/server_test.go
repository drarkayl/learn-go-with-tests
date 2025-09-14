package networkingexercise

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// Create a new HTTP request for the "/" route.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()

	// Call the handler function directly, passing the ResponseRecorder.
	// This simulates the server calling the handler.
	helloHandler(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body.
	expected := "Hello, World!"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// We'll write a new test here that actually starts the server.
// Note: This is a separate concept from the above unit test for the handler.
func TestServerIntegration(t *testing.T) {
	// This is a simple integration test that uses httptest to create a test server.
	ts := httptest.NewServer(http.HandlerFunc(helloHandler))
	defer ts.Close() // Close the server when the test finishes.

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code was incorrect, got: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}
