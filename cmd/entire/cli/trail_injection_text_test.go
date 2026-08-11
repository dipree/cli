package cli

import "testing"

// The first-turn injection documents trails without advertising unrelated
// Entire commands or relying on transient trail state.
func TestEntireTrailContextInjection(t *testing.T) {
	t.Parallel()

	const want = "Trails are Entire's replacement for pull requests and issues. A trail ties together a branch's context, discussion, findings, and review state. Use `entire trail` to view, create, update, or watch one."
	if got := entireTrailContextInjection(); got != want {
		t.Fatalf("entireTrailContextInjection() = %q, want %q", got, want)
	}
}
