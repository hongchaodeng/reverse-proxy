package proxy

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProxyImpl_RequestHandler(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// Create a new proxy
	proxy, err := New(ts.URL)
	if err != nil {
		t.Fatalf("Failed to create proxy: %v", err)
	}

	// Create a request to the proxy
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the request handler
	handler := proxy.RequestHandler()
	handler(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, want %v", status, http.StatusOK)
	}
}
