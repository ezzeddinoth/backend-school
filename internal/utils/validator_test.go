package utils

import "testing"

func TestValidateCommaSeparatedNumberListOrEmpty(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"1,2,3", true},
		{"10,20,30", true},
		{"1,a,3", false},
		{"1,,2", false},
		{"abc", false},
	}

	for _, test := range tests {
		if result := ValidateCommaSeparatedNumberListOrEmpty(test.input); result != test.expected {
			t.Errorf("ValidateCommaSeparatedNumberListOrEmpty(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestValidateUTM(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"32 N 461344 5481745", true},
		{"1 N 461344 5481745", true},
		{"60 N 461344 5481745", true},
		{"61 N 461344 5481745", false},
		{"32 I 461344 5481745", false},
		{"32 O 461344 5481745", false},
		{"", false},
		{"32 N 4613449 548174", false},
		{"32 N 461344 548174", false},
		{"32 461344 548174", false},
		{"322 N 461344 548174", false},
		{"61 N 461344 548174", false},
		{"32 N 46131 548174", false},
		{"32 N 461344 54817", false},
	}

	for _, test := range tests {
		if result := ValidateUTM(test.input); result != test.expected {
			t.Errorf("ValidateUTM(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestValidateStringOfDigits(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123456", true},
		{"00000", true},
		{"987654321", true},
		{"123abc", false},
		{"12 34", false},
		{"", false},
	}

	for _, test := range tests {
		if result := ValidateStringOfDigits(test.input); result != test.expected {
			t.Errorf("ValidateStringOfDigits(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestValidateHexColor(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"#123456", true},
		{"#000000", true},
		{"#FFFFFF", true},
		{"#aaaaaa", true},
		{"#aaaFFF", true},
		{"#aaaff", false},
		{"#aaaffff", false},
		{"123abc", false},
		{"12 34", false},
	}

	for _, test := range tests {
		if result := ValidateHexColor(test.input); result != test.expected {
			t.Errorf("ValidateHexColor(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}
