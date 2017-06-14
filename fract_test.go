// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package fract

import (
    "testing"
)

type test_fract_new_t struct {
	n, d int
	f *Fract_t
}
func Test_fract_creer(t *testing.T) {
	test := "creer"
	var test_fract_new = []test_fract_new_t {
		{1,2,&Fract_t{1,2}},
		{1,3,&Fract_t{1,3}},
		{5,15,&Fract_t{1,3}},
		{-5,15,&Fract_t{-1,3}},
		{5,-15,&Fract_t{-1,3}},
		{-5,-15,&Fract_t{1,3}},
		{5,0,&Fract_t{1,0}},
		{-5,0,&Fract_t{-1,0}},
	}
	for _, v := range test_fract_new {
		attendu := v.f
		obtenu := Creer(v.n,v.d)
		if !attendu.Egal(obtenu) {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}
}

type test_fract_egal_t struct {
	a, b *Fract_t
	c bool
}
func Test_fract_egal(t *testing.T) {
	test := "egal"
	var test_fract_egal = []test_fract_egal_t{
		{Creer(1,2), Creer(2,3),false},
		{Creer(-1,2), Creer(2,3),false},
		{Creer(2,3), Creer(1,2),false},
		{Creer(2,3), Creer(-1,2),false},
		{Creer(1,2), Creer(1,0),false},
		{Creer(-1,2), Creer(1,0),false},
		{Creer(1,0), Creer(1,2),false},
		{Creer(1,0), Creer(-1,2),false},
		{Creer(1,2), Creer(2,4),true},
		{Creer(-1,2), Creer(2,-4),true},
		{Creer(2,4), Creer(1,2),true},
		{Creer(2,-4), Creer(-1,2),true},
		{Creer(1,0), Creer(2,0),true},
		{Creer(-1,0), Creer(-2,0),true},
		{Creer(2,0), Creer(1,0),true},
		{Creer(0,0), Creer(0,0),true},
		{Creer(-2,0), Creer(-1,0),true},
		{Creer(1,0), Creer(-2,0),false},
		{Creer(-1,0), Creer(2,0),false},
		{Creer(-2,0), Creer(1,0),false},
		{Creer(2,0), Creer(-1,0),false},
	}
	for _, v := range test_fract_egal {
		attendu := v.c
		obtenu := v.a.Egal(v.b)
		if attendu != obtenu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}
}

type test_fract_compare_t struct {
	a, b *Fract_t
	c Fract_compare_t
}
func Test_fract_compare(t *testing.T) {
	test := "compare"
	var test_fract_compare = []test_fract_compare_t {
		{Creer(1,2), Creer(2,3),MINOR},
		{Creer(-1,2), Creer(2,3),LT}, // LT = MINOR
		{Creer(2,3), Creer(1,2),MAJOR},
		{Creer(2,3), Creer(-1,2),MAJOR},

		{Creer(1,2), Creer(2,4),EQ}, // EQ = AEQUAL
		{Creer(-1,2), Creer(2,-4),AEQUAL},
		{Creer(2,4), Creer(1,2),AEQUAL},
		{Creer(2,-4), Creer(-1,2),AEQUAL},

		{Creer(1,0), Creer(2,0),AEQUAL},
		{Creer(-1,0), Creer(-2,0),AEQUAL},
		{Creer(2,0), Creer(1,0),AEQUAL},
		{Creer(-2,0), Creer(-1,0),EGAL}, // EGAL = AEQUAL
		{Creer(0,0), Creer(0,0),AEQUAL},

		{Creer(1,2), Creer(1,0),MINOR},
		{Creer(-1,2), Creer(1,0),PLUS_PETIT}, // PLUS_PETIT = MINOR
		{Creer(1,0), Creer(1,2),GT}, // GT = MAJOR
		{Creer(1,0), Creer(-1,2),PLUS_GRAND}, // PLUS_GRAND = MAJOR

		{Creer(1,0), Creer(-2,0),MAJOR},
		{Creer(-1,0), Creer(2,0),MINOR},
		{Creer(-2,0), Creer(1,0),MINOR},
		{Creer(2,0), Creer(-1,0),MAJOR},

		{Creer(1,2), Creer(0,0),MAJOR},
		{Creer(-1,2), Creer(0,0),MINOR},
		{Creer(2,3), Creer(0,0),MAJOR},
		{Creer(-2,3), Creer(0,0),MINOR},
		{Creer(0,0), Creer(1,2),MINOR},
		{Creer(0,0), Creer(-1,2),MAJOR},
		{Creer(0,0), Creer(2,3),MINOR},
		{Creer(0,0), Creer(-2,3),MAJOR},

		{Creer(1,0), Creer(0,0),MAJOR},
		{Creer(-1,0), Creer(0,0),MINOR},
		{Creer(0,0), Creer(1,0),MINOR},
		{Creer(0,0), Creer(-1,0),MAJOR},
	}
	for _, v := range test_fract_compare {
		attendu := v.c
		obtenu := v.a.Compare(v.b)
		if attendu != obtenu {
			t.Errorf(test+"(%v %v): attendu %v, obtenu %v", v.a, v.b, attendu, obtenu)
		}
	}
}

type test_fract_ope_t struct {
	a, b , c *Fract_t
}

func Test_fract_add(t *testing.T) {
	var test_fract_add = []test_fract_ope_t {
		{Creer(1,2), Creer(2,3), Creer(7,6)},
		{Creer(-1,2), Creer(2,3), Creer(1,6)},
		{Creer(2,3), Creer(1,2), Creer(7,6)},
		{Creer(2,3), Creer(-1,2), Creer(1,6)},

		{Creer(1,0), Creer(1,2), Creer(1,0)},
		{Creer(1,0), Creer(-1,2), Creer(1,0)},
		{Creer(1,2), Creer(1,0), Creer(1,0)},
		{Creer(-1,2), Creer(1,0), Creer(1,0)},

		{Creer(-1,0), Creer(1,2), Creer(-1,0)},
		{Creer(-1,0), Creer(-1,2), Creer(-1,0)},
		{Creer(1,2), Creer(-1,0), Creer(-1,0)},
		{Creer(-1,2), Creer(-1,0), Creer(-1,0)},

		{Creer(0,0), Creer(1,2), Creer(0,0)},
		{Creer(0,0), Creer(-1,2), Creer(0,0)},
		{Creer(1,2), Creer(0,0), Creer(0,0)},
		{Creer(-1,2), Creer(0,0), Creer(0,0)},
	}
	for _, v := range test_fract_add {
		attendu := v.c
		obtenu := v.a.Add(v.b)
		if !attendu.Egal(obtenu) {
			t.Errorf("add(%v %v): attendu %v, obtenu %v",v.a,v.b, attendu, obtenu)
		}
	}
}

func Test_fract_sub(t *testing.T) {
	var test_fract_sub = []test_fract_ope_t {
		{Creer(1,2), Creer(2,3), Creer(-1,6)},
		{Creer(-1,2), Creer(2,3), Creer(-7,6)},
		{Creer(2,3), Creer(1,2), Creer(1,6)},
		{Creer(2,3), Creer(-1,2), Creer(7,6)},

		{Creer(1,0), Creer(1,2), Creer(1,0)},
		{Creer(1,0), Creer(-1,2), Creer(1,0)},
		{Creer(1,2), Creer(1,0), Creer(-1,0)},
		{Creer(-1,2), Creer(1,0), Creer(-1,0)},

		{Creer(-1,0), Creer(1,2), Creer(-1,0)},
		{Creer(-1,0), Creer(-1,2), Creer(-1,0)},
		{Creer(1,2), Creer(-1,0), Creer(1,0)},
		{Creer(-1,2), Creer(-1,0), Creer(1,0)},

		{Creer(0,0), Creer(1,2), Creer(0,0)},
		{Creer(0,0), Creer(-1,2), Creer(0,0)},
		{Creer(1,2), Creer(0,0), Creer(0,0)},
		{Creer(-1,2), Creer(0,0), Creer(0,0)},
	}
	for _, v := range test_fract_sub {
		attendu := v.c
		obtenu := v.a.Sub(v.b)
		if !attendu.Egal(obtenu) {
			t.Errorf("sub(%v %v): attendu %v, obtenu %v",v.a,v.b, attendu, obtenu)
		}
	}
}

func Test_fract_mul(t *testing.T) {
	test := "mul"
	var test_fract_mul = []test_fract_ope_t {
		{Creer(1,2), Creer(2,3), Creer(1,3)},
		{Creer(-1,2), Creer(2,3), Creer(-1,3)},
		{Creer(2,3), Creer(1,2), Creer(1,3)},
		{Creer(2,3), Creer(-1,2), Creer(-1,3)},

		{Creer(1,0), Creer(1,2), Creer(1,0)},
		{Creer(1,0), Creer(-1,2), Creer(-1,0)},
		{Creer(1,2), Creer(1,0), Creer(1,0)},
		{Creer(-1,2), Creer(1,0), Creer(-1,0)},

		{Creer(-1,0), Creer(1,2), Creer(-1,0)},
		{Creer(-1,0), Creer(-1,2), Creer(1,0)},
		{Creer(1,2), Creer(-1,0), Creer(-1,0)},
		{Creer(-1,2), Creer(-1,0), Creer(1,0)},

		{Creer(0,0), Creer(1,2), Creer(0,0)},
		{Creer(0,0), Creer(-1,2), Creer(0,0)},
		{Creer(1,2), Creer(0,0), Creer(0,0)},
		{Creer(-1,2), Creer(0,0), Creer(0,0)},
	}
	for _, v := range test_fract_mul {
		attendu := v.c
		obtenu := v.a.Mul(v.b)
		if !attendu.Egal(obtenu) {
			t.Errorf(test+"(%v %v): attendu %v, obtenu %v",v.a,v.b, attendu, obtenu)
		}
	}
}

func Test_fract_pgcd(t *testing.T) {
	test := "pgcd"
	p := 12
	n, d := 51*p, 2*p
	{
		attendu := p
		obtenu := pgcd1(n,d)
		if attendu != obtenu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}
	{
		attendu := p
		obtenu := pgcd2(n,d)
		if attendu != obtenu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}
	{
		attendu := p
		obtenu := pgcd1(-n,d)
		if attendu != obtenu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}
	{
		attendu := p
		obtenu := pgcd2(-n,d)
		if attendu != obtenu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}
	{
		attendu := p
		obtenu := pgcd1(n,-d)
		if attendu != obtenu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}
	{
		attendu := p
		obtenu := pgcd2(n,-d)
		if attendu != obtenu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}
	{
		attendu := p
		obtenu := pgcd1(-n,-d)
		if attendu != obtenu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}
	{
		attendu := p
		obtenu := pgcd2(-n,-d)
		if attendu != obtenu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}

	{
		attendu := n
		obtenu := pgcd1(n,0)
		if attendu != obtenu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}
	{
		attendu := n
		obtenu := pgcd2(n,0)
		if attendu != obtenu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}

	{
		attendu := d
		obtenu := pgcd1(0,d)
		if attendu != obtenu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}
	{
		attendu := d
		obtenu := pgcd2(0,d)
		if attendu != obtenu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}

	{
		attendu := 1
		obtenu := pgcd1(0,0)
		if attendu != obtenu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}
	{
		attendu := 1
		obtenu := pgcd2(0,0)
		if attendu != obtenu {
			t.Errorf(test+" : attendu %v, obtenu %v", attendu, obtenu)
		}
	}
}

