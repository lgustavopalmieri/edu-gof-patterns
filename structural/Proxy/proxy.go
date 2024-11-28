//Explanation:
//Interface (Internet): Declares a ConnectTo method which both RealInternet and ProxyInternet implement.
//Real Object (RealInternet): Implements unrestricted access, simulating a basic connect functionality.
//Proxy (ProxyInternet): Checks if the requested site is in the blocked list. If so, it denies access. Otherwise, it delegates the connection request to the RealInternet.
//Client Code: Attempts to connect to different websites through the ProxyInternet, which selectively allows or blocks the requests based on a predefined list of blocked sites.

package main

import (
	"fmt"
	"strings"
)

// Internet defines the interface for accessing websites
type Internet interface {
	ConnectTo(site string) string
}

// RealInternet struct that implements the Internet interface, simulating unrestricted access to websites
type RealInternet struct{}

// ConnectTo simulates connecting to a website
func (i *RealInternet) ConnectTo(site string) string {
	return fmt.Sprintf("Connecting to %s", site)
}

// ProxyInternet struct represents the proxy that controls access to websites
type ProxyInternet struct {
	realInternet *RealInternet
	blockedSites []string
}

// NewProxyInternet initializes the ProxyInternet with blocked sites
func NewProxyInternet() *ProxyInternet {
	return &ProxyInternet{
		realInternet: &RealInternet{},
		blockedSites: []string{"blocked.com", "adultsite.com"},
	}
}

// ConnectTo controls access by allowing connections only to non-blocked sites
func (p *ProxyInternet) ConnectTo(site string) string {
	for _, blockedSite := range p.blockedSites {
		if strings.Contains(site, blockedSite) {
			return fmt.Sprintf("Access denied to %s", site)
		}
	}
	return p.realInternet.ConnectTo(site)
}

func main() {
	internet := NewProxyInternet()

	fmt.Println(internet.ConnectTo("golang.org"))         // Should allow
	fmt.Println(internet.ConnectTo("blocked.com"))        // Should deny
	fmt.Println(internet.ConnectTo("educational.com"))    // Should allow
	fmt.Println(internet.ConnectTo("adultsite.com/page")) // Should deny
}
