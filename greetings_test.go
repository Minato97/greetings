package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := hello("Juan")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`hello("Juan") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := hello("")
	if msg != "" || err == nil {
		t.Fatalf(`hello("") = %q, %v, nil`, msg, err)
	}
}
