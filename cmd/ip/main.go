package main

import (
	"flag"
	"log"
	"net"

	"github.com/bingoohuang/ip"
)

func main() {
	iface := flag.String("iface", "", "Interface name pattern specified(eg. eth0, eth*)")
	verbose := flag.Bool("verbose", false, "Verbose output for more details")
	v4 := flag.Bool("4", false, "only show ipv4")
	v6 := flag.Bool("6", false, "only show ipv6")
	flag.Parse()

	if !*v4 && !*v6 {
		*v4 = true
	}

	_, ipList := ip.MainIP(*iface)
	log.Printf("IP: %v", ipList)
	log.Printf("Outbound IP: %v", ip.Outbound())

	if *v4 {
		allIPv4, _ := ip.ListAllIPv4()
		log.Printf("IPv4: %v", allIPv4)
	}

	if *v6 {
		allIPv6, _ := ip.ListAllIPv6()
		log.Printf("IPv6: %v", allIPv6)
	}

	if *verbose {
		ListIfaces(*v4, *v6)
		moreInfo()
	}
}

// ListIfaces 根据mode 列出本机所有IP和网卡名称.
func ListIfaces(v4, v6 bool) {
	list, err := net.Interfaces()
	if err != nil {
		log.Printf("failed to get interfaces, err: %v", err)
		return
	}

	for _, f := range list {
		listIface(f, v4, v6)
	}
}

func listIface(f net.Interface, v4, v6 bool) {
	log.Printf("iface %+v", f)

	if f.HardwareAddr == nil || f.Flags&net.FlagUp == 0 || f.Flags&net.FlagLoopback == 1 {
		return
	}

	addrs, err := f.Addrs()
	if err != nil {
		log.Printf("\t failed to f.Addrs, × err: %v", err)
		return
	}

	if len(addrs) == 0 {
		return
	}

	got := false
	for _, a := range addrs {
		var netip net.IP
		switch v := a.(type) {
		case *net.IPAddr:
			netip = v.IP
		case *net.IPNet:
			netip = v.IP
		default:
			log.Print("\t\t not .(*net.IPNet) or .(*net.IPNet) ×")
			continue
		}

		if len(netip) == net.IPv4len && !v4 || len(netip) == net.IPv6len && !v6 {
			continue
		}

		if netip.IsLoopback() {
			log.Print("\t\t IsLoopback ×")
			continue
		}

		got = true
	}

	if got {
		log.Printf("\taddrs %+v √", addrs)
	} else {
		log.Printf("\taddrs %+v ×", addrs)
	}
}

func moreInfo() {
	externalIP := ip.External()
	if externalIP == "" {
		return
	}

	log.Printf("公网IP %s", externalIP)
	if eip := net.ParseIP(externalIP); eip != nil {
		result, err := ip.TabaoAPI(externalIP)
		if err != nil {
			log.Printf("TabaoAPI %v", result)
		}
	}

	ipInt := ip.ToDecimal(net.ParseIP(externalIP))
	log.Printf("Convert %s to decimal number(base 10) : %d", externalIP, ipInt)

	ipResult := ip.FromDecimal(ipInt)
	log.Printf("Convert decimal number(base 10) %d to IPv4 address: %v", ipInt, ipResult)

	isBetween := ip.Betweens(net.ParseIP(externalIP), net.ParseIP("0.0.0.0"), net.ParseIP("255.255.255.255"))
	log.Printf("0.0.0.0 isBetween 255.255.255.255 and %s : %v", externalIP, isBetween)

	isPublicIP := ip.IsPublic(net.ParseIP(externalIP))
	log.Printf("%s is public ip: %v ", externalIP, isPublicIP)
}
