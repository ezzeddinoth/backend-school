package utils

import "regexp"

var (
	commaSeparatedNumberListOrEmptyRegex = *regexp.MustCompile(`^$|^(?:\d+(?:,\d+)*|\*)$`)
	listRangeNumberOrJoker               = regexp.MustCompile(`^(\d+(?:[,-]\d+)*|(?:[*]))$`)
	utmPattern                           = *regexp.MustCompile(`^([1-9]|[1-5][0-9]|60)[ ][C-HJ-NP-X][ ]([0-9]{6})[ ]([0-9]{7})$`)
	digitsPattern                        = *regexp.MustCompile(`^\d+$`)
	hexColorPattern                      = *regexp.MustCompile(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`)
)

func ValidateCommaSeparatedNumberListOrEmpty(str string) bool {
	return commaSeparatedNumberListOrEmptyRegex.MatchString(str)
}

func ValidateListRangeNumberOrJoker(str string) bool {
	return listRangeNumberOrJoker.MatchString(str)
}

func ValidateUTM(utm string) bool {
	return utmPattern.MatchString(utm)
}

func ValidateStringOfDigits(str string) bool {
	return digitsPattern.MatchString(str)
}

func ValidateHexColor(strn string) bool {
	return hexColorPattern.MatchString(strn)
}
