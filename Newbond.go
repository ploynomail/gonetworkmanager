package gonetworkmanager

import (
	"github.com/google/uuid"
)

type BondMasterConfig struct {
	InterfaceName  string
	IPAddress      string
	NetMask        uint32
	GatewayAddress string
	IP4Method      string
	IP6Method      string
	AutoConn       bool
	Mode           string
}

func CreateBondMaster(config *BondMasterConfig) error {
	connectionUUID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	ipaddressNum, err := IpToInt(config.IPAddress)
	if err != nil {
		return err
	}
	gatewayNum, err := IpToInt(config.GatewayAddress)
	if err != nil {
		return err
	}

	ethernetType := "bond"
	ethernetSection := "bond"
	bondSectionMode := "options"

	connection := make(map[string]map[string]interface{})
	connection[ethernetSection] = make(map[string]interface{})
	connection[ethernetSection][bondSectionMode] = make(map[string]string, 0)
	bondOptions, ok := connection[ethernetSection][bondSectionMode].(map[string]string)
	if !ok {
		bondOptions = make(map[string]string)
		connection[ethernetSection][bondSectionMode] = bondOptions
	}
	bondOptions["mode"] = config.Mode
	connection[connectionSection] = make(map[string]interface{})
	connection[connectionSection][connectionSectionID] = config.InterfaceName
	connection[connectionSection][connectionSectionType] = ethernetType

	connection[connectionSection][connectionSectionUUID] = connectionUUID.String()
	connection[connectionSection][connectionSectionIfaceName] = config.InterfaceName
	connection[connectionSection][connectionSectionAutoconnect] = config.AutoConn

	// IPv4
	connection[ip4Section] = make(map[string]interface{})
	// 貌似不起作用
	addressData := make([]map[string]interface{}, 1)
	addressData[0] = make(map[string]interface{})
	addressData[0][ip4SectionAddress] = config.IPAddress
	addressData[0][ip4SectionPrefix] = config.NetMask
	connection[ip4Section][ip4SectionAddressData] = addressData
	// order defined by network manager
	addresses := make([]uint32, 3)
	addresses[0] = ipaddressNum
	addresses[1] = config.NetMask
	addresses[2] = gatewayNum

	addressArray := make([][]uint32, 1)
	addressArray[0] = addresses
	connection[ip4Section][ip4SectionAddresses] = addressArray

	connection[ip4Section][ip4SectionGateway] = config.GatewayAddress
	connection[ip4Section][ip4SectionMethod] = config.IP4Method
	connection[ip4Section][ip4SectionNeverDefault] = true

	// IPv6 igrnoe
	connection[ip6Section] = make(map[string]interface{})
	connection[ip6Section][ip6SectionMethod] = config.IP6Method

	settings, err := NewSettings()

	if err != nil {
		return err
	}

	_, err = settings.AddConnection(connection)

	if err != nil {
		return err
	}
	return nil
}

type BondSlaveConfig struct {
	InterfaceName string
	MasterName    string
}

func CreateBondSlave(config *BondSlaveConfig) error {
	connectionUUID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	ethernetType := "802-3-ethernet"
	connectionMaster := "master"
	connectionSlaveTypeSection := "slave-type"
	connectionSlaveType := "bond"

	connection := make(map[string]map[string]interface{})
	connection[connectionSection] = make(map[string]interface{})
	connection[connectionSection][connectionSectionID] = config.MasterName + "-" + config.InterfaceName
	connection[connectionSection][connectionSectionType] = ethernetType

	connection[connectionSection][connectionSectionUUID] = connectionUUID.String()
	connection[connectionSection][connectionSectionIfaceName] = config.InterfaceName
	connection[connectionSection][connectionSectionAutoconnect] = true
	connection[connectionSection][connectionMaster] = config.MasterName
	connection[connectionSection][connectionSlaveTypeSection] = connectionSlaveType

	settings, err := NewSettings()

	if err != nil {
		return err
	}

	_, err = settings.AddConnection(connection)

	if err != nil {
		return err
	}
	return nil
}
