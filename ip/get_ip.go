package main

import "fmt"
import "net"

func main() {
	ifaces, _ := net.Interfaces()
	// handle err
	var ips []string

	for _, i := range ifaces {
		addrs, _ := i.Addrs()

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				fmt.Println("IPNet")
				ip = v.IP
			case *net.IPAddr:
				fmt.Println("IPAddr")
				ip = v.IP
			}
			if ip.String() != "127.0.0.1" {
				ips = append(ips, ip.String())
			}
		}
	}

	fmt.Println("IPs found")
	for _, val := range ips[:] {
		fmt.Printf("%s\n", val)
	}

	if len(ips) == 0 {
		fmt.Println("No IPs found")
	} else {
		if len(ips) > 1 {
			fmt.Println("More than one IPs found")
		} else {
			fmt.Println("Just one IP found")
		}
	}
}
