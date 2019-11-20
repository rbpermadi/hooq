package repository

import "fmt"

// entity
type Site struct {
	ID     int    `json:"id"`
	Link   string `json:"link"`
	Status string `json:"status"`
}

type ISiteRepo interface {
	All() []Site
	Create(Site) *Site
	Update(int, Site) (*Site, error)
}

type SiteRepo struct {
	Sites []Site
}

var _ ISiteRepo = (*SiteRepo)(nil)

func (siteRepo *SiteRepo) All() []Site {
	return siteRepo.Sites
}

func (siteRepo *SiteRepo) Create(site Site) *Site {
	site.ID = len(siteRepo.Sites) + 1
	siteRepo.Sites = append(siteRepo.Sites, site)
	return &site
}

func (siteRepo *SiteRepo) Update(index int, site Site) (*Site, error) {
	if index > len(siteRepo.Sites) {
		return nil, fmt.Errorf("Site not found")
	}

	siteRepo.Sites[index] = site

	return &site, nil
}
