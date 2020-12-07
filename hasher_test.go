package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const EXPECTED_URL string = "some-url"
const EXPECTED_HASH_STR string = "2788029805"

func TestHasher_GenerateHashFromURL(t *testing.T) {
	asrt := assert.New(t)
	hasher := new(Hasher)
	asrt.Equal(EXPECTED_HASH_STR, hasher.GenerateHashFromURL(EXPECTED_URL))
}