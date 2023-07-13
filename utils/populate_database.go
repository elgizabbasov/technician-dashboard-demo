package utils

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/elgizabbasov/technician-dashboard-demo/initializers"
	"github.com/google/uuid"
)

// POPULATING sqlite w a table and data -- Temporary table creation code
func PopulateTables() {
	// open the SQLite database
	db, err := sql.Open("sqlite3", "sqlite-database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// check if the 'site' table is empty
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM site").Scan(&count)
	if err != nil {
		panic(err)
	}

	// if the table is empty, then insert some data
	if count == 0 {
		// mock UUID generation
		siteid1 := uuid.New().String()
		siteid2 := uuid.New().String()

		InsertSite(siteid1, "Sleepy Hallow", "Cleano2", "555 Ave SW.", "T2T4G5", "Calgary", "Alberta", "Canada", "2021-10-14 13:57:01", "2020-10-14 13:57:01")
		InsertSite(siteid2, "Ramses", "Cleano2", "554 Ave SW.", "H3T5G6", "Calgary", "Alberta", "Canada", "2021-10-14 13:57:02", "2020-10-14 13:57:02")

		// mock UUID generation
		deviceid1 := uuid.New().String()
		deviceid2 := uuid.New().String()
		deviceid3 := uuid.New().String()
		deviceid4 := uuid.New().String()
		deviceid5 := uuid.New().String()

		InsertDevice(deviceid1, "Calgary North", "3.4", siteid1, "8989898", "2019-10-14 13:57:01", "2020-10-14 13:57:01")
		InsertDevice(deviceid2, "Calgary South", "3.7", siteid1, "8989896", "2019-11-14 13:57:01", "2020-10-14 13:57:01")
		InsertDevice(deviceid3, "Calgary West", "3.6", siteid1, "8989895", "2019-12-14 13:57:01", "2020-10-14 13:57:01")
		InsertDevice(deviceid4, "Calgary East", "3.5", siteid1, "8989892", "2019-01-14 13:57:01", "2020-10-14 13:57:01")
		InsertDevice(deviceid5, "Calgary Edmonton", "4.0", siteid2, "8989854", "2021-10-14 13:57:01", "2020-10-14 13:57:01")

		// mock UUID generation
		sensorid1 := uuid.New().String()
		sensorid2 := uuid.New().String()
		sensorid3 := uuid.New().String()
		sensorid4 := uuid.New().String()
		sensorid5 := uuid.New().String()
		sensorid6 := uuid.New().String()

		InsertSensor(sensorid1, "CalgaryPSensor1", "kPa", "Pressure", deviceid1, "2021-10-14 13:57:02", "2020-10-14 13:57:02")
		InsertSensor(sensorid2, "CalgaryTSensor1", "celsius", "Temperature", deviceid1, "2021-10-14 13:57:02", "2020-10-14 13:57:02")
		InsertSensor(sensorid3, "CalgaryTSensor2", "g.m^-3", "Humidity", deviceid1, "2021-10-14 13:57:02", "2020-10-14 13:57:02")
		InsertSensor(sensorid4, "CalgaryPSensor2", "kPa", "Pressure", deviceid1, "2021-10-14 13:57:02", "2020-10-14 13:57:02")
		InsertSensor(sensorid5, "CalgaryTSensor3", "celsius", "Temperature", deviceid2, "2021-10-14 13:57:02", "2020-10-14 13:57:02")
		InsertSensor(sensorid6, "CalgaryTSensor4", "g.m^-3", "Humidity", deviceid2, "2021-10-14 13:57:02", "2020-10-14 13:57:02")

		// mock 1000 rows of sensordata generation
		for i := 0; i < 2000; i++ {
			uuid8 := uuid.New().String()
			// convert random float to string with 3 decimal points
			data := strconv.FormatFloat(rand.Float64(), 'f', 3, 64)

			// between current time and a week ago
			now := time.Now().Unix()
			weekAgo := now - 604800
			randomTimestamp := rand.Int63n(now-weekAgo+1) + weekAgo
			created := time.Unix(randomTimestamp, 0).Format("2006-01-02 15:04:05")

			// fifth of mock data with first mock sensor id
			if i < 333 {
				InsertSensorData(uuid8, sensorid1, data, created)
			} else if i > 333 && i < 666 {
				InsertSensorData(uuid8, sensorid2, data, created)
			} else if i > 666 && i < 999 {
				InsertSensorData(uuid8, sensorid3, data, created)
			} else if i > 999 && i < 1333 {
				InsertSensorData(uuid8, sensorid4, data, created)
			} else if i > 1333 && i < 1666 {
				InsertSensorData(uuid8, sensorid5, data, created)
			} else {
				InsertSensorData(uuid8, sensorid6, data, created)
			}
		}

		// mock UUID generation
		serviceid1 := uuid.New().String()
		serviceid2 := uuid.New().String()
		serviceid3 := uuid.New().String()
		serviceid4 := uuid.New().String()
		serviceid5 := uuid.New().String()

		servicedataid1 := uuid.New().String()
		servicedataid2 := uuid.New().String()

		// important: MAKING SURE THAT ALL THE DEVICES HAVE AT LEAST 1 SERVICE,
		// OTHERWISE Tech Form DATA WONT BE SUBMITTED
		InsertService(serviceid1, "Checklist service", deviceid1, "2021-10-14 13:57:02", "2020-10-14 13:57:02")
		InsertService(serviceid2, "Checklist service", deviceid2, "2021-10-14 13:57:02", "2020-10-14 13:57:02")
		InsertService(serviceid3, "Checklist service", deviceid3, "2021-10-14 13:57:02", "2020-10-14 13:57:02")
		InsertService(serviceid4, "Checklist service", deviceid4, "2021-10-14 13:57:02", "2020-10-14 13:57:02")
		InsertService(serviceid5, "Checklist service", deviceid5, "2021-10-14 13:57:02", "2020-10-14 13:57:02")

		InsertServiceData(servicedataid1, "technicianReportCardInputsConcatenated?", serviceid1, "2021-10-14 13:57:02", "2020-10-14 13:57:02")
		InsertServiceData(servicedataid2, "technicianReportCardInputsConcatenated2", serviceid2, "2022-11-05 13:57:02", "2020-10-14 13:57:02")

		// sqlite test end
	} else {
		fmt.Println("Table already has data")
	}
}

func CreateSiteTable() {
	createSiteTableSQL := `CREATE TABLE IF NOT EXISTS site (
		"id" TEXT NOT NULL PRIMARY KEY,		
		"name" TEXT,
		"organization" TEXT,
		"address" TEXT,	
		"postal" TEXT,
		"city" TEXT,
		"province" TEXT,
		"country" TEXT,
		"created" TEXT,
		"updated" TEXT
	  );` // SQL Statement for Create Table

	log.Println("Create site table...")
	statement, err := initializers.DB().Prepare(createSiteTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("site table created")
}

func CreateDeviceTable() {
	createDeviceTableSQL := `CREATE TABLE IF NOT EXISTS device (
		"id" TEXT NOT NULL PRIMARY KEY,		
		"name" TEXT,
		"version" TEXT,
		"site_id" TEXT,	
		"serial_number" TEXT,
		"created" TEXT,
		"updated" TEXT,
		FOREIGN KEY (site_id) REFERENCES site (id)
	  );` // SQL Statement for Create Table

	log.Println("Create device table...")
	statement, err := initializers.DB().Prepare(createDeviceTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("device table created")
}

func CreateSensorTable() {
	createSensorTableSQL := `CREATE TABLE IF NOT EXISTS sensor (
		"id" TEXT NOT NULL PRIMARY KEY,		
		"name" TEXT,
		"unit" TEXT,
		"type" TEXT,	
		"device_id" TEXT,
		"created" TEXT,
		"updated" TEXT, 
		FOREIGN KEY (device_id) REFERENCES device (id)
	  );` // SQL Statement for Create Table

	log.Println("Create sensor table...")
	statement, err := initializers.DB().Prepare(createSensorTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("sensor table created")
}

func CreateSensorDataTable() {
	createSensorDataTableSQL := `CREATE TABLE IF NOT EXISTS sensordata (
		"id" TEXT NOT NULL PRIMARY KEY,		
		"sensor_id" TEXT,
		"data" TEXT,
		"created" TEXT,
		FOREIGN KEY (sensor_id) REFERENCES sensor (id)
	  );` // SQL Statement for Create Table

	log.Println("Create Sensor Data table...")
	statement, err := initializers.DB().Prepare(createSensorDataTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("Sensor Data table created")
}

func CreateServiceTable() {
	createServiceTableSQL := `CREATE TABLE IF NOT EXISTS service (
		"id" TEXT NOT NULL PRIMARY KEY,		
		"name" TEXT,
		"device_id" TEXT,	
		"created" TEXT,
		"updated" TEXT, 
		FOREIGN KEY (device_id) REFERENCES device (id)
	  );` // SQL Statement for Create Table

	log.Println("Create service table...")
	statement, err := initializers.DB().Prepare(createServiceTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("service table created")
}

func CreateServiceDataTable() {
	createServiceDataTableSQL := `CREATE TABLE IF NOT EXISTS servicedata (
		"id" TEXT NOT NULL PRIMARY KEY,		
		"data" TEXT,
		"service_id" TEXT,
		"created" TEXT,
		"updated" TEXT,
		"pearlash" TEXT,
		"mechanicalroom" TEXT,
		"pails" TEXT,
		FOREIGN KEY (service_id) REFERENCES service (id)
	  );` // SQL Statement for Create Table

	log.Println("Create Service Data table...")
	statement, err := initializers.DB().Prepare(createServiceDataTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("Service Data table created")
}

// TODO: SEPERATE IMAGES SO MORE THAN 1 IMAGE CAN BE STORED
// INSTEAD OF 1 BLOB
// func createServiceDataImageTable() {
// 	createServiceDataImageTable := `CREATE TABLE servicedata_images (
// 		id TEXT NOT NULL PRIMARY KEY,
// 		servicedata_id TEXT NOT NULL,
// 		image_data BLOB,
// 		FOREIGN KEY (servicedata_id) REFERENCES servicedata (id)
// 	  );` // SQL Statement for Create Table

// 	log.Println("Create Service Data Image table...")
// 	statement, err := initializers.DB().Prepare(createServiceDataImageTable) // Prepare SQL Statement
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	statement.Exec() // Execute SQL Statements
// 	log.Println("Service Data Image table created")
// }

// *** SQLITE3 Database Prepopulation helpers below ***
func InsertDevice(id string, name string, version string, siteid string, serialnumber string, created string, updated string) {
	log.Println("Inserting device record ...")
	insertStudentSQL := `INSERT INTO device(id, name, version, site_id, serial_number, created, updated) VALUES (?, ?, ?, ?, ?, ?, ?)`
	statement, err := initializers.DB().Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, name, version, siteid, serialnumber, created, updated)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func InsertSite(id string, name string, organization string, address string, postal string, city string, province string, country string, created string, updated string) {
	log.Println("Inserting site record ...")
	insertStudentSQL := `INSERT INTO site(id, name, organization, address, postal, city, province, country, created, updated) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	statement, err := initializers.DB().Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, name, organization, address, postal, city, province, country, created, updated)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func InsertService(id string, name string, device_id string, created string, updated string) {
	log.Println("Inserting service record ...")
	insertStudentSQL := `INSERT INTO service(id, name, device_id, created, updated) VALUES (?, ?, ?, ?, ?)`
	statement, err := initializers.DB().Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, name, device_id, created, updated)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func InsertServiceData(id string, data string, service_id string, created string, updated string) {
	log.Println("Inserting service data record ...")
	insertStudentSQL := `INSERT INTO servicedata(id, data, service_id, created, updated) VALUES (?, ?, ?, ?, ?)`
	statement, err := initializers.DB().Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, data, service_id, created, updated)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func InsertSensor(id string, name string, unit string, _type string, device_id string, created string, updated string) {
	log.Println("Inserting sensor record ...")
	insertStudentSQL := `INSERT INTO sensor(id, name, unit, type, device_id, created, updated) VALUES (?, ?, ?, ?, ?, ?, ?)`
	statement, err := initializers.DB().Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, name, unit, _type, device_id, created, updated)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func InsertSensorData(id string, sensor_id string, data string, created string) {
	log.Println("Inserting Sensor Data record ...")
	InsertSensorDataSQL := `INSERT INTO sensordata(id, sensor_id, data, created) VALUES (?, ?, ?, ?)`
	statement, err := initializers.DB().Prepare(InsertSensorDataSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, sensor_id, data, created)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
