package gonetworkmanager

const (
	ethernetSectionAutoNegotiate = "auto-negotiate"
	connectionSection            = "connection"
	connectionSectionID          = "id"
	connectionSectionType        = "type"
	connectionSectionUUID        = "uuid"
	connectionSectionIfaceName   = "interface-name"
	connectionSectionAutoconnect = "autoconnect"
	ip4Section                   = "ipv4"
	ip4SectionAddressData        = "address-data"
	ip4SectionAddresses          = "addresses"
	ip4SectionAddress            = "address"
	ip4SectionPrefix             = "prefix"
	ip4SectionMethod             = "method"
	ip4SectionGateway            = "gateway"
	ip4SectionNeverDefault       = "never-default"
	ip6Section                   = "ipv6"
	ip6SectionMethod             = "method"
)

// getConnectionByName
func GetConnectionByName(name string) (Connection, error) {
	setting, err := NewSettings()
	if err != nil {
		return nil, err
	}
	conns, err := setting.ListConnections()
	if err != nil {
		return nil, err
	}

	for _, conn := range conns {
		settings, err := conn.GetSettings()
		if err != nil {
			return nil, err
		}

		if settings[connectionSection][connectionSectionID].(string) == name {
			return conn, nil
		}
	}

	return nil, nil
}
