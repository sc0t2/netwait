// Package netwait provides functions for waiting on a TCP network service to become available
package netwait

import (
	"net"
	"time"
)

func dialOnce(host, port string) error {
	conn, err := net.Dial("tcp", net.JoinHostPort(host, port))
	if conn != nil {
		defer conn.Close()
	}
	return err
}

// Try attempts to connect to `host`:`port` up to `times` times, sleeping for `interval` between tries
func Try(times int, interval time.Duration, host, port string) error {
	var err error
	for i := 0; i < times; i++ {
		err = dialOnce(host, port)
		if err == nil {
			return nil
		}
		time.Sleep(interval)
	}
	return err
}
