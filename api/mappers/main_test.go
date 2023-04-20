package mappers

import (
	"log"
	"os"
	"testing"

	"github.com/CleanO2tech/clean-technician-web/initializers"
)

func TestMain(m *testing.M) {
	err := initializers.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer initializers.DB().Close()
	os.Exit(m.Run())
	// main.PopulateTables()
}
