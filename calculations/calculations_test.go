package calculations

import (
	"testing"
)

type test struct {
	exp       string
	answer    float64
	shouldErr bool
}

/*
Tests to add:
	Negatives
	Invalid operators
*/

var tests = [...]test{
	//Valid
	{"1", 1.0, false},   //Parsing
	{"1+1", 2.0, false}, //Basic operators
	{"1-1", 0.0, false},
	{"1*4", 4.0, false},
	{"1/4", 0.25, false},
	{"2^3", 8.0, false},
	{"1+1", 2.0, false},
	{"(1+1)*2", 4.0, false},
	{"1.5+1.5", 3.0, false}, //Floats
	{"1.+1", 2.0, false},
	{"1+3*2/6+2^2", 6.0, false}, //Order of operations
	{"(1+3)*2/4+2^2", 6.0, false},

	//Invalid
	{"1.0.+1", 0.0, true}, //Extra .
	{"1..0+1", 0.0, true},
	{"1++1", 0.0, true}, //Extra operators
	{"1+", 0.0, true},   //Not enough operands
	{"+", 0.0, true},
}

func TestCalculation(t *testing.T) {
	for _, currTest := range tests {
		doCalcTest(t, currTest.exp, currTest.answer, currTest.shouldErr)
	}
}

func doCalcTest(t *testing.T, expression string, expectedAnswer float64, expectedErr bool) {
	ans, err := Evaluate(expression)

	if err && !expectedErr {
		t.Fatalf("Expression '%s' failed but should not have.", expression)
		return
	}

	if !err && expectedErr {
		t.Fatalf("Expression '%s' did not fail but was expected to fail.", expression)
		return
	}

	if ans != expectedAnswer {
		t.Fatalf("Expression '%s' got answer '%f' but expected '%f'.", expression, ans, expectedAnswer)
	}
}
