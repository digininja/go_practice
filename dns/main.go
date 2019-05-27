package main

/*
http://www.golangprograms.com/find-dns-records-programmatically.html

https://golang.org/pkg/net/
*/

import (
	"context"
	"fmt"
	"net"
	"os"
)

var Nameserver = "8.8.8.8"

func CustomDialer(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{}

	return d.DialContext(ctx, "udp", net.JoinHostPort(Nameserver, "53"))
}

func customResolverViaSeparateDialer(domain string) {
	fmt.Println("customResolverViaSeparateDialer")
	r := net.Resolver{
		PreferGo: true,
		Dial:     CustomDialer,
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

func systemResolverCNAME(domain string) {
	fmt.Println("systemResolver")
	ips, err := net.LookupCNAME(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}
	for _, ip := range ips {
		fmt.Printf("%s. IN A %s\n", domain, ip)
	}

	fmt.Println()
}

func main() {
	domain := "microsoft.com"

	nameserver, _ := net.LookupNS(domain)
	for _, ns := range nameserver {
		fmt.Printf("Looking at name server: %s\n", ns.Host)
		Nameserver = ns.Host

		customResolverViaDialerBuiltIn(domain)
		customResolverViaSeparateDialer(domain)
	}
	systemResolver(domain)
	systemResolver("staging.zonetransfer.me")
	Nameserver = "nsztm1.digi.ninja"
	customResolverViaSeparateDialer("staging.zonetransfer.me")
}
