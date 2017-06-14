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
// 'Brocot' cree une suite de Brocot de rang 'n'
func Brocot(n int) []*Fract_t {
	u := *Brocot_nums(n)
	m := len(u) - 1
	r := make([]*Fract_t,m+1)
	for k := 0; k <= m; k++ {
		r[k] = Creer(u[k],u[m-k])
	}
	return r
}
// 'Brocot_nums' retourne un pointeur sur un tableau de 2^'n'+1 elements donnant la suite des numerateurs de Brocot
func Brocot_nums(n int) *[]int {
	if 0 > n {panic("Parametre invalide")}
	m := 1 << uint(n) // 2^'n'
	u := make([]int,m+1) // tableau de 2^n'+1 elements
	u[0], u[1] = 0, 1 // initialisation
	for k := 2; k <= m; k++ { // boucler
		if 0 == k & 1 { // 'k' pair
			u[k] = u[k >> 1] // k / 2
		} else { // 'k' impair
			u[k] = u[(k-1) >> 1] + u[(k+1) >> 1]
		}
	}
	return &u // retourner le pointeur
}
// 'Brocot_approx' retourne les deux fractions de la suite de Brocot encadrant 'f' a la precision 'prec'
func Brocot_approx(f, prec float64) ([2]*Fract_t) {
	if prec_min > prec {panic("Precision invalide")} // 'prec_min' est empirique
	e, r := fqa.Ent(f) // 'e' partie entiere, 'r' reste (0 <= 'r' < 1)
	t := [2]*Fract_t{Creer(0,1), Creer(1,1),} // initialiser les fractions d'encadrement (0/1 et 1/1)
	for i := 0; !fqa.Egal_f(t[1].Valeur(),t[0].Valeur(), prec); i++ {  // bloucler tant que la precision
		                                                           // n'est pas suffisante
		m := t[0].Med(t[1]) // calculer la fraction mediane
		                    // (numerateur = somme des numerateurs, denominateur = somme des denominateurs)
		if r < m.Valeur() {t[1] = m} // changer la bonne superieure
		if r > m.Valeur() {t[0] = m} // changer la borne inferieure
	}
	t[0] = t[0].AddInt(e) // ajouter la partie entiere a la borne inferieure
	t[1] = t[1].AddInt(e) // ajouter la partie entiere a la borne superieure
	return t
}
