package arp

import (
	"testing"
)

func TestArp(t *testing.T) {
	addr, err := Lookup("10.0.0.1")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("addr: %v", addr)
}