// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package fract

import (
	"github.com/plb97/fqa"
)
// http://www-irma.u-strasbg.fr/~richez/ressources/recherche/memoire_fractions_continues.pdf
// 'Cont_red' retourne :
// le tableau des elements de la suite continue (reduite (suite finie) de la fraction continue)
// le tableau des fractions continues successives (la derniere etant la meilleure approximation)
// la valeur majorant l'erreur d'approximation inferieure ou egale a 'prec'
//
// p[0] = a[0]	p[1] = a[1]*a[0] + 1	et pour n>=2	p[n] = a[n]*p[n-1] + p[n-2]
// q[0] = 1	q[1] = a[1]		et pour n>=2	q[n] = a[n]*q[n-1] + q[n-2]
// REMARQUE : en posant p[-1] = 1 et q[-1] = 0 on obtient
//		p[1] = a[1]*p[0] + p[-1] = a[1]*p[0] + 1
//		q[1] = a[1]*q[0] + q[-1] = a[1]*q[0] + 0
func Cont_red(f float64, prec float64) ([]int, []*Fract_t, float64) {
	if prec_min > prec {panic("Precision invalide")} // 'prec_min' est empirique
	var r, e = make([]int,0), make([]*Fract_t,0) // 'e' tableau des elements de la suite continue,
	                                             // 'r' tableau des fractions correspondantes
	a, b := fqa.Ent(f) // 'a' partie entiere, 'b' reste fractionnaire 0 <= 'b' < 1
	d := 1e0
	p, ap := Creer(a,1), Creer(1,0) // a/1 , 1/0 (+Inf)
	r, e = append(r,a), append(e,p) // initialiser les tableaux
	arret := func(p,ap *Fract_t, b float64) bool {
		d = 1 / float64(p.d*p.d)
		ok := fqa.Egal_f(b, 0, prec) || // 'b' est nul (ou presque)
			fqa.Egal_f(d, 0, prec) // l'ecart |'p.n/p.d' - 'f'| < 1 / p.d*p.d est inferieur a 'prec'
		return ok
	}
	for !arret(p,ap,b) { // boucler
		a, b = fqa.Ent(1 / b) // 'a' partie entiere, 'b' reste fractionnaire 0 <= 'b' < 1
		p, ap = Creer(a*p.n+ap.n,a*p.d+ap.d), p // 'p' nouvelle fraction, 'ap' fraction precedente
		r, e = append(r,a), append(e,p) // ajouter les elements aux tableaux
	}
	return r, e, d // retourner les resultats
}
// 'Cont' retourne un tableau des fractions continues qui approchent 'f'
func Cont(f float64, prec float64) ([]*Fract_t) {
	_, r, _ := Cont_red(f, prec)
	return r
}
// Cont_elmts' retourne le tableau des elements de la suite continue approchant 'f'
func Cont_elmts(f float64, prec float64) ([]int) {
	e, _, _ := Cont_red(f, prec)
	return e
}

