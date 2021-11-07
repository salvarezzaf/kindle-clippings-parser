# Kindle Highlights/Notes Parser
This is a simple GO module which is able to parse highlights or notes stored in "MyClippings.txt" on your Kindle device.

```bash
go get github.com/salvarezzaf/kindle-clippings-parser
```

MyClippings.txt is usually located in the root folder mounted by your OS when attaching your Kindle device via USB. 

```go
parser := NewClippingsParser(clippingsFilePath string)
clippings := parser.Parse()
```

 The Parse function will return clippings grouped by book title. Parsed clipping metadata is stored in a Clipping struct which holds the following:

```go
type Clipping struct {
 	title        string
	author       string
	clippingType string // Highlight or Note
	pageOrLoc    string // either a page or loc if ebook did not have pages
	loc          string // loc added if ebook has page and loc details
	clippingDate string  
	content      string
}
```

 Final output from parsing would be a map grouping clipppings by book title

```go
map[string][]Clipping { 
 "book title1" : []Clipping{Clipping{},Clipping{}},
 "book title2" : []Clipping{Clipping{},Clipping{},Clipping{}},
}
```

