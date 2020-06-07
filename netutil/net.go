package netutil

import (
	"net"
)

func Uint32ToIPV4(ip uint32) net.IP {
	var result [4]byte
	result[0] = byte(ip >> 24)
	result[1] = byte(ip >> 16)
	result[2] = byte(ip >> 8)
	result[3] = byte(ip)
	return net.IPv4(result[0], result[1], result[2], result[3])
}

func IPV4ToUint32(ip net.IP) uint32 {
	ip = ip.To4()
	return uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
}
