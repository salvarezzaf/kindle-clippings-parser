package main

import (
	"testing"
)

func TestParseClippingSuccess(t *testing.T) {

	clipping := New("clippings.txt")

	clipping.Parse()

}
