// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package fract

func Farey(n int) []*Fract_t {
	if 1 > n {panic("Parametre invalide")}
	r := make([]*Fract_t,0)
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
