package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus/v5"
)

const (
	DeviceBondInterface = DeviceInterface + ".Bond"

	/* Properties */
	DeviceBondPropertyHwAddress = DeviceBondInterface + ".HwAddress" // readable   s
)

type DeviceBond interface {
	Device

	// Hardware address of the device.
	GetPropertyHwAddress() (string, error)
}

func NewDeviceBond(objectPath dbus.ObjectPath) (*deviceBond, error) {
	var d deviceBond
	return &d, d.init(NetworkManagerInterface, objectPath)
}

type deviceBond struct {
	device
}

func (d *deviceBond) GetPropertyHwAddress() (string, error) {
	return d.getStringProperty(DeviceBondPropertyHwAddress)
}

func (d *deviceBond) MarshalJSON() ([]byte, error) {
	m, err := d.device.marshalMap()
	if err != nil {
		return nil, err
	}

	m["HwAddress"], _ = d.GetPropertyHwAddress()
	return json.Marshal(m)
}
