package regex

import (
	"regexp"
	"testing"
)

func compareWithGoStdLib(t *testing.T, regex *regex, regexStr, input string) {
	t.Helper()
	actualResult := regex.MatchString(input)
	expectedResult := regexp.MustCompile(regexStr).MatchString(input)

	if actualResult != expectedResult {
		t.Fatalf(
			"Mismatch on input %s, bytes: %x\nregex: %s, bytes:%x\ngo regexp package result: '%t'\nthis package result: '%v'",
			input, []byte(input), regexStr, []byte(regexStr), expectedResult, actualResult,
		)
	}
}
