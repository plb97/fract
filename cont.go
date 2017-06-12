// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package fract

import (
	"github.com/plb97/fqa"
)

func Cont_red(f float64, n int, prec float64) ([]int, []*Fract_t, float64) {
	if 0 >= prec {panic("Precision invalide")}
	if 0 >= n {panic("Parametre invalide")}
	var e, r = make([]int,0,n), make([]*Fract_t,0,n)
	a, b := fqa.Ent(f)
	p, ap := Creer(a,1), Creer(1,0)
	e, r = append(e,a), append(r,p)
	for i := 0; i < n-1 &&
			!fqa.Egal_f(b, 0, prec) &&
			!fqa.Egal_f(1/float64(p.d)/(float64(p.d)/b+float64(ap.d)), 0, prec); i++ {
		a, b = fqa.Ent(1 / b)
		p, ap = Creer(p.n*a+ap.n,p.d*a+ap.d), p
		e, r = append(e,a), append(r,p)
	}
	return e, r, b
}

func Cont(f float64, n int, prec float64) ([]*Fract_t) {
	_, r, _ := Cont_red(f, n, prec)
	return r
}

func Cont_elmts(f float64, n int, prec float64) ([]int) {
	e, _, _ := Cont_red(f, n, prec)
	return e
}

