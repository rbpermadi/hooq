package sitecheckapp

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/rbpermadi/hooq/internal/repository"
	"github.com/rbpermadi/hooq/pkg/sitecheck"
	"github.com/rbpermadi/hooq_test1/pkg/site_check"
)

type Index struct {
	SiteRepo repository.ISiteRepo
}

type Response struct {
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (index *Index) GetSiteCheckHandler(w http.ResponseWriter, r *http.Request) {
	sites := index.SiteRepo.All()

	response := Response{
		Data: sites,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}

func (index *Index) AddSiteCheckHandler(w http.ResponseWriter, r *http.Request) {
	var newSite repository.Site

	err := json.NewDecoder(r.Body).Decode(&newSite)

	if err != nil {
		log.Println("Error while reading JSON ", err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		response := ResponseError{
			Code:    400,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)

		return
	}

	_, err = url.ParseRequestURI(newSite.Link)

	if err != nil {
		log.Println("Error url validation : ", err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		response := ResponseError{
			Code:    400,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)

		return
	}

	newSite.Status = sitecheck.Check(newSite.Link)

	site := index.SiteRepo.Create(newSite)

	response := Response{
		Data: site,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(response)
}

func (index *Index) DeleteSiteCheckHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		response := ResponseError{
			Code:    400,
			Message: "Id must not null",
		}
		json.NewEncoder(w).Encode(response)

		return
	}

	err = index.SiteRepo.Delete(id)

	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		response := ResponseError{
			Code:    404,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	response := struct {
		Message string `json:"message"`
	}{
		Message: "site has successfully deleted",
	}

	json.NewEncoder(w).Encode(response)
}

func (index *Index) Loop() {
	for {
		sites := index.SiteRepo.All()
		for _, site := range sites {
			site.Status = site_check.Check(site.Link)
			index.SiteRepo.Update(site.ID, site)
		}

		time.Sleep(5 * time.Minute)
	}
}
