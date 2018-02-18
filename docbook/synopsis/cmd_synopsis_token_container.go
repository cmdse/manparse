package synopsis

import "encoding/xml"

type CmdSynopsisTokenContainer struct {
	Tokens []SynopsisFragment
}

func (basis *CmdSynopsisTokenContainer) unmarshalTokenAndAppend(d *xml.Decoder, start xml.StartElement, cmdToken SynopsisFragment) error {
	err := d.DecodeElement(cmdToken, &start)
	basis.Tokens = append(basis.Tokens, cmdToken)
	return err
}

func (basis *CmdSynopsisTokenContainer) handleGroup(d *xml.Decoder, start xml.StartElement) error {
	return basis.unmarshalTokenAndAppend(d, start, &Group{})
}

func (basis *CmdSynopsisTokenContainer) handleArg(d *xml.Decoder, start xml.StartElement) error {
	return basis.unmarshalTokenAndAppend(d, start, &Arg{})
}
