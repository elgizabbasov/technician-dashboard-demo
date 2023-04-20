package initializers

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// func CreateDB() error {
// 	os.Remove("sqlite-database.db") // Delete the file to avoid duplicated records.
// 	log.Println("Creating sqlite-database.db...")

// 	file, err := os.Create("sqlite-database.db") // Create SQLite file

// 	if err != nil {
// 		return err
// 	}

// 	file.Close()
// 	log.Println("sqlite-database.db created")
// 	return nil
// }

func ConnectDB() error {
	// _foreign_keys=on enables FK
	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db?_foreign_keys=on") // TODO: Parametrize name of db file?

	if err != nil {
		return err
	}

	if err = sqliteDatabase.Ping(); err != nil {
		return err
	}

	db = sqliteDatabase
	return nil
}

func DB() *sql.DB {
	return db
}
