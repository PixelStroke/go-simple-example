package greet

import (
	"regexp"
	"testing"
)

func TestOne(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := One(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`One(%q) = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

func TestOneEmpty(t *testing.T) {
	msg, err := One("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
