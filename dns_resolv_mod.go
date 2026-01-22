package main

import (
	"errors"
	"fmt"
	"net"
)

func dns_resolv_mod() {
	// resolver := new(net.Resolver)
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// addrs, err := resolver.LookupHost(ctx, "www.google.com")

	addrs, err := net.LookupHost("www.google.com")
	if err != nil {
		var dnsError *net.DNSError
		// if dnsError, isDNSError := err.(*net.DNSError); isDNSError {
		// if errors.Is(err, &net.DNSError{}) {
		if errors.As(err, &dnsError) {
			fmt.Printf("[ERROR]: %T, %v\n", err, err)
			fmt.Println(dnsError.IsTimeout, dnsError.IsTemporary, dnsError.Server)
		} else {
			fmt.Printf("[ERROR]: %T, %v\n", err, err)
		}
		return
	}

	fmt.Println(addrs)
}
