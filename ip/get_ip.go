package main

import "fmt"
import "net"

func main() {
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
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
				fmt.Printf("The IP is: %s\n", ip)
			}
		}
	}
}
