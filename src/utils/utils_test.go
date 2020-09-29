package utils_test

import (
	"testing"

	. "github.com/OtavioGallego/banking-api/src/utils"
)

var slice = []int{1, 2, 3, 4, 5}

type testScenario struct {
	slice         []int
	input         int
	expectedFound bool
}

var testScenarios = []testScenario{
	{slice, 1, true},
	{slice, 2, true},
	{slice, 3, true},
	{slice, 4, true},
	{slice, 5, true},
	{slice, 6, false},
	{slice, 7, false},
	{slice, 8, false},
	{slice, 9, false},
	{slice, 10, false},
}

func TestFind(t *testing.T) {
	for _, scenario := range testScenarios {
		found := Find(scenario.slice, scenario.input)

		if found != scenario.expectedFound {
			t.Errorf("Expected found to be %t but got %t", scenario.expectedFound, found)
		}
	}
}
