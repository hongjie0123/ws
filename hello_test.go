package ws

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "hello"
	if got != want {
		t.Errorf("Hello() = %q; want %q", got, want)
	}
}
