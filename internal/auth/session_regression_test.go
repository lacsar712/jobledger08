package auth

import "testing"

// TestNilDisplayNameRegression guards against the nil-receiver panic that
// previously occurred in (*Session).DisplayName. A nil receiver must yield an
// empty string rather than dereferencing s.Name.
func TestNilDisplayNameRegression(t *testing.T) {
	var s *Session
	got := func() (out string) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("DisplayName panicked on nil receiver: %v", r)
			}
		}()
		out = s.DisplayName()
		return
	}()
	if got != "" {
		t.Fatalf("expected empty string for nil receiver, got %q", got)
	}
}
