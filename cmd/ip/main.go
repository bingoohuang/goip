package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/bingoohuang/ip"
	"github.com/sirupsen/logrus"
)

func main() {
	iface := flag.String("iface", "", "Interface name pattern specified(eg. eth0, eth*)")
	flag.Parse()

	mainIP, ipList := ip.TryMainIP(*iface)
	logrus.Infof("TryMainIP: %s", mainIP)
	logrus.Infof("IP List: %v", ipList)

	ListIfaces()
	ListenAddr()
	GetOutboundIP()
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

		if iface.HardwareAddr == nil {
			logrus.Infof("\t iface.HardwareAddr == nil ×")
			continue
		}

		if iface.Flags&net.FlagUp == 0 {
			logrus.Infof("\t net.FlagUp = 0 ×")
			continue
		}

		if iface.Flags&net.FlagLoopback == 1 {
			logrus.Infof("\t net.FlagLoopback = 0 ×")
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			logrus.Warnf("\t failed to iface.Addrs, × err: %v", err)
			continue
		}

		logrus.Infof("\taddrs %+v", addrs)

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

			logrus.Infof("\t\t√ Got")
		}
	}
}

// ListenAddr returns the listenable TCP addresss.
func ListenAddr() {
	l, e := ListenPort(0)
	if e != nil {
		logrus.Warnf("failed to ListenPort(0), × err: %v", e)
		return
	}

	defer l.Close()

	addr := l.Addr().(*net.TCPAddr)
	logrus.Infof("ListenAddrIP %+v", addr)
}

// ListenPort listens on port
func ListenPort(port int) (net.Listener, error) {
	return net.Listen("tcp", fmt.Sprintf(":%d", port))
}

// GetOutboundIP gets preferred outbound ip of this machine.
func GetOutboundIP() {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		logrus.Warnf("failed to udp 8.8.8.8:80, × err: %v", err)
		return
	}

	defer conn.Close()

	addr := conn.LocalAddr().(*net.UDPAddr)
	logrus.Infof("OutboundIP %+v", addr)
}

// nolint lll
// https://topic.alibabacloud.com/a/go-combat-golang-get-public-ip-view-intranet-ip-detect-ip-type-verify-ip-range-ip-address-string-and-int-conversion-judge-by-ip_1_38_10267608.html

// IPInfo ...
type IPInfo struct {
	Code int `json:"code"`
	Data IP  `json:"data"`
}

// IP ...
type IP struct {
	Country   string `json:"country"`
	CountryID string `json:"country_id"`
	Area      string `json:"area"`
	AreaID    string `json:"area_id"`
	Region    string `json:"region"`
	RegionID  string `json:"region_id"`
	City      string `json:"city"`
	CityID    string `json:"city_id"`
	Isp       string `json:"isp"`
}

func moreInfo() {
	externalIP := getExternal()
	externalIP = strings.Replace(externalIP, "\n", "", -1)
	logrus.Infof("公网IP %s", externalIP)

	eip := net.ParseIP(externalIP)
	if eip != nil {
		result := TabaoAPI(externalIP)
		if result != nil {
			logrus.Infof("TabaoAPI %+v", result)
		}
	}

	ipInt := inetAton(net.ParseIP(externalIP))
	logrus.Infof("Convert %s to decimal number(base 10) : %d", externalIP, ipInt)

	ipResult := inetNtoa(ipInt)
	logrus.Infof("Convert decimal number(base 10) %d to IPv4 address: %v", ipInt, ipResult)

	isBetween := IPBetween(net.ParseIP("0.0.0.0"), net.ParseIP("255.255.255.255"), net.ParseIP(externalIP))
	logrus.Infof("0.0.0.0 isBetween 255.255.255.255 and %s : %v", externalIP, isBetween)

	isPublicIP := IsPublicIP(net.ParseIP(externalIP))
	logrus.Infof("%s is public ip: %v ", externalIP, isPublicIP)

	logrus.Infof("PulicIP:%s", ip.GetOutboundIP())
}

func getExternal() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(resp.Body)
	//s := buf.String()
	return string(content)
}

// TabaoAPI ...
func TabaoAPI(ip string) *IPInfo {
	url := "http://ip.taobao.com/service/getIpInfo.php?ip=" + ip

	resp, err := http.Get(url) // nolint gosec
	if err != nil {
		logrus.Warnf("failed http.Get(%s), × err: %v", url, err)

		return nil
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Warnf("failed ioutil.ReadAll, × err: %v", err)
		return nil
	}

	var result IPInfo

	if err := json.Unmarshal(out, &result); err != nil {
		logrus.Warnf("failed json.Unmarshal %s, × err: %v", string(out), err)
		return nil
	}

	return &result
}

// nolint gomnd
func inetNtoa(ipnr int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

// nolint gomnd
func inetAton(ipnr net.IP) int64 {
	bits := strings.Split(ipnr.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

// IPBetween ...
func IPBetween(from net.IP, to net.IP, test net.IP) bool {
	if from == nil || to == nil || test == nil {
		logrus.Warnf("An IP input is nil")
		return false
	}

	from16 := from.To16()
	to16 := to.To16()
	test16 := test.To16()

	if from16 == nil || to16 == nil || test16 == nil {
		logrus.Warnf("An IP did not convert to a 16 byte")
		return false
	}

	return bytes.Compare(test16, from16) >= 0 && bytes.Compare(test16, to16) <= 0
}

// IsPublicIP ...
// nolint gomnd
func IsPublicIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := ip.To4(); ip4 != nil {
		switch {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}

	return false
}
