package mappers

import (
	"time"

	"github.com/elgizabbasov/technician-dashboard-demo/initializers"
	"github.com/elgizabbasov/technician-dashboard-demo/models"
)

type SensorData = models.SensorData

// Get SensorData
func SensorDataGet() ([]SensorData, error) {
	rows, err := initializers.DB().Query("SELECT id, sensor_id, data, created FROM sensordata")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	sensordata := make([]SensorData, 0)

	for rows.Next() {
		singleSensorData := SensorData{}
		err = rows.Scan(&singleSensorData.ID, &singleSensorData.SensorID, &singleSensorData.Data, &singleSensorData.Created)

		if err != nil {
			return nil, err
		}

		sensordata = append(sensordata, singleSensorData)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return sensordata, err
}

// Get SensorData set by Datetime
func SensorDataGetByDate(start_time string, end_time string) ([]SensorData, error) {
	//sort by date not id
	start, _ := time.Parse("2006-01-02", start_time)
	end, _ := time.Parse("2006-01-02", end_time)
	stmt, err := initializers.DB().Query("SELECT id, sensor_id, data, created from sensordata WHERE created BETWEEN ? AND ?", start, end)

	if err != nil {
		return []SensorData{}, err
	}
	defer stmt.Close()

	sensordata := make([]SensorData, 0)

	for stmt.Next() {
		singleSensorData := SensorData{}
		err = stmt.Scan(&singleSensorData.ID, &singleSensorData.SensorID, &singleSensorData.Data, &singleSensorData.Created)

		if err != nil {
			return []SensorData{}, err
		}

		sensordata = append(sensordata, singleSensorData)

	}

	if err != nil {
		return []SensorData{}, err
	}

	return sensordata, err
}

// Optimized for graph fetching (Only gets latest 100 rows of data)
func SensorDataGetLatest(rows string, sensorid string) ([]SensorData, error) {
	stmt, err := initializers.DB().Query("SELECT id, sensor_id, data, created FROM sensordata WHERE sensor_id = ? ORDER BY created DESC LIMIT ?", sensorid, rows)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	sensordata := make([]SensorData, 0)

	for stmt.Next() {
		singleSensorData := SensorData{}
		err = stmt.Scan(&singleSensorData.ID, &singleSensorData.SensorID, &singleSensorData.Data, &singleSensorData.Created)

		if err != nil {
			return nil, err
		}

		sensordata = append(sensordata, singleSensorData)
	}

	err = stmt.Err()

	if err != nil {
		return nil, err
	}

	return sensordata, err
}

// Add a new sensor data to the sensor table
func SensorDataAdd(newSensorData SensorData) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare(`INSERT INTO sensordata(id, sensor_id, data, created) VALUES (?, ?, ?, ?)`)

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	// TODO: need additional input filtering and error handling
	_, err = stmt.Exec(newSensorData.ID, newSensorData.SensorID, newSensorData.Data, newSensorData.Created)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func SensorDataUpdate() {
}

func SensorDataDelete(id string) (bool, error) {
	tx, err := initializers.DB().Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("DELETE from sensordata WHERE id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	// TODO: need additional input filtering
	_, err = stmt.Exec(id)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
