package dependency

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{} // creating a new bufffer type
	Greet(&buffer, "Chris")
	got := buffer.String()
	want := "Hello, Chris"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
