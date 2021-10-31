package main

import (
	"testing"
)

func TestParseClippingSuccess(t *testing.T) {

	clipping := New("sample_clippings.txt")

	clipping.Parse()

}
