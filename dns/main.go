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
		fmt.Printf("%s", err)
	} else {
		for _, ip := range ips {
			fmt.Printf("%s. IN A %s\n", domain, ip.String())
		}
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
	} else {
		for _, ip := range ips {
			fmt.Printf("%s. IN A %s\n", domain, ip.String())
		}
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
	nameservers, _ := net.LookupNS(domain)
	for _, ns := range nameservers {
		fmt.Printf("Looking at name server: %s\n", ns.Host)
		Nameserver = ns.Host

		customResolverViaDialerBuiltIn(domain)
		customResolverViaSeparateDialer(domain)
	}
	systemResolver(domain)

	// With the dot on the end, the system will not use the search domains
	customResolverViaSeparateDialer("withdot.microsoft.com.")

	// Without the dot, it will retry the lookups with all my internal
	// search domains appended to it
	customResolverViaSeparateDialer("withoutdot.microsoft.com")

	// With the name server set to one of the Microsoft ones
	// this will fail as they don't do recursive lookups
	systemResolver("testing.zonetransfer.me")

	// Setting the name server to the one for zonetransfer.me, the
	// lookup will work
	Nameserver = "nsztm1.digi.ninja"
	customResolverViaSeparateDialer("testing.zonetransfer.me")

	// This will fail even though it is using the zonetransfer name server
	// for the lookup as the record is a CNAME off to a different server
	// and this name server doesn't do recursive lookups.
	customResolverViaSeparateDialer("staging.zonetransfer.me")
}
