package synopsis

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

type RepeatAttr bool

func (repeat *RepeatAttr) UnmarshalXMLAttr(attr xml.Attr) (err error) {
	switch attr.Value {
	case "repeat":
		*repeat = true
	case "norepeat", "":
		*repeat = false
	default:
		err = errors.New("'rep' attribute expects 'repeat' or 'norepeat' value.")
	}
	return err
}
