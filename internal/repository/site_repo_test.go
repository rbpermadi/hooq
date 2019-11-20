package repository_test

import (
	"testing"

	"github.com/rbpermadi/hooq/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestSiteRepo_All(t *testing.T) {
	siteRepo := repository.SiteRepo{}

	sitesNull := siteRepo.All()

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

	// test when data not null
	t.Run("success", func(t *testing.T) {
		assert.Equal(t, site2.Link, lastSite.Link)
	})

	// test when data null
	t.Run("data null", func(t *testing.T) {
		assert.True(t, len(sitesNull) == 0)
	})
}

func TestSiteRepo_Create(t *testing.T) {
	siteRepo := repository.SiteRepo{}

	newSite := repository.Site{
		Link:   "http://google.com",
		Status: "healthy",
	}

	site := siteRepo.Create(newSite)

	assert.Equal(t, newSite.Link, site.Link)
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

		assert.Nil(t, err)
		assert.Equal(t, site.Status, updateResult.Status)
	})

	// test when data null
	t.Run("data null", func(t *testing.T) {
		_, err := siteRepo.Update(10, *site)

		assert.NotNil(t, err)
		assert.Equal(t, "Site not found", err.Error())
	})
}
