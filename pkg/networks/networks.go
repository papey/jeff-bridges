package networks

import (
	"net"
)

// GetAll function gets all ips of all interfaces on the host
func GetAll() ([]net.IP, error) {

	// get all ifaces
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	ips, err := extractIfaces(ifaces)
	if err != nil {
		return nil, err
	}

	return ips, err

}

// https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
func extractIfaces(ifaces []net.Interface) ([]net.IP, error) {

	var ips []net.IP

	// loop on all interfaces
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		// loop
		for _, addr := range addrs {
			// ensure variable
			var ip net.IP

			// Switch between possible values
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// if it's an IPv4, add to slice
			if ip.To4() != nil {
				ips = append(ips, ip)
			}

		}
	}

	return ips, nil

}
