package parser

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClippingParserTestSuite struct {
	suite.Suite
	expectedClippings map[string][]Clipping
}


func (suite *ClippingParserTestSuite) SetupSuite()  {
	suite.expectedClippings = initExpectedClippings()
}

func (suite *ClippingParserTestSuite) TestParseClippingSuccess() {

	clipping := NewClippingsParser("sample_clippings.txt")

	clippings, _ := clipping.Parse()

	assert.True(suite.T(), len(clippings) == 2)
	assert.ObjectsAreEqualValues(clippings,suite.expectedClippings)

}

func (suite *ClippingParserTestSuite) TestParseClippingError() {

	clipping := NewClippingsParser("not_exists.txt")

	_, err := clipping.Parse()

	assert.EqualError(suite.T(),err,"Clipping file not found in provided file path")

}

func TestClippingParserTestSuite(t *testing.T) {
	suite.Run(t,new(ClippingParserTestSuite))
}

func initExpectedClippings() map[string][]Clipping {
	clipping1 := Clipping{
		title: "How to Win Friends and Influence People", 
		author: "Dale Carnegie", 
		clippingType: "Highlight", 
		pageOrLoc: "9", 
		loc: "132-134",
		clippingDate: "Sunday 5 September 2021 17:47:10",
		content: "highlight 1",
	}

	clipping2 := Clipping{
		title: "How to Win Friends and Influence People", 
		author: "Dale Carnegie", 
		clippingType: "Highlight", 
		pageOrLoc: "9", 
		loc: "138-140",
		clippingDate: "Sunday 5 September 2021 17:48:21",
		content: "highlight 2",
	}

	clipping3 := Clipping{
		title: "Investing In ETF For Dummies", 
		author: "Russell Wild", 
		clippingType: "Note", 
		pageOrLoc: "93", 
		loc: "1413",
		clippingDate: "Sunday 3 October 2021 22:24:26",
		content: "note",
	}
	
    clippings := make(map[string][]Clipping)
	clippings[clipping1.title] = []Clipping{clipping1,clipping2}
	clippings[clipping3.title] = []Clipping{clipping3}

	return clippings
}
