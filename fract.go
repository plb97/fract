// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package fract

import (
	"fmt"
	"math"
	"github.com/plb97/fqa"
)

const prec_min = 1e-18

var PGCD = pgcd1

func pgcd1(a, b int) int { // le pgcd est toujours > 0
	if 0 > a { a = -a }
	if 0 > b { b = -b }
	if a < b {
		a, b = b, a
	}
	if 0 == b {
		if 0 == a {return 1}
		return a
	}
	for r := a % b; 0 < r; r = a % b {
		a, b = b, r
	}
	return b
}
func pgcd2(a, b int) int { // le pgcd est toujours > 0
	if 0 > a { a = -a }
	if 0 > b { b = -b }
	if a < b {
		a, b = b, a
	}
	if 0 == b {
		if 0 == a {return 1}
		return a
	}
	for r := a - b; 0 < r; r = a - b {
		if b > r {
			a, b = b, r
		} else {
			a = r
		}
	}
	return b
}

type Fract_compare_t int
const (
	MINOR Fract_compare_t = iota - 1
	AEQUAL
	MAJOR
	
	LT = MINOR
	EQ = AEQUAL
	GT = MAJOR

	PLUS_PETIT = MINOR
	EGAL = AEQUAL
	PLUS_GRAND = MAJOR
	
)

type Fract_t struct {
	n, d int
}
var (
	m_inf = &Fract_t{-1,0} // - l'infini
	p_inf = &Fract_t{+1,0} // + l'infini
	nan = &Fract_t{0,0} // inderteminee
)
// creation d'une fraction a partir du numerateur et de denominateur
func Creer(n, d int) *Fract_t {
	if 0 == d {
		if 0 < n {return p_inf}
		if 0 > n {return m_inf}
		return nan
	}
	if 0 > d {
		n, d = -n, -d
	}
	p := PGCD(n,d)
	return &Fract_t{n:n/p, d:d/p}
}
// fraction sous forme de chaine
func (r Fract_t)String() string {
	if 1 == r.d {
		return fmt.Sprintf("[%d]",r.n)
	}
	if 0 == r.d {
		return fmt.Sprintf("[%d/%d]",r.n,r.d)
	}
	e, a := fqa.Divent(r.n, r.d)
	if 0 == e {
		return fmt.Sprintf("[%d/%d]",a,r.d)
	}
	return fmt.Sprintf("[%d%+d/%d]",e,a,r.d)
}
// valeur flottante d'une fraction
func (r *Fract_t)Valeur() float64 {
	if r.Egal(m_inf) {return math.Inf(-1)}
	if r.Egal(p_inf) {return math.Inf(+1)}
	if r.Egal(nan) {return math.NaN()}
	return float64(r.n) / float64(r.d)
}
// elements de la fraction (numerateur, denominateur)
func (r *Fract_t)Elmt() (int, int) {
	return r.n, r.d
}
// comparaison d'une fraction
func (r *Fract_t)Compare(f *Fract_t) Fract_compare_t {
	if r.Egal(f) {return AEQUAL}
	if m_inf.Egal(r) || p_inf.Egal(f) {return MINOR}
	if p_inf.Egal(r) || m_inf.Egal(f) {return MAJOR}
	d := r.n * f.d - r.d * f.n
	if nan.Egal(r) {
		d = -f.n
	} else if nan.Egal(f) {
		d = r.n
	}
	if 0 > d {return MINOR}
	if 0 < d {return MAJOR}
	return AEQUAL
}
// verification de l'egalite d'une fraction
func (r *Fract_t)Egal(f *Fract_t) bool {
	// les fractions son reduites 
	// et normalisees (+inf=1/0 ou -inf=-1/0) donc...
	return r.d == f.d && r.n == f. n
}
// addition d'une fraction
func (r *Fract_t)Add(f *Fract_t) *Fract_t {
	return Creer(r.n * f.d + r.d * f.n, r.d * f.d)
}
// soustraction d'une fraction
func (r *Fract_t)Sub(f *Fract_t) *Fract_t {
	return Creer(r.n * f.d - r.d * f.n, r.d * f.d)
}
// multiplication par une fraction
func (r *Fract_t)Mul(f *Fract_t) *Fract_t {
	return Creer(r.n * f.n, r.d * f.d)
}
// division par une fraction
func (r *Fract_t)Div(f *Fract_t) *Fract_t {
	return Creer(r.n * f.d, r.d * f.n)
}
// addition d'un entier
func (r *Fract_t)AddInt(i int) *Fract_t {
	return Creer(r.n + r.d * i, r.d)
}
// soustraction d'un entier
func (r *Fract_t)SubInt(i int) *Fract_t {
	return Creer(r.n - r.d * i, r.d)
}
// multiplication par un entier
func (r *Fract_t)MulInt(i int) *Fract_t {
	return Creer(r.n * i, r.d)
}
// division par un entier ( i peut etre nul)
func (r *Fract_t)DivInt(i int) *Fract_t {
	return Creer(r.n, r.d * i)
}
// determinant
func Det(r *Fract_t,f *Fract_t) int {
	return r.n * f.d - r.d * f.n
}
// mediane
func (r *Fract_t)Med(f *Fract_t) *Fract_t {
	return Creer(r.n + f.n, r.d + f.d)
}
// verification de l'egalite de deux suites
func Egal_s(a,b []*Fract_t) bool {
	l := len(a)
	if l != len(b) {return false}
	for i := 0; i < l; i++ {
		if !a[i].Egal(b[i]) {return false}
	}
	return true
}
