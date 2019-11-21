package sitecheckapp_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/rbpermadi/hooq/internal/app/sitecheckapp"
	"github.com/rbpermadi/hooq/internal/repository"
)

func assertJSON(t *testing.T, actual []byte, data interface{}) {
	expected, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	//remove trailing new line
	actual = bytes.TrimRight(actual, "\n")

	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
}

func TestGetSiteCheckHandler(t *testing.T) {
	t.Run("data not null", func(t *testing.T) {
		site := repository.Site{
			ID:     1,
			Link:   "http://google.com",
			Status: "healthy",
		}

		siteRepo := repository.SiteRepo{
			Sites: []repository.Site{},
		}

		siteRepo.Sites = append(siteRepo.Sites, site)

		siteCheckApp := sitecheckapp.Index{
			SiteRepo: &siteRepo,
		}

		req, _ := http.NewRequest("GET", "/site_check", nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(siteCheckApp.GetSiteCheckHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := sitecheckapp.Response{
			Data: []repository.Site{},
		}

		expected.Data = siteRepo.Sites

		assertJSON(t, rr.Body.Bytes(), expected)
	})

	t.Run("data null", func(t *testing.T) {
		siteRepo := repository.SiteRepo{
			Sites: []repository.Site{},
		}

		siteCheckApp := sitecheckapp.Index{
			SiteRepo: &siteRepo,
		}

		req, _ := http.NewRequest("GET", "/site_check", nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(siteCheckApp.GetSiteCheckHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := sitecheckapp.Response{
			Data: []repository.Site{},
		}

		assertJSON(t, rr.Body.Bytes(), expected)
	})
}

func TestAddSiteCheckHandler(t *testing.T) {
	siteRepo := repository.SiteRepo{
		Sites: []repository.Site{},
	}

	siteCheckApp := sitecheckapp.Index{
		SiteRepo: &siteRepo,
	}

	t.Run("bad request : post data not json", func(t *testing.T) {

		jsonRequest, _ := json.Marshal("testtest")
		body := bytes.NewBuffer(jsonRequest)
		req, _ := http.NewRequest("POST", "/site_check/add", body)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(siteCheckApp.AddSiteCheckHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

	t.Run("bad request : url not failed", func(t *testing.T) {

		site := repository.Site{
			Link: "test123",
		}

		jsonRequest, _ := json.Marshal(site)
		body := bytes.NewBuffer(jsonRequest)
		req, _ := http.NewRequest("POST", "/site_check/add", body)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(siteCheckApp.AddSiteCheckHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

	})

	t.Run("success", func(t *testing.T) {

		site := repository.Site{
			Link: "http://google.com",
		}

		jsonRequest, _ := json.Marshal(site)
		body := bytes.NewBuffer(jsonRequest)
		req, _ := http.NewRequest("POST", "/site_check/add", body)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(siteCheckApp.AddSiteCheckHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

	})
}

func TestDeleteSiteCheckHandler(t *testing.T) {
	site := repository.Site{
		ID:     1,
		Link:   "http://google.com",
		Status: "healthy",
	}

	siteRepo := repository.SiteRepo{
		Sites: []repository.Site{},
	}

	siteRepo.Sites = append(siteRepo.Sites, site)

	siteCheckApp := sitecheckapp.Index{
		SiteRepo: &siteRepo,
	}

	t.Run("bad request : id must not null", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/site_check/delete", nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(siteCheckApp.DeleteSiteCheckHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := sitecheckapp.ResponseError{
			Code:    400,
			Message: "Id must not null",
		}

		assertJSON(t, rr.Body.Bytes(), expected)
	})

	t.Run("not found : site not found", func(t *testing.T) {

		req, _ := http.NewRequest("DELETE", "/site_check/delete", nil)

		req.URL.RawQuery = url.Values{"id": {"0"}}.Encode()

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(siteCheckApp.DeleteSiteCheckHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := sitecheckapp.ResponseError{
			Code:    404,
			Message: "Site not found",
		}

		assertJSON(t, rr.Body.Bytes(), expected)
	})

	t.Run("success", func(t *testing.T) {

		req, _ := http.NewRequest("DELETE", "/site_check/delete", nil)

		req.URL.RawQuery = url.Values{"id": {"1"}}.Encode()

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(siteCheckApp.DeleteSiteCheckHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := struct {
			Message string `json:"message"`
		}{
			Message: "site has successfully deleted",
		}

		assertJSON(t, rr.Body.Bytes(), expected)
	})
}

func TestHomePageHandler(t *testing.T) {
	siteRepo := repository.SiteRepo{
		Sites: []repository.Site{},
	}

	siteCheckApp := sitecheckapp.Index{
		SiteRepo: &siteRepo,
	}

	req, _ := http.NewRequest("GET", "", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(siteCheckApp.HomePageHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
