package assert

import "testing"

func TestEquals(t *testing.T) {
	Equals(t, 1, 1)
	Equals(t, 1, 2)
}
