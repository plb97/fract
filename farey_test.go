// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package fract

import (
    "testing"
)

type test_farey_t struct {
	n int
	f []*Fract_t
}
func Test_farey(t *testing.T) {
	test := "farey"
	var test_farey = []test_farey_t {
		{1,[]*Fract_t{Creer(0,1), Creer(1,1)}},
		{2,[]*Fract_t{Creer(0,1), Creer(1,2), Creer(1,1)}},
		{3,[]*Fract_t{Creer(0,1), Creer(1,3), Creer(1,2), Creer(2,3), Creer(1,1)}},
		{4,[]*Fract_t{Creer(0,1), Creer(1,4), Creer(1,3), Creer(1,2), Creer(2,3), Creer(3,4), Creer(1,1)}},
		{5,[]*Fract_t{Creer(0,1), Creer(1,5), Creer(1,4), Creer(1,3), Creer(2,5), Creer(1,2), Creer(3,5),
			Creer(2,3), Creer(3,4), Creer(4,5), Creer(1,1)}},
		{6,[]*Fract_t{Creer(0,1), Creer(1,6), Creer(1,5), Creer(1,4), Creer(1,3), Creer(2,5), Creer(1,2),
			Creer(3,5), Creer(2,3), Creer(3,4), Creer(4,5), Creer(5,6), Creer(1,1)}},
		{7,[]*Fract_t{Creer(0,1), Creer(1,7), Creer(1,6), Creer(1,5), Creer(1,4), Creer(2,7), Creer(1,3),
			Creer(2,5), Creer(3,7), Creer(1,2), Creer(4,7), Creer(3,5), Creer(2,3), Creer(5,7), Creer(3,4),
			Creer(4,5), Creer(5,6), Creer(6,7), Creer(1,1)}},
		{8,[]*Fract_t{Creer(0,1), Creer(1,8), Creer(1,7), Creer(1,6), Creer(1,5), Creer(1,4), Creer(2,7),
			Creer(1,3), Creer(3,8), Creer(2,5), Creer(3,7), Creer(1,2), Creer(4,7), Creer(3,5), Creer(5,8),
			Creer(2,3), Creer(5,7), Creer(3,4), Creer(4,5), Creer(5,6), Creer(6,7), Creer(7,8),
			Creer(1,1)}},
	}
	for _, v := range test_farey {
		attendu := v.f
		obtenu := Farey(v.n)
		if !Egal_s(attendu,obtenu) {
			t.Errorf(test+"(%d): attendu %v, obtenu %v",v.n, attendu, obtenu)
		}
	}
}

