# showip
show host IP addresses


```bash
$ ip                                                                                                                                            [二  3/31 11:48:19 2020]
INFO[0000] TryMainIP:192.168.162.17
INFO[0000] go iface {Index:1 MTU:16384 Name:lo0 HardwareAddr: Flags:up|loopback|multicast}
INFO[0000] 	 iface.HardwareAddr == nil ×
INFO[0000] go iface {Index:2 MTU:1280 Name:gif0 HardwareAddr: Flags:pointtopoint|multicast}
INFO[0000] 	 iface.HardwareAddr == nil ×
INFO[0000] go iface {Index:3 MTU:1280 Name:stf0 HardwareAddr: Flags:0}
INFO[0000] 	 iface.HardwareAddr == nil ×
INFO[0000] go iface {Index:4 MTU:0 Name:XHC1 HardwareAddr: Flags:0}
INFO[0000] 	 iface.HardwareAddr == nil ×
INFO[0000] go iface {Index:5 MTU:0 Name:XHC0 HardwareAddr: Flags:0}
INFO[0000] 	 iface.HardwareAddr == nil ×
INFO[0000] go iface {Index:6 MTU:0 Name:XHC20 HardwareAddr: Flags:0}
INFO[0000] 	 iface.HardwareAddr == nil ×
INFO[0000] go iface {Index:7 MTU:0 Name:VHC128 HardwareAddr: Flags:0}
INFO[0000] 	 iface.HardwareAddr == nil ×
INFO[0000] go iface {Index:8 MTU:1500 Name:en5 HardwareAddr:ac:de:48:00:11:22 Flags:up|broadcast|multicast}
INFO[0000] 	addrs [fe80::aede:48ff:fe00:1122/64]
INFO[0000] 		√ Got
INFO[0000] go iface {Index:9 MTU:1500 Name:bridge0 HardwareAddr:82:68:2b:61:34:01 Flags:up|broadcast|multicast}
INFO[0000] 	addrs []
INFO[0000] go iface {Index:10 MTU:1500 Name:ap1 HardwareAddr:f2:18:98:a5:12:27 Flags:broadcast|multicast}
INFO[0000] 	 net.FlagUp = 0 ×
INFO[0000] go iface {Index:11 MTU:1500 Name:en0 HardwareAddr:f0:18:98:a5:12:27 Flags:up|broadcast|multicast}
INFO[0000] 	addrs [fe80::d2:7528:7170:8891/64 192.168.162.17/24]
INFO[0000] 		√ Got
INFO[0000] 		√ Got
INFO[0000] go iface {Index:12 MTU:1500 Name:en1 HardwareAddr:82:68:2b:61:34:01 Flags:up|broadcast|multicast}
INFO[0000] 	addrs []
INFO[0000] go iface {Index:13 MTU:1500 Name:en2 HardwareAddr:82:68:2b:61:34:00 Flags:up|broadcast|multicast}
INFO[0000] 	addrs []
INFO[0000] go iface {Index:14 MTU:1500 Name:en3 HardwareAddr:82:68:2b:61:34:05 Flags:up|broadcast|multicast}
INFO[0000] 	addrs []
INFO[0000] go iface {Index:15 MTU:1500 Name:en4 HardwareAddr:82:68:2b:61:34:04 Flags:up|broadcast|multicast}
INFO[0000] 	addrs []
INFO[0000] go iface {Index:16 MTU:2304 Name:p2p0 HardwareAddr:02:18:98:a5:12:27 Flags:up|broadcast|multicast}
INFO[0000] 	addrs []
INFO[0000] go iface {Index:17 MTU:1484 Name:awdl0 HardwareAddr:42:88:97:6b:60:32 Flags:up|broadcast|multicast}
INFO[0000] 	addrs [fe80::4088:97ff:fe6b:6032/64]
INFO[0000] 		√ Got
INFO[0000] go iface {Index:18 MTU:1500 Name:llw0 HardwareAddr:42:88:97:6b:60:32 Flags:up|broadcast|multicast}
INFO[0000] 	addrs [fe80::4088:97ff:fe6b:6032/64]
INFO[0000] 		√ Got
INFO[0000] go iface {Index:19 MTU:1380 Name:utun0 HardwareAddr: Flags:up|pointtopoint|multicast}
INFO[0000] 	 iface.HardwareAddr == nil ×
INFO[0000] go iface {Index:20 MTU:2000 Name:utun1 HardwareAddr: Flags:up|pointtopoint|multicast}
INFO[0000] 	 iface.HardwareAddr == nil ×
INFO[0000] ListenAddrIP [::]:52463
INFO[0000] OutboundIP 192.168.162.17:55587
INFO[0022] 公网IP 60.247.93.190
INFO[0022] TabaoAPI &{Code:0 Data:{Country:中国 CountryID:CN Area: AreaID: Region:北京 RegionID:110000 City:北京 CityID:110100 Isp:电信}}
INFO[0022] Convert 60.247.93.190 to decimal number(base 10) : 1022844350
INFO[0022] Convert decimal number(base 10) 1022844350 to IPv4 address: 60.247.93.190
INFO[0022] 0.0.0.0 isBetween 255.255.255.255 and 60.247.93.190 : true
INFO[0022] 60.247.93.190 is public ip: true
INFO[0022] PulicIP:192.168.162.17
```
