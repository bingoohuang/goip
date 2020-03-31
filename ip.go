package ip

import (
	"fmt"
	"net"
	"path/filepath"
	"strings"
)

// ListAllIPv4 list all IPv4 addresses.
func ListAllIPv4(ifaceNames ...string) ([]string, error) {
	list, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get interfaces, err: %w", err)
	}

	ips := make([]string, 0)

	matcher := newIfaceNameMatcher(ifaceNames)

	for _, iface := range list {
		if iface.HardwareAddr == nil || iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback == 1 {
			continue
		}

		if !matcher.Matches(iface.Name) {
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

// TryMainIP tries to get the main IP address and the IP addresses.
func TryMainIP(ifaceName ...string) (string, []string) {
	ips, _ := ListAllIPv4(ifaceName...)
	if len(ips) == 1 {
		return ips[0], ips
	}

	if oip := GetOutboundIP(); oip != "" && contains(ips, oip) {
		return oip, ips
	}

	if len(ips) > 0 {
		return ips[0], ips
	}

	return "", ips
}

func contains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}

	return false
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

type ifaceNameMatcher struct {
	ifacePatterns map[string]bool
}

func newIfaceNameMatcher(ss []string) ifaceNameMatcher {
	return ifaceNameMatcher{ifacePatterns: MakeSliceMap(ss)}
}

func (i ifaceNameMatcher) Matches(name string) bool {
	if len(i.ifacePatterns) == 0 {
		return true
	}

	if _, ok := i.ifacePatterns[name]; ok {
		return true
	}

	for k := range i.ifacePatterns {
		if ok, _ := filepath.Match(k, name); ok {
			return true
		}
	}

	return false
}
