package sitecheck_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rbpermadi/hooq/pkg/sitecheck"
)

func TestCheck(t *testing.T) {
	// test when data null
	t.Run("healthy", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`OK`))
			w.WriteHeader(200)
		}))

		defer server.Close()

		sc := sitecheck.Check(server.URL)

		if sc != "healthy" {
			t.Errorf("Check() failed, expected %v, got %v", "healthy", sc)
		}
	})

	t.Run("unhealthy", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(1 * time.Second)
			w.Write([]byte(`OK`))
			w.WriteHeader(500)
		}))

		defer server.Close()

		sc := sitecheck.Check(server.URL)

		if sc != "unhealthy" {
			t.Errorf("Check() failed, expected %v, got %v", "unhealthy", sc)
		}
	})
}
