package fract

import (
    "testing"
	"math"
)

type test_cont_t struct {
	f float64
	r []*Fract_t
}
func Test_cont(t *testing.T) {
	var test_cont = []test_cont_t {
		{math.Pi,[]*Fract_t{New(3,1),New(22,7),New(333,106),New(355,113),New(103993,33102),}},
	}
	for _, v := range test_cont {
		expected := v.r
		n := len(expected)
		actual := Cont(v.f,n,1e-20)
		if n != len(actual) {
			t.Errorf("cont(%f,%d): expected n %v, actual %v",v.f,n,n, len(actual))
		} else {
			for i := 0; i < n; i++ {
				if !expected[i].Equal(actual[i]) {
					t.Errorf("cont(%f,%d): [%d] expected %v, actual %v",v.f,n, i, expected[i], actual[i])
				}
			}
		}
		actual = Cont(v.f,10,1e-9)
		if n != len(actual) {
			t.Errorf("cont(%f,%d): expected n %v, actual %v",v.f,n,n, len(actual))
		} else {
			for i := 0; i < n; i++ {
				if !expected[i].Equal(actual[i]) {
					t.Errorf("cont(%f,%d): [%d] expected %v, actual %v",v.f,n, i, expected[i], actual[i])
				}
			}
		}
	}
}

