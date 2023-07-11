// https://go.dev/tour/methods/18
package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var s string
	for n, cel := range ip {
		if n < 3 {
			s += fmt.Sprintf("%v.", cel)
		} else {
			s += fmt.Sprintf("%v", cel)
		}
	}
	return s

}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
