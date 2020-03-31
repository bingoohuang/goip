# showip

show host IP addresses

## API

```go
import "github.com/bingoohuang/ip"

// ListAllIPv4 list all IPv4 addresses.
// The input argument ifaceNames are used to specified interface names (filename wild match pattern supported also, like eth*)
allIPv4s, _ := ip.ListAllIPv4()
allEth0IPv4s, _ := ip.ListAllIPv4("eth0")
allEth0En0IPv4s, _ := ip.ListAllIPv4("eth0", "en0")
allEnIPv4s, _ := ip.ListAllIPv4("en*")

// GetOutboundIP  gets preferred outbound ip of this machine.
outboundIP := ip.GetOutboundIP()

// TryMainIP tries to get the main IP address and the IP addresses.
mainIP, ipList := ip.TryMainIP()
```

## Usages

on mac:

```bash
./ip                                                                                                                                                        [二  3/31 13:26:11 2020]
INFO[0000] TryMainIP: 192.168.162.17
INFO[0000] IP List: [192.168.162.17]
INFO[0000] iface {Index:1 MTU:16384 Name:lo0 HardwareAddr: Flags:up|loopback|multicast}
INFO[0000]       iface.HardwareAddr == nil ×
INFO[0000] iface {Index:2 MTU:1280 Name:gif0 HardwareAddr: Flags:pointtopoint|multicast}
INFO[0000]       iface.HardwareAddr == nil ×
INFO[0000] iface {Index:3 MTU:1280 Name:stf0 HardwareAddr: Flags:0}
INFO[0000]       iface.HardwareAddr == nil ×
INFO[0000] iface {Index:4 MTU:0 Name:XHC1 HardwareAddr: Flags:0}
INFO[0000]       iface.HardwareAddr == nil ×
INFO[0000] iface {Index:5 MTU:0 Name:XHC0 HardwareAddr: Flags:0}
INFO[0000]       iface.HardwareAddr == nil ×
INFO[0000] iface {Index:6 MTU:0 Name:XHC20 HardwareAddr: Flags:0}
INFO[0000]       iface.HardwareAddr == nil ×
INFO[0000] iface {Index:7 MTU:0 Name:VHC128 HardwareAddr: Flags:0}
INFO[0000]       iface.HardwareAddr == nil ×
INFO[0000] iface {Index:8 MTU:1500 Name:en5 HardwareAddr:ac:de:48:00:11:22 Flags:up|broadcast|multicast}
INFO[0000]      addrs [fe80::aede:48ff:fe00:1122/64]
INFO[0000]              √ Got
INFO[0000] iface {Index:9 MTU:1500 Name:bridge0 HardwareAddr:82:68:2b:61:34:01 Flags:up|broadcast|multicast}
INFO[0000]      addrs []
INFO[0000] iface {Index:10 MTU:1500 Name:ap1 HardwareAddr:f2:18:98:a5:12:27 Flags:broadcast|multicast}
INFO[0000]       net.FlagUp = 0 ×
INFO[0000] iface {Index:11 MTU:1500 Name:en0 HardwareAddr:f0:18:98:a5:12:27 Flags:up|broadcast|multicast}
INFO[0000]      addrs [fe80::d2:7528:7170:8891/64 192.168.162.17/24]
INFO[0000]              √ Got
INFO[0000]              √ Got
INFO[0000] iface {Index:12 MTU:1500 Name:en1 HardwareAddr:82:68:2b:61:34:01 Flags:up|broadcast|multicast}
INFO[0000]      addrs []
INFO[0000] iface {Index:13 MTU:1500 Name:en2 HardwareAddr:82:68:2b:61:34:00 Flags:up|broadcast|multicast}
INFO[0000]      addrs []
INFO[0000] iface {Index:14 MTU:1500 Name:en3 HardwareAddr:82:68:2b:61:34:05 Flags:up|broadcast|multicast}
INFO[0000]      addrs []
INFO[0000] iface {Index:15 MTU:1500 Name:en4 HardwareAddr:82:68:2b:61:34:04 Flags:up|broadcast|multicast}
INFO[0000]      addrs []
INFO[0000] iface {Index:16 MTU:2304 Name:p2p0 HardwareAddr:02:18:98:a5:12:27 Flags:up|broadcast|multicast}
INFO[0000]      addrs []
INFO[0000] iface {Index:17 MTU:1484 Name:awdl0 HardwareAddr:42:88:97:6b:60:32 Flags:up|broadcast|multicast}
INFO[0000]      addrs [fe80::4088:97ff:fe6b:6032/64]
INFO[0000]              √ Got
INFO[0000] iface {Index:18 MTU:1500 Name:llw0 HardwareAddr:42:88:97:6b:60:32 Flags:up|broadcast|multicast}
INFO[0000]      addrs [fe80::4088:97ff:fe6b:6032/64]
INFO[0000]              √ Got
INFO[0000] iface {Index:19 MTU:1380 Name:utun0 HardwareAddr: Flags:up|pointtopoint|multicast}
INFO[0000]       iface.HardwareAddr == nil ×
INFO[0000] iface {Index:20 MTU:2000 Name:utun1 HardwareAddr: Flags:up|pointtopoint|multicast}
INFO[0000]       iface.HardwareAddr == nil ×
INFO[0000] ListenAddrIP [::]:62061
INFO[0000] OutboundIP 192.168.162.17:49360
INFO[0007] 公网IP 60.247.93.190
INFO[0007] TabaoAPI &{Code:0 Data:{Country:中国 CountryID:CN Area: AreaID: Region:北京 RegionID:110000 City:北京 CityID:110100 Isp:电信}}
INFO[0007] Convert 60.247.93.190 to decimal number(base 10) : 1022844350
INFO[0007] Convert decimal number(base 10) 1022844350 to IPv4 address: 60.247.93.190
INFO[0007] 0.0.0.0 isBetween 255.255.255.255 and 60.247.93.190 : true
INFO[0007] 60.247.93.190 is public ip: true
INFO[0007] PulicIP:192.168.162.17
```

