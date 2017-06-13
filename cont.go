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
// 'Cond_red' retourne :
// le tableau des elements('n' au plus depend de 'prec') de la suite continue
// le tableau des elements('n' au plus depend de 'prec') des fractions continues correspondantes
// la valeur approchee a la precision 'prec' au plus (depend de la valeur 'n') de la partie fractionnaire de 'f'
// la partie entiere de 'f'
// la valeur majorant l'erreur d'approximation
// 'e' est la reduite (suite finie) de la fraction continue
//
// p[0] = a[0]	p[1] = a[1]*a[0] + 1	et pour n>=2	p[n] = a[n]*p[n-1] + p[n-2]
// q[0] = 1	q[1] = a[1]		et pour n>=2	q[n] = a[n]*q[n-1] + q[n-2]
// REMARQUE : en posant p[-1] = 1 et q[-1] = 0 on obtient
//		p[1] = a[1]*p[0] + p[-1] = a[1]*p[0] + 1
//		q[1] = a[1]*q[0] + q[-1] = a[1]*q[0] + 0
// TODO l'utilisation de 'n' et 'prec' comme criteres d'arret pourrait etre a revoir ('prec' devrait rester seul)
func Cont_red(f float64, n int, prec float64) ([]int, []*Fract_t, float64, int, float64) {
	if 0 >= prec {panic("Precision invalide")}
	if 0 >= n {panic("Parametre invalide")}
	var e, r = make([]int,0,n), make([]*Fract_t,0,n) // 'e' tableau des elements de la suite continue, 'r' tableau des fractions correspondantes
	a, b := fqa.Ent(f) // 'a' partie entiere, 'b' reste fractionnaire 0 <= 'b' < 1
	c, d := a, 1e0
	p, ap := Creer(a,1), Creer(1,0) // a/1 , 1/0 (+Inf)
	e, r = append(e,a), append(r,p) // initialiser les tableaux
	arret := func(p,ap *Fract_t, b float64, i int) bool {
		d = 1 / float64(p.d*p.d)
		ok := i >= n || // 'n' est atteint
			fqa.Egal_f(b, 0, prec) || // 'b' est nul (ou presque)
			fqa.Egal_f(d, 0, prec) // l'ecart |'p.n/p.d' - 'f'| < 1 / p.d*p.d est nul (ou presque)
		return ok
	}
	for i := 1; !arret(p,ap,b,i); i++ { // boucler
		a, b = fqa.Ent(1 / b) // 'a' partie entiere, 'b' reste fractionnaire 0 <= 'b' < 1
		p, ap = Creer(a*p.n+ap.n,a*p.d+ap.d), p // 'p' nouvelle fraction, 'ap' fraction precedente
		e, r = append(e,a), append(r,p) // ajouter les elements aux tableaux
	}
	return e, r, b, c, d // retourner les resultats
}
// 'Cont' retourne un tableau (d'au plus 'n' elements') des fractions continues qui approchent 'f'
func Cont(f float64, n int, prec float64) ([]*Fract_t) {
	_, r, _, _, _ := Cont_red(f, n, prec)
	return r
}
// Cont_elmts' retourne le tableau des au plus 'n' elements de la suite continue approchant 'f'
func Cont_elmts(f float64, n int, prec float64) ([]int) {
	e, _, _, _, _ := Cont_red(f, n, prec)
	return e
}

