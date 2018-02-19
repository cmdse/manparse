package section

import (
	"encoding/xml"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Element interface{}

type Nodes []*Node

func (elements *Nodes) Append(node *Node) {
	*elements = append(*elements, node)
}

// http://tdg.docbook.org/tdg/4.5/refsect1.html
type Section struct {
	XMLName  xml.Name `xml:"refsect1"`
	Children Nodes    `xml:",any"`
	RefSection
}

func (section Section) String() string {
	var rows = make([]string, 0, 20)
	for _, ch := range section.Children {
		rows = append(rows, ch.Flatten())
	}
	return spew.Sprintf("'%v'\n%v\n\n\n", section.Title, strings.Join(rows, "\n"))
}
