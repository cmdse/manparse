package main

import (
	"fmt"
	"os"

	"github.com/cmdse/manparse/pim"

	"github.com/cmdse/manparse/docbook"
)

func checkArgs() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "missing xml file to extract model")
		os.Exit(1)
	}
}

func checkUnmarshalErr(err error, file string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not extract information from file '%v': %v", file, err)
		os.Exit(1)
	}
}

func main() {
	var err error
	var mandoc *docbook.ManDocBookXml
	checkArgs()
	file := os.Args[1]
	mandoc, err = docbook.Unmarshal(file)
	checkUnmarshalErr(err, file)
	fmt.Printf("COMMAND: %#v\n", mandoc.CmdSynopsis.Command)
	pim.ExtractPIMFromDocBook(mandoc, os.Stdout)
}
