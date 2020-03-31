package ip

import (
	"fmt"
	"net"
	"strings"
)

// ListAllIPv4 list all IPv4 addresses.
func ListAllIPv4(ifaceNames ...string) ([]string, error) {
	list, err := net.Interfaces()

	if err != nil {
		return nil, fmt.Errorf("failed to get interfaces, err: %w", err)
	}

	ips := make([]string, 0)

	ifaceNamesMap := MakeSliceMap(ifaceNames)

	for _, iface := range list {
		if iface.HardwareAddr == nil || iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback == 1 {
			continue
		}

		if len(ifaceNamesMap) > 0 && !ifaceNamesMap[iface.Name] {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			if ipnet.IP.IsLoopback() {
				continue
			}

			if ipv4 := ipnet.IP.To4(); ipv4 != nil {
				ips = append(ips, ipv4.String())
			}
		}
	}

	return ips, nil
}

// GetOutboundIP  gets preferred outbound ip of this machine.
func GetOutboundIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()

	localAddr := conn.LocalAddr().String()

	return localAddr[0:strings.LastIndex(localAddr, ":")]
}

// TryMainIP tries to get the main IP address.
func TryMainIP(ifaceName ...string) string {
	ips, _ := ListAllIPv4(ifaceName...)
	if len(ips) == 1 {
		return ips[0]
	}

	oip := GetOutboundIP()
	if oip == "" {
		return oip
	}

	if len(ips) > 0 {
		return ips[0]
	}

	return ""
}

// MakeSliceMap makes a map[string]bool from the string slice.
func MakeSliceMap(ss []string) map[string]bool {
	m := make(map[string]bool)

	for _, s := range ss {
		if s != "" {
			m[s] = true
		}
	}

	return m
}
