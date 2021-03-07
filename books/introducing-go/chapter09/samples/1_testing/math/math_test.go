package math

import "testing"

func TestAverage(t *testing.T) {
	v := Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

func TestMin(t *testing.T) {
	v := Min([]float64{12, 1, 5})
	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}

func TestMax(t *testing.T) {
	v := Max([]float64{3, 12, 1, 5})
	if v != 12 {
		t.Error("Expected 12, got ", v)
	}
}

type testpair struct {
	values []float64
	result float64
}

var testsAverage = []testpair{
	{[]float64{},0},
	{[]float64{1,2}, 1.5 },
	{[]float64{1,1,1,1,1,1}, 1 },
	{[]float64{-1,1}, 0 },
}

var testsMin = []testpair{
	{[]float64{1, 2, 3, 4, 5, 6, 7}, 1},
	{[]float64{7, 6, 5, 4, 3, 2, 1}, 1},
	{[]float64{5, 3, 2, 1, 6, 7, 4}, 1},
}

var testsMax = []testpair{
	{[]float64{1, 2, 3, 4, 5, 6, 7}, 7},
	{[]float64{7, 6, 5, 4, 3, 2, 1}, 7},
	{[]float64{5, 3, 2, 1, 6, 7, 4}, 7},
}

func TestMathPackage(t *testing.T) {
	testMathFunc(t, testsAverage, Average)
	testMathFunc(t, testsMin, Min)
	testMathFunc(t, testsMax, Max)
}

func testMathFunc(t *testing.T, testSet []testpair, f func([]float64) float64)  {
	for _, pair := range testSet {
		v := f(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
