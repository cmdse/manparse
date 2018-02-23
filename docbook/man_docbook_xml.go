package docbook

import (
	"github.com/cmdse/manparse/docbook/section"
	"github.com/cmdse/manparse/docbook/synopsis"
)

// http://tdg.docbook.org/tdg/4.5/cmdsynopsis.html
type ManDocBookXml struct {
	Name        string               `xml:"refnamediv>refname"`
	Purpose     string               `xml:"refnamediv>refpurpose"`
	CmdSynopsis synopsis.CmdSynopsis `xml:"refsynopsisdiv>cmdsynopsis"`
	Sections    []*section.Section   `xml:"refsect1"`
}
