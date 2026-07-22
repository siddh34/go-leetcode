package maxactivesectionsaftertrade

import "testing"

func TestMaxActiveSectionsAfterTrade(t *testing.T) {
	testCases := []struct {
		sections string
		expected int
	}{
		{sections: "01", expected: 1},
		{sections: "0100", expected: 4},
		{sections: "1000100", expected: 7},
		{sections: "01010", expected: 4},
	}
	
	for _, tc := range testCases {
		result := maxActiveSectionsAfterTrade(tc.sections)
		if result != tc.expected {
			t.Errorf("For sections %s, expected %d but got %d", tc.sections, tc.expected, result)
		}
	}
}
