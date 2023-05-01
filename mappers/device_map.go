package mappers

import (
	"database/sql"

	"github.com/elgizabbasov/technician-dashboard-demo/initializers"
	"github.com/elgizabbasov/technician-dashboard-demo/models"
)

type Device = models.Device

// Get Devices
func DevicesGet() ([]Device, error) {
	rows, err := initializers.DB().Query("SELECT id, name, version, site_id, serial_number, created, updated from device")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	device := make([]Device, 0)

	for rows.Next() {
		singleDevice := Device{}
		err = rows.Scan(&singleDevice.ID, &singleDevice.Name, &singleDevice.Version, &singleDevice.SiteID, &singleDevice.SerialNumber, &singleDevice.Created, &singleDevice.Updated)

		if err != nil {
			return nil, err
		}

		device = append(device, singleDevice)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return device, err
}

// Get Device by ID
func DeviceGetByID(id string) (Device, error) {
	stmt, err := initializers.DB().Prepare("SELECT id, name, version, site_id, serial_number, created, updated from device WHERE id = ?")

	if err != nil {
		return Device{}, err
	}

	device := Device{}

	sqlErr := stmt.QueryRow(id).Scan(&device.ID, &device.Name, &device.Version, &device.SiteID, &device.SerialNumber, &device.Created, &device.Updated)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Device{}, nil
		}
		return Device{}, sqlErr
	}
	return device, nil
}

// Get Device by Name
func DeviceGetByName(name string) (Device, error) {
	stmt, err := initializers.DB().Prepare("SELECT id, name, created, updated from device WHERE name = ?")

	if err != nil {
		return Device{}, err
	}

	device := Device{}

	sqlErr := stmt.QueryRow(name).Scan(&device.ID, &device.Name, &device.Created, &device.Updated)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Device{}, nil
		}
		return Device{}, sqlErr
	}
	return device, nil
}

// Add a new device to the device table
func DeviceAdd(newDevice Device) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO device (id, name, version, site_id, serial_number, created, updated) VALUES (?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	// TODO: need additional input filtering and error handling
	_, err = stmt.Exec(newDevice.ID, newDevice.Name, newDevice.Version, newDevice.SiteID, newDevice.SerialNumber, newDevice.Created, newDevice.Updated)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func DeviceUpdate(aDevice Device, id string) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE device SET name = ?, version = ?, site_id = ?, serial_number = ?, updated = ? WHERE id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	// TODO: need additional input filtering and error handling
	_, err = stmt.Exec(aDevice.Name, aDevice.Version, aDevice.SiteID, aDevice.SerialNumber, aDevice.Updated, id)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func DeviceDelete(id string) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("DELETE from device WHERE id = ?")

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
