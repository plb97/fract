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
func Test_fract_new(t *testing.T) {
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
		expected := v.f
		actual := New(v.n,v.d)
		if !expected.Egal(actual) {
			t.Errorf("new: expected %v, actual %v", expected, actual)
		}
	}
}

type test_fract_equal_t struct {
	a, b *Fract_t
	c bool
}
func Test_fract_equal(t *testing.T) {
	var test_fract_equal = []test_fract_equal_t {
		{New(1,2),New(2,3),false},
		{New(-1,2),New(2,3),false},
		{New(2,3),New(1,2),false},
		{New(2,3),New(-1,2),false},
		{New(1,2),New(1,0),false},
		{New(-1,2),New(1,0),false},
		{New(1,0),New(1,2),false},
		{New(1,0),New(-1,2),false},
		{New(1,2),New(2,4),true},
		{New(-1,2),New(2,-4),true},
		{New(2,4),New(1,2),true},
		{New(2,-4),New(-1,2),true},
		{New(1,0),New(2,0),true},
		{New(-1,0),New(-2,0),true},
		{New(2,0),New(1,0),true},
		{New(0,0),New(0,0),true},
		{New(-2,0),New(-1,0),true},
		{New(1,0),New(-2,0),false},
		{New(-1,0),New(2,0),false},
		{New(-2,0),New(1,0),false},
		{New(2,0),New(-1,0),false},
	}
	for _, v := range test_fract_equal {
		expected := v.c
		actual := v.a.Egal(v.b)
		if expected != actual {
			t.Errorf("equal: expected %v, actual %v", expected, actual)
		}
	}
}

type test_fract_compare_t struct {
	a, b *Fract_t
	c Fract_compare_t
}
func Test_fract_compare(t *testing.T) {
	var test_fract_compare = []test_fract_compare_t {
		{New(1,2),New(2,3),MINOR},
		{New(-1,2),New(2,3),MINOR},
		{New(2,3),New(1,2),MAJOR},
		{New(2,3),New(-1,2),MAJOR},

		{New(1,2),New(2,4),AEQUAL},
		{New(-1,2),New(2,-4),AEQUAL},
		{New(2,4),New(1,2),AEQUAL},
		{New(2,-4),New(-1,2),AEQUAL},

		{New(1,0),New(2,0),AEQUAL},
		{New(-1,0),New(-2,0),AEQUAL},
		{New(2,0),New(1,0),AEQUAL},
		{New(-2,0),New(-1,0),AEQUAL},
		{New(0,0),New(0,0),AEQUAL},

		{New(1,2),New(1,0),MINOR},
		{New(-1,2),New(1,0),MINOR},
		{New(1,0),New(1,2),MAJOR},
		{New(1,0),New(-1,2),MAJOR},

		{New(1,0),New(-2,0),MAJOR},
		{New(-1,0),New(2,0),MINOR},
		{New(-2,0),New(1,0),MINOR},
		{New(2,0),New(-1,0),MAJOR},

		{New(1,2),New(0,0),MAJOR},
		{New(-1,2),New(0,0),MINOR},
		{New(2,3),New(0,0),MAJOR},
		{New(-2,3),New(0,0),MINOR},
		{New(0,0),New(1,2),MINOR},
		{New(0,0),New(-1,2),MAJOR},
		{New(0,0),New(2,3),MINOR},
		{New(0,0),New(-2,3),MAJOR},

		{New(1,0),New(0,0),MAJOR},
		{New(-1,0),New(0,0),MINOR},
		{New(0,0),New(1,0),MINOR},
		{New(0,0),New(-1,0),MAJOR},
	}
	for _, v := range test_fract_compare {
		expected := v.c
		actual := v.a.Compare(v.b)
		if expected != actual {
			t.Errorf("compare(%v %v): expected %v, actual %v", v.a, v.b, expected, actual)
		}
	}
}

type test_fract_ope_t struct {
	a, b , c *Fract_t
}

