package netwait

import (
	"testing"
	"time"
)

func TestTry(t *testing.T) {
	t.Run("google.com", func(t *testing.T) {
		got := Try(3, time.Second*2, "google.com", "80")
		if got != nil {
			t.Fatalf("expected nil but got %q\n", got)
		}
	})

	t.Run("nonexistent.test", func(t *testing.T) {
		before := time.Now()
		got := Try(3, time.Second*1, "nonexistent.test", "1234")
		if got == nil {
			t.Fatalf("got nil but expected error")
		} else if time.Since(before) < time.Second*3 {
			t.Fatalf("did not wait for 3 seconds")
		}
	})
}
