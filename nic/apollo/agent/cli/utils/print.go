package utils

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/pensando/sw/nic/apollo/agent/gen/pds"
)

// MactoStr converts a uint64 to a MAC string
func MactoStr(mac uint64) string {
	var bytes [6]byte

	bytes[0] = byte(mac & 0xFF)
	bytes[1] = byte((mac >> 8) & 0xFF)
	bytes[2] = byte((mac >> 16) & 0xFF)
	bytes[3] = byte((mac >> 24) & 0xFF)
	bytes[4] = byte((mac >> 32) & 0xFF)
	bytes[5] = byte((mac >> 40) & 0xFF)

	macStr := fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", bytes[5], bytes[4], bytes[3], bytes[2], bytes[1], bytes[0])

	return macStr
}

// IPAddrToStr converts HAL proto IP address to string
func IPAddrToStr(ipAddr *pds.IPAddress) string {
	if ipAddr.GetAf() == pds.IPAF_IP_AF_INET {
		v4Addr := ipAddr.GetV4Addr()
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, v4Addr)
		return ip.String()
	} else if ipAddr.GetAf() == pds.IPAF_IP_AF_INET6 {
		v6Addr := ipAddr.GetV6Addr()
		ip := make(net.IP, 16)
		copy(ip, v6Addr)
		return ip.String()
	} else {
		return "-"
	}

}

// IPPrefixToStr converts prefix to string
func IPPrefixToStr(pfx *pds.IPPrefix) string {
	return fmt.Sprintf("%s/%d", IPAddrToStr(pfx.GetAddr()), pfx.GetLen())
}
