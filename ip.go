package ip

import (
	"fmt"
	"net"
	"path/filepath"
	"strings"
)

// ListAllIPv4 list all IPv4 addresses.
// ifaceNames are used to specified interface names (filename wild match pattern supported also, like eth*)
func ListAllIPv4(ifaceNames ...string) ([]string, error) {
	ips := make([]string, 0)

	_, err := ListAllIP(func(ip net.IP) bool {
		yes := len(ip.To4()) == net.IPv4len
		if yes {
			ips = append(ips, ip.To4().String())
		}

		return yes
	}, ifaceNames...)

	return ips, err
}

// ListAllIPv6 list all IPv6 addresses.
// ifaceNames are used to specified interface names (filename wild match pattern supported also, like eth*)
func ListAllIPv6(ifaceNames ...string) ([]string, error) {
	ips := make([]string, 0)

	_, err := ListAllIP(func(ip net.IP) bool {
		yes := len(ip.To4()) != net.IPv4len
		if yes {
			ips = append(ips, ip.To16().String())
		}

		return yes
	}, ifaceNames...)

	return ips, err
}

// ListAllIP list all IP addresses.
func ListAllIP(predicate func(net.IP) bool, ifaceNames ...string) ([]net.IP, error) {
	list, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get interfaces, err: %w", err)
	}

	ips := make([]net.IP, 0)

	matcher := newIfaceNameMatcher(ifaceNames)

	for _, i := range list {
		f := i.Flags
		if i.HardwareAddr == nil || f&net.FlagUp == 0 || f&net.FlagLoopback == 1 || !matcher.Matches(i.Name) {
			continue
		}

		addrs, err := i.Addrs()
		if err != nil {
			continue
		}

		for _, a := range addrs {
			var ip net.IP
			switch v := a.(type) {
			case *net.IPAddr:
				ip = v.IP
			case *net.IPNet:
				ip = v.IP
			default:
				continue
			}

			if !containsIp(ips, ip) && predicate(ip) {
				ips = append(ips, ip)
			}
		}
	}

	return ips, nil
}

func containsIp(ips []net.IP, ip net.IP) bool {
	for _, j := range ips {
		if j.Equal(ip) {
			return true
		}
	}

	return false
}

// Outbound  gets preferred outbound ip of this machine.
func Outbound() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}

	defer conn.Close()

	s := conn.LocalAddr().String()
	return s[:strings.LastIndex(s, ":")]
}

// MainIP tries to get the main IP address and the IP addresses.
func MainIP(ifaceName ...string) (string, []string) {
	ips, _ := ListAllIPv4(ifaceName...)
	if len(ips) == 1 {
		return ips[0], ips
	}

	if oip := Outbound(); oip != "" && contains(ips, oip) {
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
