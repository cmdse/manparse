package docbook

import (
	"encoding/xml"
	"io"
	"os"

	"golang.org/x/net/html/charset"
)

func init() {
	xml.HTMLEntity["bsol"] = "\\"
	xml.HTMLEntity["sol"] = "/"
}

func makeReader(name string) (io.Reader, error) {
	reader, err := os.Open(name)
	return reader, err
}

func makeDecoder(reader io.Reader) *xml.Decoder {
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoder.Strict = true
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
