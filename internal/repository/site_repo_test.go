package repository_test

import (
	"testing"

	"github.com/rbpermadi/hooq/internal/repository"
)

func TestSiteRepo_All(t *testing.T) {
	siteRepo := repository.SiteRepo{}

	// test when data null
	t.Run("data null", func(t *testing.T) {
		sitesNull := siteRepo.All()
		if len(sitesNull) != 0 {
			t.Errorf("All() failed, expected %v, got %v", 0, len(sitesNull))
		}
	})

	// test when data not null
	t.Run("success", func(t *testing.T) {
		site := repository.Site{
			ID:     1,
			Link:   "http://google.com",
			Status: "healthy",
		}

		site2 := repository.Site{
			ID:     2,
			Link:   "http://github.com",
			Status: "unhealthy",
		}

		siteRepo.Create(site)
		siteRepo.Create(site2)

		sites := siteRepo.All()
		lastSite := sites[len(sites)-1]

		if site2.Link != lastSite.Link {
			t.Errorf("All() failed, expected %v, got %v", site2.Link, lastSite.Link)
		}
	})
}

func TestSiteRepo_Create(t *testing.T) {
	siteRepo := repository.SiteRepo{}

	newSite := repository.Site{
		Link:   "http://google.com",
		Status: "healthy",
	}

	site := siteRepo.Create(newSite)

	if newSite.Link != site.Link {
		t.Errorf("Create() failed, expected %v, got %v", newSite.Link, site.Link)
	}
}

func TestSiteRepo_Update(t *testing.T) {
	siteRepo := repository.SiteRepo{}

	newSite := repository.Site{
		Link:   "http://google.com",
		Status: "healthy",
	}

	site := siteRepo.Create(newSite)

	site.Status = "unhealhty"

	// test when data not null
	t.Run("success", func(t *testing.T) {
		updateResult, err := siteRepo.Update(0, *site)

		if err != nil {
			t.Errorf("Update() failed, expected %v, got %v", "nil", err)
		}

		if site.Status != updateResult.Status {
			t.Errorf("Update() failed, expected %v, got %v", site.Status, updateResult.Status)
		}
	})

	// test when data null
	t.Run("data not found", func(t *testing.T) {
		_, err := siteRepo.Update(10, *site)

		if err == nil {
			t.Errorf("Update() failed, expected %v, got %v", "not nil", err)
		}

		if err.Error() != "Site not found" {
			t.Errorf("Update() failed, expected %v, got %v", "Site not found", err.Error())
		}
	})
}
