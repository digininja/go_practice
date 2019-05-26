package main

import (
	"context"
	"fmt"
	"net"
	"os"
)

var Nameserver = "8.8.8.8"

func GoogleDNSDialer(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{}

	return d.DialContext(ctx, "udp", net.JoinHostPort(Nameserver, "53"))
}

func customResolverViaSeparateDialer(domain string) {
	fmt.Println("customResolverViaSeparateDialer")
	r := net.Resolver{
		PreferGo: true,
		Dial:     GoogleDNSDialer,
	}
	ctx := context.Background()
	ips, err := r.LookupIPAddr(ctx, domain)
	if err != nil {
		panic(err)
	}
	for _, ip := range ips {
		fmt.Printf("%s. IN A %s\n", domain, ip.String())
	}
	fmt.Println()
}

func customResolverViaDialerBuiltIn(domain string) {
	fmt.Println("customResolverViaDialerBuiltIn")

	var resolver *net.Resolver
	if Nameserver != "" {
		resolver = &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{}
				return d.DialContext(ctx, "udp", net.JoinHostPort(Nameserver, "53"))
			},
		}
	} else {
		resolver = net.DefaultResolver
	}

	ips, err := resolver.LookupIPAddr(context.Background(), domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs with custom nameserver: %v\n", err)
		os.Exit(1)
	}
	for _, ip := range ips {
		fmt.Printf("%s. IN A %s\n", domain, ip.String())
	}

	fmt.Println()
}

func systemResolver(domain string) {
	fmt.Println("systemResolver")
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}
	for _, ip := range ips {
		fmt.Printf("%s. IN A %s\n", domain, ip.String())
	}

	fmt.Println()
}

func main() {
	domain := "microsoft.com"

	customResolverViaDialerBuiltIn(domain)
	customResolverViaSeparateDialer(domain)
	systemResolver(domain)
}
