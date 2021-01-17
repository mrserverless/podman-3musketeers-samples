package main

import "testing"

func TestHello(t *testing.T) {
	// given
	want := "Hello, world"

	// when
	got := Hello()
	
	// then
	if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}