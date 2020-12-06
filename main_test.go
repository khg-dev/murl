package main

import (
	"fmt"
	"testing"
)

func testFormatUint(i uint64, base int) string {
	return "abcde"
}
func TestGenerateShortUrl(t *testing.T) {
	old := strconvFormatUint
	defer func() { strconvFormatUint = old }()

	strconvFormatUint = func(i uint64, base int) string {
		t.Fail()
		return "abcde"
	}

	short := generateShortUrl("test")
	fmt.Println(short)

}
