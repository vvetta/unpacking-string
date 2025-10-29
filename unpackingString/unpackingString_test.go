package unpackingstring_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	unpackingstring "github.com/vvetta/unpacking-string/unpackingString"
)

func TestUnpackWoEscape(t *testing.T) {
	testStrings := []string{"a4bc2d5e", "abcd", "45", "", "a0", "a1", "a2", "a22"}
	results := []string{"aaaabccddddde", "abcd", "", "", "a", "a", "aa", ""}

	for i, testString := range testStrings {
		result, _ := unpackingstring.Unpack(testString)
		if results[i] != result {
			t.Errorf("test case: %s error; correct(%s) != result(%s)", testStrings[i], results[i], result)
		} else {
			t.Logf("test case: %s PASS", testStrings[i])
		}
	}
}

func TestUnpack_twoNumInRow(t *testing.T) {
	testStrings := []string{"45554343", "a22"}
	results := []string{"", ""}

	for i, testString := range testStrings {
		result, err := unpackingstring.Unpack(testString)
		require.Equal(t, unpackingstring.ErrTwoNumInRow, err)
		if result != results[i] {
			t.Errorf("test case: %s error; correct(%s) != result(%s)", testStrings[i], results[i], result)
		} else {
			t.Logf("test case: %s PASS", testStrings[i])
		}
	}
}

func TestUnpackWEscape(t *testing.T) {
	testStrings := []string{`qwe\4\5`, `qwe\45`, `\45`, `\4\4\5\6\7\8`}
	results := []string{"qwe45", "qwe44444", "44444", "445678"}

	for i, testString := range testStrings {
		result, _ := unpackingstring.Unpack(testString)
		if results[i] != result {
			t.Errorf("test case: %s error; correct(%s) != result(%s)", testStrings[i], results[i], result)
		} else {
			t.Logf("test case: %s PASS", testStrings[i])
		}
	}
}
