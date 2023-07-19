// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"database/sql/driver"
	"fmt"
	"net"
	"reflect"
)

type VisitorIP net.IP

func ParseVisitorIP(ip string) VisitorIP {
	return VisitorIP(net.ParseIP(ip))
}

func (v VisitorIP) NetIP() net.IP {
	return net.IP(v)
}

func (v VisitorIP) RegionCode() (string, error) {
	return findRegionCodeByIPAddress(v.NetIP())
}

func (v VisitorIP) String() string {
	return v.NetIP().String()
}

func (v VisitorIP) GormDataType() string {
	return "string"
}

func (v *VisitorIP) Scan(value interface{}) error {
	ipAddressBytes, ok := value.([]byte)
	if !ok {
		valueType := reflect.TypeOf(value)
		return fmt.Errorf("%v is %s , not []byte", value, valueType)
	}

	ipAddressString := string(ipAddressBytes)
	*v = ParseVisitorIP(ipAddressString)
	return nil
}

func (v VisitorIP) Value() (driver.Value, error) {
	value := v.NetIP().String()
	return value, nil
}
