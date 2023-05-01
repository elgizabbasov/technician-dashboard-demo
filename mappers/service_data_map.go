package mappers

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/elgizabbasov/technician-dashboard-demo/initializers"
	"github.com/elgizabbasov/technician-dashboard-demo/models"
)

type ServiceData = models.ServiceData

// Get ServiceData
func ServiceDataGet() ([]ServiceData, error) {
	rows, err := initializers.DB().Query("SELECT id, data, service_id, created, updated from servicedata")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	servicedata := make([]ServiceData, 0)

	for rows.Next() {
		singleServiceData := ServiceData{}
		err = rows.Scan(&singleServiceData.ID, &singleServiceData.Data, &singleServiceData.ServiceID, &singleServiceData.Created, &singleServiceData.Updated)

		if err != nil {
			return nil, err
		}

		servicedata = append(servicedata, singleServiceData)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return servicedata, err
}

// Get ServiceData by range of dates
func ServiceDataGetByDate(start_time string, end_time string) ([]ServiceData, error) {
	//sort by date not id
	start, _ := time.Parse("2006-01-02", start_time)
	end, _ := time.Parse("2006-01-02", end_time)
	stmt, err := initializers.DB().Query("SELECT id, data, service_id, created, updated from servicedata WHERE created BETWEEN ? AND ?", start, end)

	if err != nil {
		return []ServiceData{}, err
	}
	defer stmt.Close()

	servicedata := make([]ServiceData, 0)

	for stmt.Next() {
		singleServiceData := ServiceData{}
		err = stmt.Scan(&singleServiceData.ID, &singleServiceData.Data, &singleServiceData.ServiceID, &singleServiceData.Created, &singleServiceData.Updated)

		if err != nil {
			return []ServiceData{}, err
		}

		servicedata = append(servicedata, singleServiceData)

	}

	if err != nil {
		return []ServiceData{}, err
	}

	return servicedata, err
}

func ServiceDataGetByServiceID(serviceid string) ([]ServiceData, error) {
	stmt, err := initializers.DB().Query("SELECT id, data, service_id, created, updated from servicedata WHERE service_id = ?", serviceid)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	servicedata := make([]ServiceData, 0)

	for stmt.Next() {
		singleServiceData := ServiceData{}
		err = stmt.Scan(&singleServiceData.ID, &singleServiceData.Data, &singleServiceData.ServiceID, &singleServiceData.Created, &singleServiceData.Updated)
		if err != nil {
			return nil, err
		}

		servicedata = append(servicedata, singleServiceData)
	}

	if stmt.Err() != nil {
		return nil, stmt.Err()
	}
	return servicedata, nil
}

// Add a new servicedata to the servicedata table
func ServiceDataAdd(newServiceData ServiceData) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO servicedata (id, data, service_id, created, updated, pearlash, mechanicalroom, pails) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	// TODO: need additional input filtering and error handling
	_, err = stmt.Exec(newServiceData.ID, newServiceData.Data, newServiceData.ServiceID, newServiceData.Created, newServiceData.Updated, strings.Join(newServiceData.Pearlash, ","), strings.Join(newServiceData.Mechanicalroom, ","), strings.Join(newServiceData.Pails, ","))

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

// Upload a new image to the given path
func ServiceDataImageAdd(image *multipart.FileHeader, localFolder string) error {
	// Open the uploaded file
	file, err := image.Open()
	if err != nil {
		return fmt.Errorf("error opening the file: %v", err)
	}
	defer file.Close()

	// Read the contents of the file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading the file: %v", err)
	}

	// Check if the folder exists, if not, create it
	_, err = os.Stat(localFolder)
	if os.IsNotExist(err) {
		err = os.MkdirAll(localFolder, 0755)
		if err != nil {
			return fmt.Errorf("error creating folder: %v", err)
		}
	}

	// Save the file to the folder
	filePath := filepath.Join(localFolder, image.Filename)
	err = ioutil.WriteFile(filePath, data, 0755)
	if err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}

	return nil
}

func ServiceDataUpdate(aServiceData ServiceData, id string) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE servicedata SET data = ?, service_id = ?, updated = ? WHERE id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	// TODO: need additional input filtering and error handling
	_, err = stmt.Exec(aServiceData.Data, aServiceData.ServiceID, aServiceData.Updated, id)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func ServiceDataDelete(id string) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("DELETE from servicedata WHERE id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	// TODO: need additional input filtering & error handling
	_, err = stmt.Exec(id)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
