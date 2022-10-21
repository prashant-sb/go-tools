package resolv

import (
	"net"
	"regexp"
	"strings"

	log "github.com/golang/glog"
)

// Validate hostname format
func IsValidName(name string) bool {
	name = strings.Trim(name, " ")
	re, _ := regexp.Compile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`)
	return re.MatchString(name)
}

// Validate IPv4 or IPv6 addresses
func IsValidIP(ipAddress string) bool {
	ip := net.ParseIP(ipAddress)

	return !(ip.To4() == nil && ip.To16() == nil)
}

// Lookup for Host
func ResolveIP(name string) error {
	ips, err := net.LookupIP(name)
	if err != nil {
		return err
	}

	for _, ip := range ips {
		log.Info(name + " IN A " + ip.String())
	}

	return nil
}

// Performs reverse lookup from IP to Host
func ResolveHost(ipAddress string) error {
	hosts, err := net.LookupAddr(ipAddress)
	if err != nil {
		return err
	}
	for _, dname := range hosts {
		log.Info(ipAddress+".in-addr.arpa", " name = "+dname)
	}

	return nil
}
