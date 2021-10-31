package main

import (
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

func TestFindPatternMatchesInStringFound(t *testing.T) {
	actual := findPatternMatchesInString("Match numbers ([0-9]+)","Match numbers 1234")

	assert.Contains(t,actual,"1234")
}

func TestFindPatternMatchesInStringNotFound(t *testing.T) {
	actual := findPatternMatchesInString("Match numbers ([0-9]+)","Match numbers no numbers here")

	assert.Empty(t,actual)
}

func TestFindPatternMatchesInStringTitleAuthorFound(t *testing.T) {
	actual := findPatternMatchesInString(`(.*?)\s*\((.*?)\)`, "This is a book title   (Author Name)")
	
	assert.Contains(t,actual,"This is a book title", "Author Name")
}

func TestContainsElement(t *testing.T) {
	aSlice := []string{"test"}

	assert.True(t, contains(aSlice,"test"))
}

func TestNotContainsElement(t *testing.T) {
	aSlice := []string{"test"}

	assert.False(t, contains(aSlice,"notThere"))
}
