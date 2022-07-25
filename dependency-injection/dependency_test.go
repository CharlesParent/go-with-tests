package dependency

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Charles")

	got := buffer.String()
	want := "Hello, Charles"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
