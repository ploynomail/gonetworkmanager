package main

import (
	"fmt"

	"github.com/ploynomail/gonetworkmanager/v2"
)

func DeleteBridge(name string) {
	// Get the connection by name
	conn, err := gonetworkmanager.GetConnectionByName(name)
	if err != nil {
		fmt.Println(err)
	}
	if conn == nil {
		return
	}

	// Delete the connection
	if err := conn.Delete(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	config := &gonetworkmanager.BondMasterConfig{
		InterfaceName:  "bond0",
		IPAddress:      "172.22.0.22",
		NetMask:        24,
		GatewayAddress: "172.22.0.1",
		IP4Method:      "manual",
		IP6Method:      "ignore",
		AutoConn:       true,
	}
	if err := gonetworkmanager.CreateBondMaster(config); err != nil {
		panic(err)
	}
	slaveConfing8 := &gonetworkmanager.BridgeSlaveConfig{
		InterfaceName: "eth1",
		MasterName:    "bond0",
	}
	if err := gonetworkmanager.CreateBridgeSlave(slaveConfing8); err != nil {
		panic(err)
	}
	// slaveConfing9 := &gonetworkmanager.BridgeSlaveConfig{
	// 	InterfaceName: "enp0s9",
	// 	MasterName:    "bond0",
	// }
	// if err := gonetworkmanager.CreateBridgeSlave(slaveConfing9); err != nil {
	// 	panic(err)
	// }
}
