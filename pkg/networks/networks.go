package networks

import (
	"errors"
	"net"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/milosgajdos83/tenus"
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

// ComputeIP will return next available ip for the bridge
func ComputeIP(ips []net.IP, network *net.IPNet) (net.IP, *net.IPNet, error) {

	// get mask
	mask := network.Mask
	// compute len
	len, _ := mask.Size()

	for {

		// is ip already assigned ?
		exists := false

		// ask for next subnet of size len
		net, exced := cidr.NextSubnet(network, len)
		if exced != false {
			return nil, nil, errors.New("Can't compute new ip")
		}

		// forge ip, use 1 as host number
		ip, err := cidr.Host(net, 1)
		if err != nil {
			return nil, nil, err
		}

		// check if ip is free
		for _, i := range ips {
			if i.String() == ip.String() {
				exists = true
				network = net
				break
			}
		}

		if !exists {
			return ip, net, err
		}

	}
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

// CreateBridge will create bridge with name and assign ip
func CreateBridge(ip net.IP, net *net.IPNet, name string) error {

	br, err := tenus.NewBridgeWithName(name)
	if err != nil {
		return err
	}

	err = br.SetLinkIp(ip, net)
	if err != nil {
		return err
	}

	err = br.SetLinkUp()

	return err

}
