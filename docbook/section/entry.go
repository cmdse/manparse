package section

import (
	"encoding/xml"
	"fmt"
)

type ManEntry struct {
	XMLName   xml.Name `xml:"citerefentry"`
	Title     string   `xml:"refentrytitle"`
	ManVolume string   `xml:"manvolnum"`
}

func (entry ManEntry) String() string {
	return fmt.Sprintf("%v(%v)", entry.Title, entry.ManVolume)
}
