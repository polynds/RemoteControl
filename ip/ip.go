package ip

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

func GetMac() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v\n", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}

func GetLocalMac() (mac string) {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	for _, inter := range interfaces {
		fmt.Println(inter.Name)
		mac := inter.HardwareAddr //获取本机MAC地址
		fmt.Println("MAC = ", mac)
	}
	return mac
}

func GetIps() (ips []string) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interfaces ipAddress: %v\n", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isVailIpNet := address.(*net.IPNet)
		// 检查ip地址判断是否回环地址
		if isVailIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}

func ClientIp() (string, error) {
	for _, ip := range GetIps() {
		if strings.Contains(ip, "192.168.") {
			return ip, nil
		}
	}

	return "", errors.New("Can not find the client ip address!")
}
