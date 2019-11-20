package sitecheckapp

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rbpermadi/hooq/internal/repository"
)

type Index struct {
	SiteRepo repository.ISiteRepo
}

type Response struct {
	Data interface{} `json:"data"`
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
	}

	newSite.Status = "unchecked"

	site := index.SiteRepo.Create(newSite)

	response := Response{
		Data: site,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(response)
}
