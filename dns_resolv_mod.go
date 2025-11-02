package main

import (
	"fmt"
	"net"
)

func dns_resolv_mod() {
	// resolver := new(net.Resolver)
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// addrs, err := resolver.LookupHost(ctx, "www.google.com")

	addrs, err := net.LookupHost("www.google.com")
	if err != nil {
		if dnsError, isDNSError := err.(*net.DNSError); isDNSError {
			fmt.Printf("[ERROR]: %T, %v\n", err, err)
			fmt.Println(dnsError.IsTimeout, dnsError.IsTemporary, dnsError.Server)
		} else {
			fmt.Printf("[ERROR]: %T, %v\n", err, err)
		}
		return
	}

	fmt.Println(addrs)
}
