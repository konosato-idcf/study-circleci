package main

import "testing"

func TestHello(t *testing.T) {
	actual := Hello("CircleCI")
	expected := "Hello CircleCI"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
