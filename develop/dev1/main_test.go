package main

import "testing"

func TestExactTime(t *testing.T) {
	host1 := "0.beevik-ntp.pool.ntp.org"
	_, err1 := ExactTime(host1)
	if err1 != nil {
		t.Fatalf("Error is not empty: %s host %s", err1.Error(), host1)
	}
	host2 := "wrong host"
	_, err2 := ExactTime(host2)
	if err2 == nil {
		t.Fatalf("Error is empty: %v host %s", err2, host2)
	}
}
