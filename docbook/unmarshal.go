package docbook

import (
	"encoding/xml"
	"io"
	"os"

	"golang.org/x/net/html/charset"
)

func makeReader(name string) (io.Reader, error) {
	reader, err := os.Open(name)
	return reader, err
}

func makeDecoder(reader io.Reader) *xml.Decoder {
	xml.HTMLEntity["bsol"] = "\u005C"
	xml.HTMLEntity["sol"] = "\u002F"
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoder.Strict = false
	decoder.Entity = xml.HTMLEntity
	return decoder
}

func decode(decoder *xml.Decoder) (mandoc *ManDocBookXml, err error) {
	mandoc = &ManDocBookXml{}
	err = decoder.Decode(mandoc)
	return mandoc, err
}

func Unmarshal(path string) (mandoc *ManDocBookXml, err error) {
	var reader io.Reader
	reader, err = makeReader(path)
	if err != nil {
		return nil, err
	}
	decoder := makeDecoder(reader)
	mandoc, err = decode(decoder)
	return mandoc, err
}
