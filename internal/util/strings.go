package util

import "strings"

func CaseInsensitiveSubstring(s, substr string) bool {
	// Convert both the original string and substring to lowercase
	lowercaseS := strings.ToLower(s)
	lowercaseSubstr := strings.ToLower(substr)

	// Check if the lowercase substring exists in the lowercase original string
	return strings.Contains(lowercaseS, lowercaseSubstr)
}

func ConvertStringAmountToFloat(input string) string {
	input = strings.ReplaceAll(input, ",", "")
	input = strings.ReplaceAll(input, "$", "")
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "+", "")

	return input
}
