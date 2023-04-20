package mappers

import (
	"database/sql"

	"github.com/CleanO2tech/clean-technician-web/initializers"
	"github.com/CleanO2tech/clean-technician-web/models"
)

type Sensor = models.Sensor

// Get Sensors
func SensorsGet() ([]Sensor, error) {
	rows, err := initializers.DB().Query("SELECT id, name, unit, type, device_id, created, updated from sensor")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	sensor := make([]Sensor, 0)

	for rows.Next() {
		singleSensor := Sensor{}
		err = rows.Scan(&singleSensor.ID, &singleSensor.Name, &singleSensor.Unit, &singleSensor.Type, &singleSensor.DeviceID, &singleSensor.Created, &singleSensor.Updated)

		if err != nil {
			return nil, err
		}

		sensor = append(sensor, singleSensor)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return sensor, err
}

// Get Sensor by ID
func SensorGetByID(id string) (Sensor, error) {
	stmt, err := initializers.DB().Prepare("SELECT id, name, unit, type, device_id, created, updated from sensor WHERE id = ?")

	if err != nil {
		return Sensor{}, err
	}

	sensor := Sensor{}

	sqlErr := stmt.QueryRow(id).Scan(&sensor.ID, &sensor.Name, &sensor.Unit, &sensor.Type, &sensor.DeviceID, &sensor.Created, &sensor.Updated)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Sensor{}, nil
		}
		return Sensor{}, sqlErr
	}
	return sensor, nil
}

func SensorGetByDeviceID(deviceid string) ([]Sensor, error) {
	stmt, err := initializers.DB().Query("SELECT id, name, unit, type, device_id, created, updated from sensor WHERE device_id = ?", deviceid)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	sensor := make([]Sensor, 0)

	for stmt.Next() {
		singleSensor := Sensor{}
		err = stmt.Scan(&singleSensor.ID, &singleSensor.Name, &singleSensor.Unit, &singleSensor.Type, &singleSensor.DeviceID, &singleSensor.Created, &singleSensor.Updated)
		if err != nil {
			return nil, err
		}

		sensor = append(sensor, singleSensor)
	}

	if stmt.Err() != nil {
		return nil, stmt.Err()
	}
	return sensor, nil
}

// Get Sensor by Name
func SensorGetByName(name string) (Sensor, error) {
	stmt, err := initializers.DB().Prepare("SELECT id, name, unit, type, device_id, created, updated from sensor WHERE name = ?")

	if err != nil {
		return Sensor{}, err
	}

	sensor := Sensor{}

	sqlErr := stmt.QueryRow(name).Scan(&sensor.ID, &sensor.Name, &sensor.Unit, &sensor.Type, &sensor.DeviceID, &sensor.Created, &sensor.Updated)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Sensor{}, nil
		}
		return Sensor{}, sqlErr
	}
	return sensor, nil
}

// Add a new sensor to the sensor table
func SensorAdd(newSensor Sensor) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO sensor (id, name, unit, type, device_id, created, updated) VALUES (?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	// TODO: need additional input filtering and error handling
	_, err = stmt.Exec(newSensor.ID, newSensor.Name, newSensor.Unit, newSensor.Type, newSensor.DeviceID, newSensor.Created, newSensor.Updated)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func SensorUpdate(aSensor Sensor, id string) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE sensor SET name = ?, unit = ?, type = ?, device_id = ?, updated = ? WHERE id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	// TODO: need additional input filtering and error handling
	_, err = stmt.Exec(aSensor.Name, aSensor.Unit, aSensor.Type, aSensor.DeviceID, aSensor.Updated, id)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func SensorDelete(id string) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("DELETE from sensor WHERE id = ?")

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
