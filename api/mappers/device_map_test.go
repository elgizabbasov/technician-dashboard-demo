package mappers

import (
	"testing"
	"time"
)

func TestDevicesGet(t *testing.T) {
	devices, error := DevicesGet()

	if error != nil {
		t.Errorf("DevicesGet() FAILED: %v", error)
	}

	if len(devices) == 0 {
		t.Error("DevicesGet() FAILED: no devices found")
	} else {
		t.Logf("DevicesGet() PASSED: found %d devices", len(devices))
	}
}

func TestDeviceAdd(t *testing.T) {
	// Create a test device
	newDevice := Device{
		ID:           "123",
		Name:         "Test Device",
		Version:      "1.0",
		SiteID:       "456",
		SerialNumber: "789",
		Created:      time.Now().Format("2006-01-02 15:04:05"),
		Updated:      time.Now().Format("2006-01-02 15:04:05"),
	}

	// Add the test device
	success, err := DeviceAdd(newDevice)
	if err != nil {
		t.Errorf("DeviceAdd() FAILED: %v", err)
	}

	// Check that the device was added successfully
	if !success {
		t.Errorf("DeviceAdd() FAILED: device was not added to the database.")
	} else {
		t.Logf("DeviceAdd() PASSED: device was added to the database.")
	}
}
