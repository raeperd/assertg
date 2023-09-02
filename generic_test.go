package assert

import "testing"

func TestComaparable(t *testing.T) {
	t.Parallel()

	t.Run("Equals", func(t *testing.T) {
		t.Parallel()
		Equals(t, true, true)
		Equals(t, false, false)
		Equals(t, 1, 1)
		Equals(t, 1.2, 1.2)
		Equals(t, "string", "string")

		ptr := &struct{ value int }{10}
		eqPtr := ptr
		Equals(t, ptr, eqPtr)
	})

	t.Run("NotEquals", func(t *testing.T) {
		t.Parallel()
		NotEquals(t, true, false)
		NotEquals(t, true, false)
		NotEquals(t, 1, 2)
		NotEquals(t, 1.2, 1.3)
		NotEquals(t, "string", "strings")

		ptr := &struct{ value int }{10}
		nEqPtr := &struct{ value int }{10}
		NotEquals(t, ptr, nEqPtr)
	})

}
