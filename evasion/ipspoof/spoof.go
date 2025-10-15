package ipspoof

import (
	"crypto/rand"
	"net"
)

func GenerateByCountry(country string) string {
	if country == "US" {
		// Use AWS US-East block: 3.0.0.0/8
		b := make([]byte, 3)
		rand.Read(b)
		return net.IPv4(3, b[0], b[1], b[2]).String()
	}
	ip := make([]byte, 4)
	rand.Read(ip)
	return net.IPv4(ip[0], ip[1], ip[2], ip[3]).String()
}
