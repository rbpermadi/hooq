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

func (siteRepo *SiteRepo) Update(id int, site Site) (*Site, error) {
	for i := range siteRepo.Sites {
		if siteRepo.Sites[i].ID == id {
			siteRepo.Sites[i] = site
			return &site, nil
		}
	}

	return nil, fmt.Errorf("Site not found")
}

func (siteRepo *SiteRepo) Delete(id int) error {
	sites := []Site{}
	status := false

	for i := range siteRepo.Sites {
		if siteRepo.Sites[i].ID == id {
			status = true
		} else {
			sites = append(sites, siteRepo.Sites[i])
		}
	}

	if status == false {
		return fmt.Errorf("Site not found")
	}

	siteRepo.Sites = sites

	return nil
}
