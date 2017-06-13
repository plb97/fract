// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package fract

import (
	"testing"
	"math"
	"github.com/plb97/fqa"
)

type test_cont_t struct {
	f float64
	r []*Fract_t
}
func Test_cont(t *testing.T) {
	test := "cont"
	var test_cont = []test_cont_t {
		{math.Pi,[]*Fract_t{Creer(3,1), Creer(22,7), Creer(333,106), Creer(355,113), Creer(103993,33102),}},
	}
	for _, v := range test_cont {
		attendu := v.r
		n := len(attendu)
		p := 1e-20
		obtenu := Cont(v.f,n,p)
		if n < len(obtenu) {
			t.Errorf(test+"(%f,%d): attendu n %v >= obtenu %v",v.f,n,n, len(obtenu))
		} else {
			for i := 0; i < len(obtenu); i++ {
				if !attendu[i].Egal(obtenu[i]) {
					t.Errorf(test+"(%f,%d): [%d] attendu %v, obtenu %v",v.f,n, i, attendu[i], obtenu[i])
				}
			}
		}
		p = 1e-9
		obtenu = Cont(v.f,n,p)
		if n < len(obtenu) {
			t.Errorf(test+"(%f,%d): attendu n %v >= obtenu %v",v.f,n,n, len(obtenu))
		} else {
			for i := 0; i < len(obtenu); i++ {
				if !attendu[i].Egal(obtenu[i]) {
					t.Errorf(test+"(%f,%d): [%d] attendu %v, obtenu %v",v.f,n, i, attendu[i], obtenu[i])
				}
			}
		}
	}
}

func Test_elmts(t *testing.T) {
	test := "elmts"
	f := math.Pi
	fc := []int{3,7,15,1,292}
	n := 5
	p := 1e-7
	{
		lo := Cont_elmts(f,n,p)
		for i, attendu := range fc {
			obtenu := lo[i]
			if attendu != obtenu {
				t.Errorf(test+" : attendu %v != obtenu %v", attendu, obtenu)
			}
		}
	}

}

func Test_red(t *testing.T) {
	test := "red"
	f := math.Phi
	li := []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}
	n := len(li) // suffisamment grand
	p := 1e-6 // precision souhaitee (pas trop grande)
	e,r,b := Cont_red(f,n,p)
	{
		if len(li) < len(e) {
			t.Errorf(test+" : attendu %v >= obtenu %v", len(li), len(e))
		}
		for i, obtenu := range e {
			attendu := li[i]
			if attendu != obtenu {
				t.Errorf(test+" : attendu %v != obtenu %v", attendu, obtenu)
			}
		}
	}
	{
		attendu := f - 1
		obtenu := b
		if !fqa.Egal_f(obtenu,attendu,p) {
			t.Errorf(test+" : attendu %v != obtenu %v prec=%v ecart=%v", attendu, obtenu,p,obtenu - attendu)
		}
	}
	{
		attendu := f
		obtenu := r[len(r)-1].Valeur()
		if !fqa.Egal_f(obtenu,attendu,p) {
			t.Errorf(test+" : attendu %v != obtenu %v prec=%v ecart=%v", attendu, obtenu,p,obtenu - attendu)
		}
	}
}