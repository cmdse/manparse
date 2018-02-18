package synopsis

import (
	"encoding/xml"

	"github.com/cmdse/manparse/docbook/shared"
)

// http://tdg.docbook.org/tdg/4.5/cmdsynopsis.html
type CmdSynopsis struct {
	CmdSynopsisTokenContainer
	XMLName xml.Name `xml:"cmdsynopsis"`
	Command string
}

func (synopsis *CmdSynopsis) handleCmd(d *xml.Decoder, start xml.StartElement) (err error) {
	container := &struct {
		XMLName xml.Name `xml:"command"`
		Command string   `xml:",chardata"`
	}{}
	err = d.DecodeElement(container, &start)
	synopsis.Command = container.Command
	return err
}

func (synopsis *CmdSynopsis) DecodeTags(d *xml.Decoder, start xml.StartElement) (err error) {
	inElement := start.Name.Local
	switch inElement {
	case "arg":
		err = synopsis.handleArg(d, start)
	case "group":
		err = synopsis.handleGroup(d, start)
	case "command":
		err = synopsis.handleCmd(d, start)
	}
	return err
}

func (synopsis *CmdSynopsis) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	synopsis.Tokens = make([]SynopsisFragment, 0, 20)
	return shared.UnmarshalFromDecoder(d, synopsis)
}
