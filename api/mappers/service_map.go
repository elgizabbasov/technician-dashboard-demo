package mappers

import (
	"database/sql"

	"github.com/CleanO2tech/clean-technician-web/initializers"
	"github.com/CleanO2tech/clean-technician-web/models"
)

type Service = models.Service

// Get Services
func ServicesGet() ([]Service, error) {
	rows, err := initializers.DB().Query("SELECT id, name, device_id, created, updated from service")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	service := make([]Service, 0)

	for rows.Next() {
		singleService := Service{}
		err = rows.Scan(&singleService.ID, &singleService.Name, &singleService.DeviceID, &singleService.Created, &singleService.Updated)

		if err != nil {
			return nil, err
		}

		service = append(service, singleService)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return service, err
}

// Get Service by ID
func ServiceGetByID(id string) (Service, error) {
	stmt, err := initializers.DB().Prepare("SELECT id, name, device_id, created, updated from service WHERE id = ?")

	if err != nil {
		return Service{}, err
	}

	service := Service{}

	sqlErr := stmt.QueryRow(id).Scan(&service.ID, &service.Name, &service.DeviceID, &service.Created, &service.Updated)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Service{}, nil
		}
		return Service{}, sqlErr
	}
	return service, nil
}

func ServiceGetByDeviceID(deviceid string) (Service, error) {
	stmt, err := initializers.DB().Prepare("SELECT id, name, device_id, created, updated from service WHERE device_id = ?")

	if err != nil {
		return Service{}, err
	}

	service := Service{}

	sqlErr := stmt.QueryRow(deviceid).Scan(&service.ID, &service.Name, &service.DeviceID, &service.Created, &service.Updated)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Service{}, nil
		}
		return Service{}, sqlErr
	}
	return service, nil
}

// Add a new service to the service table
func ServiceAdd(newService Service) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO service (id, name, device_id, created, updated) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	// TODO: need additional input filtering and error handling
	_, err = stmt.Exec(newService.ID, newService.Name, newService.DeviceID, newService.Created, newService.Updated)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func ServiceUpdate(aService Service, id string) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE service SET name = ?, device_id = ?, updated = ? WHERE id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	// TODO: need additional input filtering and error handling
	_, err = stmt.Exec(aService.Name, aService.DeviceID, aService.Updated, id)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func ServiceDelete(id string) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("DELETE from service WHERE id = ?")

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
