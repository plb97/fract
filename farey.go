package fract

import (

)

func Farey(n int) []*Fract_t {
	if 1 > n {panic("Invalid n")}
	r := make([]*Fract_t,0)
	a,b,c,d := 0, 1, 1, n
	r = append(r,New(a,b))
	for c <= n {
		k := (n+b)/d
		e,f := k*c-a,k*d-b
		a,b,c,d = c,d,e,f
		r = append(r,New(a,b))
	}
	return r
}
