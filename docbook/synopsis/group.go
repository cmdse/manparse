package synopsis

import (
	"encoding/xml"
	"fmt"

	"github.com/cmdse/manparse/docbook/shared"
)

// http://tdg.docbook.org/tdg/4.5/group.html
type Group struct {
	XMLName xml.Name `xml:"group"`
	CmdTokenAttributes
	CmdSynopsisTokenContainer
}

func (group *Group) DecodeTags(d *xml.Decoder, start xml.StartElement) (err error) {
	inElement := start.Name.Local
	switch inElement {
	case "arg":
		err = group.handleArg(d, start)
	case "group":
		err = group.handleGroup(d, start)
	}
	return err
}

func (group *Group) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	group.Tokens = make([]SynopsisFragment, 0, 20)
	return shared.UnmarshalFromDecoder(d, group)
}

func (group *Group) String() string {
	tokenFmt := "\n"
	for _, token := range group.Tokens {
		tokenFmt = fmt.Sprintf("%v\t\t%v", tokenFmt, token)
	}
	return fmt.Sprintf("Group{\n\tchoice: %v,\n\trepeat: %v,\n\tchildren: %v \n\t}\n", group.Choice, group.Repeat, tokenFmt)
}
