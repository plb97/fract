// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package fract
// http://jcresson.perso.univ-pau.fr/nombre-ford.pdf
// 'Farey' retourne un tableau des fractions de la suite de Farey 'n' ('n' >= 1) qui contient 2^'n'+1 elements
func Farey(n int) []*Fract_t {
	if 1 > n {panic("Parametre invalide")}
	m := 1<<uint(n)+1
	r := make([]*Fract_t,0,m)
	a,b,c,d := 0, 1, 1, n
	r = append(r, Creer(a,b))
	for c <= n {
		k := (n+b)/d
		e,f := k*c-a,k*d-b
		a,b,c,d = c,d,e,f
		r = append(r, Creer(a,b))
	}
	return r
}
