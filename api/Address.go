package api

import (
	"strconv"
)

type Address struct {
	Host string
	Port int
}

func (address *Address) GetAddressString() string {
	portString := strconv.Itoa(address.Port)
	return address.Host + ":" + portString
}
