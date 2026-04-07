package ws

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "你好"
	if got != want {
		t.Errorf("Hello() = %q; want %q", got, want)
	}
}
