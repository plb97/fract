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
		{1,[]*Fract_t{New(0,1),New(1,1)}},
		{2,[]*Fract_t{New(0,1),New(1,2),New(1,1)}},
		{3,[]*Fract_t{New(0,1),New(1,3),New(1,2),New(2,3),New(1,1)}},
		{4,[]*Fract_t{New(0,1),New(1,4),New(1,3),New(1,2),New(2,3),New(3,4),New(1,1)}},
		{5,[]*Fract_t{New(0,1),New(1,5),New(1,4),New(1,3),New(2,5),New(1,2),New(3,5),New(2,3),New(3,4),New(4,5),New(1,1)}},
		{6,[]*Fract_t{New(0,1),New(1,6),New(1,5),New(1,4),New(1,3),New(2,5),New(1,2),New(3,5),New(2,3),New(3,4),New(4,5),New(5,6),New(1,1)}},
		{7,[]*Fract_t{New(0,1),New(1,7),New(1,6),New(1,5),New(1,4),New(2,7),New(1,3),New(2,5),New(3,7),New(1,2),New(4,7),New(3,5),New(2,3),New(5,7),New(3,4),New(4,5),New(5,6),New(6,7),New(1,1)}},
		{8,[]*Fract_t{New(0,1),New(1,8),New(1,7),New(1,6),New(1,5),New(1,4),New(2,7),New(1,3),New(3,8),New(2,5),New(3,7),New(1,2),New(4,7),New(3,5),New(5,8),New(2,3),New(5,7),New(3,4),New(4,5),New(5,6),New(6,7),New(7,8),New(1,1)}},
	}
	for _, v := range test_farey {
		attendu := v.f
		obtenu := Farey(v.n)
		if !Equal_s(attendu,obtenu) {
			t.Errorf(test+"(%d): attendu %v, obtenu %v",v.n, attendu, obtenu)
		}
	}
}

