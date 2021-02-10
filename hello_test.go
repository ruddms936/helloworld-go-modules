package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	want := "안녕, 세상."
	fmt.Println(want)
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
