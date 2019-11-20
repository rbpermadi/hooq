package sitecheckapp

import (
	"encoding/json"
	"log"
	"net/http"
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

	newSite.Status = sitecheck.Check(newSite.Link)

	site := index.SiteRepo.Create(newSite)

	response := Response{
		Data: site,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(response)
}

func (index *Index) Loop() {
	for {
		sites := index.SiteRepo.All()
		for idx, site := range sites {
			site.Status = site_check.Check(site.Link)
			index.SiteRepo.Update(idx, site)
		}
		time.Sleep(5 * time.Minute)
	}
}
