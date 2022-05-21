package email

import "testing"

func TestGreeting(t *testing.T) {
	if Greeting("Kenan") != "Hello Kenan" {
		t.Error("Error!")
	}
}
