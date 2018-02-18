package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/cmdse/manparse/docbook"
	"golang.org/x/net/html/charset"
)

func makeReader(name string) io.Reader {
	reader, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return reader
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "missing xml file to extract model")
		os.Exit(1)
	}
	file := os.Args[1]
	mandoc := docbook.ManDocBookXml{}
	decoder := xml.NewDecoder(makeReader(file))
	decoder.CharsetReader = charset.NewReaderLabel
	decoder.Strict = true
	xml.HTMLEntity["bsol"] = "\\"
	xml.HTMLEntity["sol"] = "/"
	decoder.Entity = xml.HTMLEntity
	err := decoder.Decode(&mandoc)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("COMMAND: %#v\n", mandoc.CmdSynopsis.Command)
	//fmt.Printf("TOKENS: %v\n", mandoc.CmdSynopsis.Tokens)

	for _, sec := range mandoc.Sections {
		if sec.Title == "OPTIONS" {
			fmt.Printf("%v\n", sec)
		}
	}
}
