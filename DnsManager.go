package gonetworkmanager

import (
	"errors"

	"github.com/godbus/dbus/v5"
)

const (
	DnsManagerInterface  = NetworkManagerInterface + ".DnsManager"
	DnsManagerObjectPath = "/org/freedesktop/NetworkManager/DnsManager"

	/* Property */
	DnsManagerPropertyMode          = DnsManagerInterface + ".Mode"          // readable s
	DnsManagerPropertyRcManager     = DnsManagerInterface + ".RcManager"     // readable s
	DnsManagerPropertyConfiguration = DnsManagerInterface + ".Configuration" // readable aa{sv}
)

type DnsConfigurationData struct {
	Nameservers []string
	Priority    int32
	Interface   string
	Vpn         bool
}

type DnsManager interface {
	GetPath() dbus.ObjectPath
	GetPropertyConfiguration() ([]DnsConfigurationData, error)
	GetPropertyRcManager() (string, error)
	GetPropertyMode() (string, error)
}

type dnsManager struct {
	dbusBase
}

func NewDnsManager() (DnsManager, error) {
	var d dnsManager
	return &d, d.init(NetworkManagerInterface, DnsManagerObjectPath)
}

func (d *dnsManager) GetPath() dbus.ObjectPath {
	return d.obj.Path()
}

func (d *dnsManager) GetPropertyConfiguration() ([]DnsConfigurationData, error) {
	configurations, err := d.getSliceMapStringVariantProperty(DnsManagerPropertyConfiguration)
	if err != nil {
		return nil, err
	}

	// for k, c := range configurations {
	// 	fmt.Println(k, c)
	// }

	ret := make([]DnsConfigurationData, len(configurations))
	for i, conf := range configurations {
		if serversVar, exist := conf["nameservers"]; exist {
			servers, ok := serversVar.Value().([]string)
			if !ok {
				return nil, errors.New("unexpected variant type for nameservers")
			}
			ret[i].Nameservers = servers
		}

		if priorityVar, exist := conf["priority"]; exist {
			priority, ok := priorityVar.Value().(int32)
			if !ok {
				return nil, errors.New("unexpected variant type for priority")
			}
			ret[i].Priority = priority
		}

		if interfaceVar, exist := conf["interface"]; exist {
			iface, ok := interfaceVar.Value().(string)
			if !ok {
				return nil, errors.New("unexpected variant type for interface")
			}
			ret[i].Interface = iface
		}

		if vpnVar, exist := conf["vpn"]; exist {
			vpn, ok := vpnVar.Value().(bool)
			if !ok {
				return nil, errors.New("unexpected variant type for vpn")
			}
			ret[i].Vpn = vpn
		}
	}

	return ret, nil
}

func (d *dnsManager) GetPropertyRcManager() (string, error) {
	return d.getStringProperty(DnsManagerPropertyRcManager)
}

func (d *dnsManager) GetPropertyMode() (string, error) {
	return d.getStringProperty(DnsManagerPropertyMode)
}
