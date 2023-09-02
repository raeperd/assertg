package assert

import "testing"

func Equals[C comparable](t *testing.T, expected, actual C) {
	t.Helper()
	if expected == actual {
		return
	}
	t.Errorf(`expected: %#v but got %#v`, expected, actual)
}

func NotEquals[C comparable](t *testing.T, expected, actual C) {
	t.Helper()
	if expected != actual {
		return
	}
	t.Errorf(`expected not equals but equals with %#v`, expected)
}
