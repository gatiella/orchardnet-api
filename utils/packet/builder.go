package packet

import (
	"encoding/binary"
	"net"
)

func BuildTCPSYN(srcIP, dstIP string, dstPort int) []byte {
	ipHdr := make([]byte, 20)
	ipHdr[0] = 0x45
	ipHdr[1] = 0x00
	binary.BigEndian.PutUint16(ipHdr[2:], 40)
	binary.BigEndian.PutUint16(ipHdr[4:], 0)
	binary.BigEndian.PutUint16(ipHdr[6:], 0x4000)
	ipHdr[8] = 64
	ipHdr[9] = 6
	copy(ipHdr[12:16], net.ParseIP(srcIP).To4())
	copy(ipHdr[16:20], net.ParseIP(dstIP).To4())

	tcpHdr := make([]byte, 20)
	binary.BigEndian.PutUint16(tcpHdr[0:], 0) // src port = 0
	binary.BigEndian.PutUint16(tcpHdr[2:], uint16(dstPort))
	binary.BigEndian.PutUint32(tcpHdr[4:], 1000000)
	binary.BigEndian.PutUint32(tcpHdr[8:], 0)
	tcpHdr[12] = 0x50
	tcpHdr[13] = 0x02 // SYN flag
	binary.BigEndian.PutUint16(tcpHdr[14:], 8192)

	return append(ipHdr, tcpHdr...)
}

func BuildUDPPacket(srcIP, dstIP string, dstPort int, payload []byte) []byte {
	ipHdr := make([]byte, 20)
	ipHdr[0] = 0x45
	ipHdr[1] = 0x00
	totalLen := 28 + len(payload)
	binary.BigEndian.PutUint16(ipHdr[2:], uint16(totalLen))
	binary.BigEndian.PutUint16(ipHdr[4:], 0)
	binary.BigEndian.PutUint16(ipHdr[6:], 0x4000)
	ipHdr[8] = 64
	ipHdr[9] = 17 // UDP
	copy(ipHdr[12:16], net.ParseIP(srcIP).To4())
	copy(ipHdr[16:20], net.ParseIP(dstIP).To4())

	udpHdr := make([]byte, 8)
	binary.BigEndian.PutUint16(udpHdr[0:], 0)
	binary.BigEndian.PutUint16(udpHdr[2:], uint16(dstPort))
	binary.BigEndian.PutUint16(udpHdr[4:], uint16(8+len(payload)))
	// UDP checksum omitted

	return append(append(ipHdr, udpHdr...), payload...)
}
