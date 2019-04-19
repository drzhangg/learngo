package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{a,b,c int}{
		{3,4,5},
		{5,12,13},
		{8,15,17},
		{12,35,37},
		{30000,40000,50000},
	}

	for _,tt := range tests{
		if actual := calcTriangle(tt.a,tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d %d);  got %d; expected %d",tt.a,tt.b,actual,tt.c)
		}
	}
}

func BenchmarkTriangle(b *testing.B) {
	a,bb := 3,4
	r := 5
	for i := 0; i < b.N; i++ {
		actual := calcTriangle(a,bb)
		if actual != r {
			b.Errorf("calcTriangle(%d %d);  got %d; expected %d",a,bb,actual,r)
		}
	}
}
