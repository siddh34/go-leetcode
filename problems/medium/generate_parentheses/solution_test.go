package generate_parentheses

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	expected := []string{"()"}
	result := generateParenthesis(1)
	if len(result) != len(expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	}
}