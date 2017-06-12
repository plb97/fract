	package fract

import (
    "testing"
)

type test_brocot_t struct {
	n int
	f []*Fract_t
}
func Test_brocot(t *testing.T) {
	var test_brocot = []test_brocot_t {
		{0,[]*Fract_t{New(0,1),New(1,0)}},
		{1,[]*Fract_t{New(0,1),New(1,1),New(1,0)}},
		{2,[]*Fract_t{New(0,1),New(1,2),New(1,1),New(2,1),New(1,0)}},
		{3,[]*Fract_t{New(0,1),New(1,3),New(1,2),New(2,3),New(1,1),New(3,2),New(2,1),New(3,1),New(1,0)}},
		{4,[]*Fract_t{New(0,1),New(1,4),New(1,3),New(2,5),New(1,2),New(3,5),New(2,3),New(3,4),New(1,1),New(4,3),New(3,2),New(5,3),New(2,1),New(5,2),New(3,1),New(4,1),New(1,0)}},
	}
	for _, v := range test_brocot {
		expected := v.f
		actual := Brocot(v.n)
		if !Equal_s(expected,actual) {
			t.Errorf("brocot(%d): expected %v, actual %v",v.n, expected, actual)
		}
	}
}

