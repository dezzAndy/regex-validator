package automatas

import (
	"net/netip"
	"regexp"
	"strings"
)

var regexIpv4 = regexp.MustCompile(`^(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)$`)

func RegexValidarIP(s string) bool {
	if regexIpv4.MatchString(s) {
		return true
	}
	// IPv6 (incluye compresión :: y zona %interfaz) vía net/netip
	if strings.Contains(s, ":") {
		addr, err := netip.ParseAddr(s)
		return err == nil && addr.IsValid()
	}
	return false
}
