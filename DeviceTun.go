package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus/v5"
)

const (
	DeviceTunInterface = DeviceInterface + ".Tun"

	/* Properties */
	DeviceTunPropertyHwAddress = DeviceTunInterface + ".HwAddress" // readable   s
)

type DeviceTun interface {
	Device

	// Hardware address of the device.
	GetPropertyHwAddress() (string, error)
}

func NewDeviceTun(objectPath dbus.ObjectPath) (DeviceTun, error) {
	var d deviceTun
	return &d, d.init(NetworkManagerInterface, objectPath)
}

type deviceTun struct {
	device
}

func (d *deviceTun) GetPropertyHwAddress() (string, error) {
	return d.getStringProperty(DeviceTunPropertyHwAddress)
}

func (d *deviceTun) MarshalJSON() ([]byte, error) {
	m, err := d.device.marshalMap()
	if err != nil {
		return nil, err
	}

	m["HwAddress"], _ = d.GetPropertyHwAddress()
	return json.Marshal(m)
}
