package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExistsTrue(t *testing.T) {
	actual := fileExists("sample_clippings.txt")

    assert.True(t,actual)
}

func TestFileExistsFalse(t *testing.T) {
	actual := fileExists("notexists.txt")

    assert.False(t,actual)
}

func TestMatchByRegexPatternFound(t *testing.T) {
	actual := matchByRegex("Match numbers ([0-9]+)","Match numbers 1234")

	assert.Contains(t,actual,"1234")
}

func TestMatchByRegexPatternNotFound(t *testing.T) {
	actual := matchByRegex("Match numbers ([0-9]+)","Match numbers no numbers here")

	assert.Empty(t,actual)
}

func TestContainsElement(t *testing.T) {
	aSlice := []string{"test"}

	assert.True(t, contains(aSlice,"test"))
}

func TestNotContainsElement(t *testing.T) {
	aSlice := []string{"test"}

	assert.False(t, contains(aSlice,"notThere"))
}

func TestIsError(t *testing.T) {
	
	
}