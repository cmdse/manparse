package synopsis

import (
	"encoding/xml"
	"errors"
)

// Choice indicates whether the Arg is required (Req or Plain) or optional (Opt).
// Arguments identified as Plain are required, but are shown without additional decoration.
type ChoiceType uint8

const (
	ChoiceOptional ChoiceType = iota
	ChoicePlain
	ChoiceRequired
)

func ChoiceFromAttr(attr xml.Attr) (choice ChoiceType, err error) {
	switch attr.Value {
	case "opt", "":
		choice = ChoiceOptional
	case "plain":
		choice = ChoicePlain
	case "req":
		choice = ChoiceRequired
	default:
		err = errors.New("unexpected value for 'choice' attribute")
	}
	return choice, err
}

func (choice *ChoiceType) UnmarshalXMLAttr(attr xml.Attr) error {
	inferredChoice, err := ChoiceFromAttr(attr)
	*choice = inferredChoice
	return err
}

func (choice ChoiceType) String() string {
	switch choice {
	case ChoiceOptional:
		return "optional"
	case ChoicePlain:
		return "plain"
	case ChoiceRequired:
		return "required"
	default:
		return ""
	}
}
