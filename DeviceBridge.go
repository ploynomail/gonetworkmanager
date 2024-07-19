package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus/v5"
)

const (
	DeviceBridgeInterface = DeviceInterface + ".Bridge"

	/* Properties */
	DeviceBridgePropertyHwAddress = DeviceBridgeInterface + ".HwAddress" // readable   s
)

type DeviceBridge interface {
	Device

	// Hardware address of the device.
	GetPropertyHwAddress() (string, error)
}

func NewDeviceBridge(objectPath dbus.ObjectPath) (DeviceBridge, error) {
	var d deviceBridge
	return &d, d.init(NetworkManagerInterface, objectPath)
}

type deviceBridge struct {
	device
}

func (d *deviceBridge) GetPropertyHwAddress() (string, error) {
	return d.getStringProperty(DeviceDummyPropertyHwAddress)
}

func (d *deviceBridge) MarshalJSON() ([]byte, error) {
	m, err := d.device.marshalMap()
	if err != nil {
		return nil, err
	}

	m["HwAddress"], _ = d.GetPropertyHwAddress()
	return json.Marshal(m)
}
