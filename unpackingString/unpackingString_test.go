package unpackingstring_test

import (
	"testing"

	unpackingstring "github.com/vvetta/unpacking-string/unpackingString"
)

func TestUnpackWoEscape(t *testing.T) {
	testStrings := []string{"a4bc2d5e", "abcd", "45", "", "a0", "a1", "a2"}
	results := []string{"aaaabccddddde", "abcd", "", "", "a", "a", "aa"}

	for i, testString := range testStrings {
		result, _ := unpackingstring.Unpack(testString)
		if results[i] != result {
			t.Errorf("test case: %s error; correct(%s) != result(%s)", testStrings[i], results[i], result)
		}
		t.Logf("test case: %s PASS", testStrings[i])
	}
}

func TestUnpackWEscape(t *testing.T) {
	testStrings := []string{`qwe\4\5`, `qwe\45`, `\45`, `\4\4\5\6\7\8`}
	results := []string{"qwe45", "qwe44444", "44444", "445678"}

	for i, testString := range testStrings {
		result, _ := unpackingstring.Unpack(testString)
		if results[i] != result {
			t.Errorf("test case: %s error; correct(%s) != result(%s)", testStrings[i], results[i], result)
		}
		t.Logf("test case: %s PASS", testStrings[i])
	}
}
