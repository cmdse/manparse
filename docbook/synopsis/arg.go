package synopsis

import (
	"encoding/xml"
	"fmt"
)

type CmdTokenAttributes struct {
	// opt, plain, req
	Choice ChoiceType `xml:"choice,attr"`
	// repeat, noreapeat
	Repeat RepeatAttr `xml:"rep,attr"`
}

// http://tdg.docbook.org/tdg/4.5/arg.html
type Arg struct {
	CmdTokenAttributes
	XMLName  xml.Name `xml:"arg"`
	Literal  string   `xml:",chardata"`
	Variable string   `xml:"replaceable,omitempty"`
}

func (arg Arg) String() string {
	return fmt.Sprintf("Arg{ choice: %v, literal: '%v', variable: '%v', repeat: %v }\n", arg.Choice, arg.Literal, arg.Variable, arg.Repeat)
}
