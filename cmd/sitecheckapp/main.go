package main

import (
	"log"
	"net/http"

	"github.com/rbpermadi/hooq/internal/app/sitecheckapp"
	"github.com/rbpermadi/hooq/internal/repository"
)

func main() {
	siteRepo := repository.SiteRepo{
		Sites: []repository.Site{},
	}

	siteCheckApp := sitecheckapp.Index{
		SiteRepo: &siteRepo,
	}

	http.HandleFunc("/site_check", siteCheckApp.GetSiteCheckHandler)
	http.HandleFunc("/site_check/add", siteCheckApp.AddSiteCheckHandler)
	log.Fatal(http.ListenAndServe("localhost:7171", nil))
}