on linux:

```bash
$ ./ip-v1.0.0-amd64-glibc2.28 -iface docker0
INFO[0000] TryMainIP: 172.17.0.1
INFO[0000] IP List: [172.17.0.1]
INFO[0000] go iface {Index:1 MTU:65536 Name:lo HardwareAddr: Flags:up|loopback}
INFO[0000]       iface.HardwareAddr == nil ×
INFO[0000] go iface {Index:2 MTU:1500 Name:eth0 HardwareAddr:52:54:00:ef:16:bd Flags:up|broadcast|multicast}
INFO[0000]      addrs [192.168.1.17/24 192.168.1.7/32]
INFO[0000]              √ Got
INFO[0000]              √ Got
INFO[0000] go iface {Index:3 MTU:1500 Name:br-8983f91a1c88 HardwareAddr:02:42:49:a5:88:9f Flags:up|broadcast|multicast}
INFO[0000]      addrs [172.18.0.1/16]
INFO[0000]              √ Got
INFO[0000] go iface {Index:4 MTU:1500 Name:docker0 HardwareAddr:02:42:b3:f8:ea:89 Flags:up|broadcast|multicast}
INFO[0000]      addrs [172.17.0.1/16]
INFO[0000]              √ Got
INFO[0000] go iface {Index:7 MTU:1500 Name:br-d4979d31f397 HardwareAddr:02:42:65:78:5e:15 Flags:up|broadcast|multicast}
INFO[0000]      addrs [172.19.0.1/16]
INFO[0000]              √ Got
INFO[0000] go iface {Index:33 MTU:1500 Name:vethb98680d HardwareAddr:ce:8b:3c:1c:c1:9a Flags:up|broadcast|multicast}
INFO[0000]      addrs []
INFO[0000] go iface {Index:37 MTU:1500 Name:veth69f8b3d HardwareAddr:3e:1d:13:80:f0:a7 Flags:up|broadcast|multicast}
INFO[0000]      addrs []
INFO[0000] go iface {Index:71 MTU:1500 Name:veth0d4f771 HardwareAddr:56:92:a6:2c:6f:75 Flags:up|broadcast|multicast}
INFO[0000]      addrs []
INFO[0000] ListenAddrIP [::]:38328
INFO[0000] OutboundIP 192.168.1.17:34517
INFO[0000] 公网IP 123.206.185.162
INFO[0000] TabaoAPI &{Code:0 Data:{Country:中国 CountryID:CN Area: AreaID: Region:上海 RegionID:310000 City:上海 CityID:310100 Isp:电信}}
INFO[0000] Convert 123.206.185.162 to decimal number(base 10) : 2077145506
INFO[0000] Convert decimal number(base 10) 2077145506 to IPv4 address: 123.206.185.162
INFO[0000] 0.0.0.0 isBetween 255.255.255.255 and 123.206.185.162 : true
INFO[0000] 123.206.185.162 is public ip: true
INFO[0000] PulicIP:192.168.1.17
```
