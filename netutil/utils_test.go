package netutil

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ipv4ToUint32(t *testing.T) {
	ip := net.ParseIP("192.168.1.1")
	ipint := IPV4ToUint32(ip)
	assert.Equal(t, net.IP(net.IP{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0xc0, 0xa8, 0x1, 0x1}), ip)
	assert.Equal(t, uint32(0xc0a80101), ipint)
	ipto := Uint32ToIPV4(ipint)
	assert.Equal(t, ip.To4(), ipto.To4())
}
