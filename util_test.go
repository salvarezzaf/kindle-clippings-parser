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

func TestFindPatternMatchesInStringTitleAuthorFound(t *testing.T) {
	actual := findPatternMatchesInString(titleAuthorRegexPattern, "This is a book title   (Author Name)")
	
	assert.ElementsMatch(t,actual,[]string{"This is a book title", "Author Name"})
}

func TestFindPatternMatchesInStringClippingTypeAndPageFound(t *testing.T) {
 actual := findPatternMatchesInString(clippingTypePageOrLocRegexPattern,"Your Highlight on page 154")
 assert.ElementsMatch(t,[]string{"Highlight","on","page","154"},actual)
} 

func TestFindPatternMatchesInStringClippingTypeAndMultiplePagesFound(t *testing.T) {
	actual := findPatternMatchesInString(clippingTypePageOrLocRegexPattern,"Your Highlight on page 154-155")
	assert.ElementsMatch(t,[]string{"Highlight","on","page","154-155"},actual)
   } 

func TestFindPatternMatchesInStringClippingTypeAndLocFound(t *testing.T) {
	actual := findPatternMatchesInString(clippingTypePageOrLocRegexPattern,"Your Highlight at location 1200")
	assert.ElementsMatch(t,[]string{"Highlight","at","location","1200"},actual)
}

func TestFindPatternMatchesInStringClippingTypeAndMultipleLocsFound(t *testing.T) {
	actual := findPatternMatchesInString(clippingTypePageOrLocRegexPattern,"Your Highlight at location 1200-1300")
	assert.ElementsMatch(t,[]string{"Highlight","at","location","1200-1300"},actual)
}

func TestFindPatternMatchesInStringClippingLocOnlyFound(t *testing.T) {
	actual := findPatternMatchesInString(clippingLocRegexPattern,"| location 1200-1300")
	assert.ElementsMatch(t,[]string{"1200-1300"},actual)
}

func TestFindPatternMatchesInStringClippingDateAndTimeFound(t *testing.T) {
	actual := findPatternMatchesInString(clippingDateTimeRegexPattern,"| Added on Sunday, 3 October 2021 22:24:26")
	assert.ElementsMatch(t,[]string{"Sunday","3","October","2021","22:24:26"},actual)
}

func TestContainsElement(t *testing.T) {
	aSlice := []string{"test"}

	assert.True(t, contains(aSlice,"test"))
}

func TestNotContainsElement(t *testing.T) {
	aSlice := []string{"test"}

	assert.False(t, contains(aSlice,"notThere"))
}

func TestRemoveUnicodeSpecials(t *testing.T) {
	assert.True(t,IsUnicodeSpecial('\ufeff'))
}