func Test_fract_add(t *testing.T) {
	var test_fract_add = []test_fract_ope_t {
		{New(1,2),New(2,3),New(7,6)},
		{New(-1,2),New(2,3),New(1,6)},
		{New(2,3),New(1,2),New(7,6)},
		{New(2,3),New(-1,2),New(1,6)},

		{New(1,0),New(1,2),New(1,0)},
		{New(1,0),New(-1,2),New(1,0)},
		{New(1,2),New(1,0),New(1,0)},
		{New(-1,2),New(1,0),New(1,0)},

		{New(-1,0),New(1,2),New(-1,0)},
		{New(-1,0),New(-1,2),New(-1,0)},
		{New(1,2),New(-1,0),New(-1,0)},
		{New(-1,2),New(-1,0),New(-1,0)},

		{New(0,0),New(1,2),New(0,0)},
		{New(0,0),New(-1,2),New(0,0)},
		{New(1,2),New(0,0),New(0,0)},
		{New(-1,2),New(0,0),New(0,0)},
	}
	for _, v := range test_fract_add {
		expected := v.c
		actual := v.a.Add(v.b)
		if !expected.Egal(actual) {
			t.Errorf("add(%v %v): expected %v, actual %v",v.a,v.b, expected, actual)
		}
	}
}

func Test_fract_sub(t *testing.T) {
	var test_fract_sub = []test_fract_ope_t {
		{New(1,2),New(2,3),New(-1,6)},
		{New(-1,2),New(2,3),New(-7,6)},
		{New(2,3),New(1,2),New(1,6)},
		{New(2,3),New(-1,2),New(7,6)},

		{New(1,0),New(1,2),New(1,0)},
		{New(1,0),New(-1,2),New(1,0)},
		{New(1,2),New(1,0),New(-1,0)},
		{New(-1,2),New(1,0),New(-1,0)},

		{New(-1,0),New(1,2),New(-1,0)},
		{New(-1,0),New(-1,2),New(-1,0)},
		{New(1,2),New(-1,0),New(1,0)},
		{New(-1,2),New(-1,0),New(1,0)},

		{New(0,0),New(1,2),New(0,0)},
		{New(0,0),New(-1,2),New(0,0)},
		{New(1,2),New(0,0),New(0,0)},
		{New(-1,2),New(0,0),New(0,0)},
	}
	for _, v := range test_fract_sub {
		expected := v.c
		actual := v.a.Sub(v.b)
		if !expected.Egal(actual) {
			t.Errorf("sub(%v %v): expected %v, actual %v",v.a,v.b, expected, actual)
		}
	}
}

func Test_fract_mul(t *testing.T) {
	var test_fract_mul = []test_fract_ope_t {
		{New(1,2),New(2,3),New(1,3)},
		{New(-1,2),New(2,3),New(-1,3)},
		{New(2,3),New(1,2),New(1,3)},
		{New(2,3),New(-1,2),New(-1,3)},

		{New(1,0),New(1,2),New(1,0)},
		{New(1,0),New(-1,2),New(-1,0)},
		{New(1,2),New(1,0),New(1,0)},
		{New(-1,2),New(1,0),New(-1,0)},

		{New(-1,0),New(1,2),New(-1,0)},
		{New(-1,0),New(-1,2),New(1,0)},
		{New(1,2),New(-1,0),New(-1,0)},
		{New(-1,2),New(-1,0),New(1,0)},

		{New(0,0),New(1,2),New(0,0)},
		{New(0,0),New(-1,2),New(0,0)},
		{New(1,2),New(0,0),New(0,0)},
		{New(-1,2),New(0,0),New(0,0)},
	}
	for _, v := range test_fract_mul {
		expected := v.c
		actual := v.a.Mul(v.b)
		if !expected.Egal(actual) {
			t.Errorf("mul(%v %v): expected %v, actual %v",v.a,v.b, expected, actual)
		}
	}
}
