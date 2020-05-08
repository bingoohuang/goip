package main

import (
	"flag"
	"net"

	"github.com/bingoohuang/ip"
	"github.com/sirupsen/logrus"
)

func main() {
	iface := flag.String("iface", "", "Interface name pattern specified(eg. eth0, eth*)")
	flag.Parse()

	mainIP, ipList := ip.MainIP(*iface)
	logrus.Infof("MainIP: %s, IP List: %v", mainIP, ipList)
	logrus.Infof("OutboundIP: %v", ip.Outbound())

	allIPv4, _ := ip.ListAllIPv4()
	logrus.Infof("IPv4: %v", allIPv4)

	allIPv6, _ := ip.ListAllIPv6()
	logrus.Infof("IPv6: %v", allIPv6)

	ListIfaces()
	moreInfo()
}

// ListIfaces 根据mode 列出本机所有IP和网卡名称
func ListIfaces() {
	list, err := net.Interfaces()
	if err != nil {
		logrus.Warnf("failed to get interfaces, err: %v", err)
		return
	}

	for _, iface := range list {
		logrus.Infof("iface %+v", iface)

		if iface.HardwareAddr == nil || iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback == 1 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			logrus.Warnf("\t failed to iface.Addrs, × err: %v", err)
			continue
		}

		if len(addrs) == 0 {
			continue
		}

		got := false

		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok {
				logrus.Infof("\t\t not addr.(*net.IPNet) ×")
				continue
			}

			if ipnet.IP.IsLoopback() {
				logrus.Infof("\t\t IsLoopback ignored")
				continue
			}

			got = true
		}

		if got {
			logrus.Infof("\taddrs %+v √ Got", addrs)
		} else {
			logrus.Infof("\taddrs %+v × Failed", addrs)
		}
	}
}

func moreInfo() {
	externalIP := ip.External()
	logrus.Infof("公网IP %s", externalIP)

	if externalIP == "" {
		return
	}

	eip := net.ParseIP(externalIP)
	if eip != nil {
		result := ip.TabaoAPI(externalIP)
		if result != nil {
			logrus.Infof("TabaoAPI %+v", result)
		}
	}

	ipInt := ip.ToDecimal(net.ParseIP(externalIP))
	logrus.Infof("Convert %s to decimal number(base 10) : %d", externalIP, ipInt)

	ipResult := ip.FromDecimal(ipInt)
	logrus.Infof("Convert decimal number(base 10) %d to IPv4 address: %v", ipInt, ipResult)

	isBetween := ip.Betweens(net.ParseIP(externalIP), net.ParseIP("0.0.0.0"), net.ParseIP("255.255.255.255"))
	logrus.Infof("0.0.0.0 isBetween 255.255.255.255 and %s : %v", externalIP, isBetween)

	isPublicIP := ip.IsPublic(net.ParseIP(externalIP))
	logrus.Infof("%s is public ip: %v ", externalIP, isPublicIP)
}
