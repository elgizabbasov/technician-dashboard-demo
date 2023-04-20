package mappers

import (
	"database/sql"

	"github.com/elgizabbasov/technician-dashboard-demo/initializers"
	"github.com/elgizabbasov/technician-dashboard-demo/models"
)

type Site = models.Site

// Get Sites
func SitesGet() ([]Site, error) {
	rows, err := initializers.DB().Query("SELECT id, name, organization, address, postal, city, province, country, created, updated from site")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	site := make([]Site, 0)

	for rows.Next() {
		singleSite := Site{}
		err = rows.Scan(&singleSite.ID, &singleSite.Name, &singleSite.Organization, &singleSite.Address, &singleSite.Postal, &singleSite.City, &singleSite.Province, &singleSite.Country, &singleSite.Created, &singleSite.Updated)

		if err != nil {
			return nil, err
		}

		site = append(site, singleSite)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return site, err
}

// Get Site by ID
func SiteGetByID(id string) (Site, error) {
	stmt, err := initializers.DB().Prepare("SELECT id, name, organization, address, postal, city, province, country, created, updated from site WHERE id = ?")

	if err != nil {
		return Site{}, err
	}

	site := Site{}

	sqlErr := stmt.QueryRow(id).Scan(&site.ID, &site.Name, &site.Organization, &site.Address, &site.Postal, &site.City, &site.Province, &site.Country, &site.Created, &site.Updated)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Site{}, nil
		}
		return Site{}, sqlErr
	}
	return site, nil
}

// Get Site by Name
func SiteGetByName(name string) (Site, error) {
	stmt, err := initializers.DB().Prepare("SELECT id, name, created, updated from site WHERE name = ?")

	if err != nil {
		return Site{}, err
	}

	site := Site{}

	sqlErr := stmt.QueryRow(name).Scan(&site.ID, &site.Name, &site.Created, &site.Updated)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Site{}, nil
		}
		return Site{}, sqlErr
	}
	return site, nil
}

// Add a new site to the site table
func SiteAdd(newSite Site) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO site (id, name, organization, address, postal, city, province, country, created, updated) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	// TODO: need additional input filtering and error handling
	_, err = stmt.Exec(newSite.ID, newSite.Name, newSite.Organization, newSite.Address, newSite.Postal, newSite.City, newSite.Province, newSite.Country, newSite.Created, newSite.Updated)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func SiteUpdate(aSite Site, id string) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE site SET name = ?, organization = ?, address = ?, postal = ?, city = ?, province = ?, country = ?, updated = ? WHERE id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	// TODO: need additional input filtering and error handling
	_, err = stmt.Exec(aSite.Name, aSite.Organization, aSite.Address, aSite.Postal, aSite.City, aSite.Province, aSite.Country, aSite.Updated, id)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func SiteDelete(id string) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("DELETE from site WHERE id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	// TODO: need additional input filtering and error handling
	_, err = stmt.Exec(id)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
