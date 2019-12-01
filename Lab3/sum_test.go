package Lab3

import "testing"

func assertEquals(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected and actual differ:\n<%v>\n<%v>\n", expected, actual)
	}
}

func Test_sumOfNothing_isZero(t *testing.T) {
	assertEquals(t, 0, Sum())
}

func Test_sumOfOne_isOne(t *testing.T) {
	assertEquals(t, 1, Sum(1))
}

func Test_sumOfTwoNumbers(t *testing.T) {
	assertEquals(t, 12, Sum(5, 7))
}