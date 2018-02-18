package shared

import "encoding/xml"

type TagDecoder interface {
	DecodeTags(d *xml.Decoder, start xml.StartElement) error
}

func UnmarshalFromDecoder(d *xml.Decoder, decoder TagDecoder) (err error) {
	token, _ := d.Token()
	for token != nil && err == nil {
		if ttype, ok := token.(xml.StartElement); ok {
			err = decoder.DecodeTags(d, ttype)
		}
		token, _ = d.Token()
	}
	return err
}
