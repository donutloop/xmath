package gcd

import "testing"

func TestGcd(t *testing.T) {

	x := Eulidean(287, 91)
	t.Log(x)
	if x != 7 {
		t.Fatal("gcd is wrong")
	}
}
