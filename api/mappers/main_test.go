package mappers

import (
	"log"
	"os"
	"testing"

	"github.com/elgizabbasov/technician-dashboard-demo/initializers"
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
