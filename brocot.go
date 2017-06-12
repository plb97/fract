// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package fract

import (
	"github.com/plb97/fqa"
)

// http://mapage.noos.fr/r.ferreol/atelecharger/textes/brocot.pdf
func Brocot(n int) []*Fract_t {
	if 0 > n {panic("Parametre invalide")}
	u := *Brocot_nums(uint(n))
	m := len(u) - 1
	r := make([]*Fract_t,m+1)
	for k := 0; k <= m; k++ {
		r[k] = Creer(u[k],u[m-k])
	}
	return r
}

func Brocot_nums(n uint) *[]int {
	m := 1 << n
	u := make([]int,m+1)
	u[0], u[1] = 0, 1
	for k := 2; k <= m; k++ {
		if 0 == k%2 {
			u[k] = u[k/2]
		} else {
			u[k] = u[(k-1)/2] + u[(k+1)/2]
		}
	}
	return &u
}

func Brocot_approx(f, prec float64) ([2]*Fract_t) {
	if 0 > prec {panic("Precision invalide")}
	e, r := fqa.Ent(f)
	t := [2]*Fract_t{Creer(0,1), Creer(1,1),}
	for i := 0; !fqa.Egal_f(t[1].Valeur(),t[0].Valeur(), prec); i++ {
		m := t[0].Med(t[1])
		if r < m.Valeur() {t[1] = m}
		if r > m.Valeur() {t[0] = m}
	}
	t[0] = t[0].AddInt(e)
	t[1] = t[1].AddInt(e)
	return t
}
