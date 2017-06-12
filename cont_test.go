// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
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
	test := "cont"
	var test_cont = []test_cont_t {
		{math.Pi,[]*Fract_t{New(3,1),New(22,7),New(333,106),New(355,113),New(103993,33102),}},
	}
	for _, v := range test_cont {
		attendu := v.r
		n := len(attendu)
		obtenu := Cont(v.f,n,1e-20)
		if n != len(obtenu) {
			t.Errorf(test+"(%f,%d): attendu n %v, obtenu %v",v.f,n,n, len(obtenu))
		} else {
			for i := 0; i < n; i++ {
				if !attendu[i].Egal(obtenu[i]) {
					t.Errorf(test+"(%f,%d): [%d] attendu %v, obtenu %v",v.f,n, i, attendu[i], obtenu[i])
				}
			}
		}
		obtenu = Cont(v.f,10,1e-9)
		if n != len(obtenu) {
			t.Errorf(test+"(%f,%d): attendu n %v, obtenu %v",v.f,n,n, len(obtenu))
		} else {
			for i := 0; i < n; i++ {
				if !attendu[i].Egal(obtenu[i]) {
					t.Errorf("cont(%f,%d): [%d] attendu %v, obtenu %v",v.f,n, i, attendu[i], obtenu[i])
				}
			}
		}
	}
}

