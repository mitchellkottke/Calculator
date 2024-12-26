package calculations

import (
	"testing"
)

type test struct {
	exp       string
	answer    float64
	shouldErr bool
}

var tests = [...]test{
	{"1+1", 2.0, false},
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
