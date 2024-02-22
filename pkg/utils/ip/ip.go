package ip

import (
	"net"
)

func GetIPv4NetworkIPs() ([]string, error) {
	ips, err := GetNetworkIPs()
	if err != nil {
		return nil, err
	}

	ip4s := make([]string, 0)
	for _, ip := range ips {
		if net.ParseIP(ip).To4() != nil {
			ip4s = append(ip4s, ip)
		}
	}

	return ip4s, nil
}

func GetNetworkIPs() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	ips := make([]string, 0)
	for _, i := range ifaces {
		if i.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ips = append(ips, v.IP.String())
			case *net.IPAddr:
				ips = append(ips, v.IP.String())
			}
		}
	}

	return ips, nil
}
