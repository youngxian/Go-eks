package main

import "testing"
func TestGreetingSpecific(t *testing.T) {
	greeting := CreateGreeting("John")
	if greeting != "Hello, Today is Wednesday\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, John\n")
	}
}
func TestGreetingDefault(t *testing.T) {
	greeting := CreateGreeting("")
	if greeting != "Hello, Guest\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Guest\n")
	}
}
