package fract

import (
	"util"
)

func Cont_red(f float64, n int, prec float64) ([]int, []*Fract_t, float64) {
	if 0 >= prec {panic("Invalid prec")}
	if 0 >= n {panic("Invalid n")}
	var e, r = make([]int,0,n), make([]*Fract_t,0,n)
	a, b := util.Ent(f)
	p, ap := New(a,1), New(1,0)
	e, r = append(e,a), append(r,p)
	for i := 0; i < n-1 &&
				!util.Equal_f(b, 0, prec) &&
				!util.Equal_f(1/float64(p.d)/(float64(p.d)/b+float64(ap.d)), 0, prec); i++ {
		a, b = util.Ent(1 / b)
		p, ap = New(p.n*a+ap.n,p.d*a+ap.d), p
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

