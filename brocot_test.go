// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package fract

import (
    "testing"
	"math"
	"reflect"
)

type test_brocot_t struct {
	n int
	f []*Fract_t
}
func Test_brocot(t *testing.T) {
	test := "brocot"
	var test_brocot = []test_brocot_t {
		{0,[]*Fract_t{Creer(0,1), Creer(1,0)}},
		{1,[]*Fract_t{Creer(0,1), Creer(1,1), Creer(1,0)}},
		{2,[]*Fract_t{Creer(0,1), Creer(1,2), Creer(1,1), Creer(2,1), Creer(1,0)}},
		{3,[]*Fract_t{Creer(0,1), Creer(1,3), Creer(1,2), Creer(2,3), Creer(1,1), Creer(3,2), Creer(2,1), Creer(3,1), Creer(1,0)}},
		{4,[]*Fract_t{Creer(0,1), Creer(1,4), Creer(1,3), Creer(2,5), Creer(1,2), Creer(3,5), Creer(2,3), Creer(3,4), Creer(1,1), Creer(4,3), Creer(3,2), Creer(5,3), Creer(2,1), Creer(5,2), Creer(3,1), Creer(4,1), Creer(1,0)}},
	}
	for _, v := range test_brocot {
		attendu := v.f
		obtenu := Brocot(v.n)
		if !Egal_s(attendu,obtenu) {
			t.Errorf(test+"(%d): attendu %v, obtenu %v",v.n, attendu, obtenu)
		}
	}
}

func Test_brocot_approx(t *testing.T) {
	test := "approx"
	f := math.Pi
	p := 1e-4
	a := Brocot_approx(f,p)
	{
		attendu := f
		obtenu := a[0].Valeur()
		ecart := obtenu - attendu
		if ecart < -p || p < ecart {
			t.Errorf(test+" : precision=%.4f ecart=%.5f attendu %.4f, obtenu %.4f",p, ecart , attendu, obtenu)
		}
	}
	{
		attendu := f
		obtenu := a[1].Valeur()
		ecart := obtenu - attendu
		if ecart < -p || p < ecart {
			t.Errorf(test+" : precision=%.4f ecart=%.5f attendu %.4f, obtenu %.4f",p, ecart , attendu, obtenu)
		}
	}
	{
		ecart := a[1].Valeur() - a[0].Valeur()
		if ecart < 0 || p < ecart {
			t.Errorf(test+" : precision=%.4f ecart=%.5f",p, ecart )
		}
	}
	{
		attendu := -1
		obtenu := Det(a[0],a[1]) // le determinant de deux termes consecutifs est egal a '-1'
		if attendu != obtenu {
			t.Errorf(test+" : attendu %v != obtenu %v", attendu, obtenu)
		}
	}
}

func Test_brocot_nums(t *testing.T) {
	test := "nums"
	la := [][]int { 0:{0,1}, 1:{0,1,1}, 2:{0,1,1,2,1}, 3:{0,1,1,2,1,3,2,3,1},4:{0,1,1,2,1,3,2,3,1,4,3,5,2,5,3,4,1}}
	for i := range la {
		attendu := la[i]
		obtenu := *Brocot_nums(i)
		if !reflect.DeepEqual(attendu,obtenu) {
			t.Errorf(test+" : attendu %v != obtenu %v", attendu, obtenu)
		}
	}
}