package assertion

import (
	"fmt"
	"net"
)

func AssertTcpSocketAcceptsConnection(addr, msg string) (bool, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return false, fmt.Errorf("%s: %s", msg, err)
	}
	defer conn.Close()
	return true, nil
}