package automatas

import "testing"

func TestValidarIP(t *testing.T) {
	casos := []struct {
		entrada  string
		esperado bool
	}{
		{"192.168.1.1", true},
		{"255.255.255.255", true},
		{"0.0.0.0", true},
		{"256.1.1.1", false},
		{"01.2.3.4", false},
		{"192.168.1", false},
		{"192.168.1.1.1", false},
		{"abcd:ef01:2345:6789:abcd:ef01:2345:6789", true},
		{"2001:db8::1", true},
		{"::1", true},
		{"fe80::1%eth0", true},
		{"12345::1", false},
		{"g:1::1", false},
		{"", false},
	}

	for _, c := range casos {
		resultado := RegexValidarIP(c.entrada)
		if resultado != c.esperado {
			t.Errorf("Para entrada %q se esperaba %v, pero se obtuvo %v", c.entrada, c.esperado, resultado)
		}
	}
}
