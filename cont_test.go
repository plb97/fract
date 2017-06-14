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
		p := 1e-5
		obtenu := Cont(v.f,p)
		n := len(obtenu)
		if n > len(attendu) {
			n = len(attendu)
		}
		for i := 0; i < n; i++ {
			if !attendu[i].Egal(obtenu[i]) {
				t.Errorf(test+"(%f,%d): [%d] attendu %v, obtenu %v",v.f,n, i, attendu[i], obtenu[i])
			}
		}
		p = 1e-5
		obtenu = Cont(v.f,p)
		n = len(obtenu)
		if n > len(attendu) {
			n = len(attendu)
		}
		for i := 0; i < n; i++ {
			if !attendu[i].Egal(obtenu[i]) {
				t.Errorf(test+"(%f,%d): [%d] attendu %v, obtenu %v",v.f,n, i, attendu[i], obtenu[i])
			}
		}
		}
}

func Test_elmts(t *testing.T) {
	test := "elmts"
	f := math.Pi
	fc := []int{3,7,15,1,292}
	{
		p := prec_min
		lo := Cont_elmts(f,p)
		n := len(lo)
		if n > len(fc) {
			n = len(fc)
		}
		for i := 0; i < n; i++ {
			attendu := fc[i]
			obtenu := lo[i]
			if attendu != obtenu {
				t.Errorf(test+" : attendu %v != obtenu %v", attendu, obtenu)
			}
		}
	}

}

func Test_prec(t *testing.T) {
	test := "prec"
	f := math.Pi
	{
		ctr := 0
		defer func() {
			if r := recover(); r != nil {
				ctr++
				attendu := "Precision invalide"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		p := prec_min / 2
		Cont_elmts(f,p)
		if 1 != ctr {
			t.Errorf(test+": attendu %v != obtenu %v\n", 1, ctr)
		}
	}

}

func Test_red(t *testing.T) {
	test := "red"
	f := math.Phi
	li := []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}
	p := 1e-8 // precision souhaitee (pas trop grande)
	r, e, d := Cont_red(f,p)
	{
		n := len(r)
		if n > len(li) {
			n = len(li)
		}
		for i := 0; i < i; i++ {
			attendu := li[i]
			obtenu := r[i]
			if attendu != obtenu {
				t.Errorf(test+" : attendu %v != obtenu %v", attendu, obtenu)
			}
		}
	}
	{
		attendu := f
		obtenu := e[len(e)-1].Valeur()
		if !fqa.Egal_f(obtenu,attendu,p) {
			t.Errorf(test+" : attendu %v != obtenu %v prec=%v ecart=%v",attendu,obtenu,p,obtenu - attendu)
		}
	}
	{
		attendu := f
		obtenu := e[len(e)-1].Valeur()
		if !fqa.Egal_f(obtenu,attendu,p) {
			t.Errorf(test+" : attendu %v != obtenu %v prec=%v ecart=%v",attendu,obtenu,p,obtenu - attendu)
		}
	}
	{
		attendu := p
		obtenu := d
		if obtenu > attendu {
			t.Errorf(test + " : attendu %v < obtenu %v", attendu, obtenu)
		}
	}
}