package regex

import (
	"regexp"
	"testing"
)

type testCase struct {
	name  string
	regex string
	input string
}

func TestFSMWithGoStdLibCharacterLiterals(t *testing.T) {
	nestedExpRegex := "a(b(c))"

	testsCases := []testCase{
		{"nested expressions", nestedExpRegex, "abcd"},
	}

	for _, tc := range testsCases {
		t.Run(
			tc.name, func(t *testing.T) {
				compareWithGoStdLib(t, NewRegex(tc.regex), tc.regex, tc.input)
			},
		)
	}
}

func TestCharacterLiteralRegex(t *testing.T) {
	abcRegex := "abc"

	testsCases := []testCase{
		{"empty string", abcRegex, ""},
		{"non matching string", abcRegex, "xxx"},
		{"matching string", abcRegex, "abc"},
		{"partial matching string", abcRegex, "ab"},
		{"nested expressions", "a(b(c))", "abcd"},
	}

	for _, tc := range testsCases {
		t.Run(
			tc.name, func(t *testing.T) {
				compareWithGoStdLib(t, NewRegex(tc.regex), tc.regex, tc.input)
			},
		)
	}
}

func FuzzTesting(f *testing.F) {
	abcRegex := "abc"
	f.Add(abcRegex, "abc")
	f.Add(abcRegex, "abcs")
	f.Add(abcRegex, "tdd")
	f.Add(abcRegex, "")

	nestedExpRegex := "c(a(b))"
	f.Add(nestedExpRegex, "cca")
	f.Add(nestedExpRegex, "cab")
	f.Add(nestedExpRegex, "cab")

	f.Fuzz(func(t *testing.T, regex, input string) {
		if _, err := regexp.Compile(regex); err != nil {
			t.Skip()
		}

		compareWithGoStdLib(t, NewRegex(regex), regex, input)
	})

}

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
