package section

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
)

var whiteSpaces = regexp.MustCompile(`^\s+$/m`)

type Node struct {
	Type     string  `xml:"-"`
	Element  Element `xml:"-"`
	Children Nodes   `xml:",any"`
}

func (node *Node) Flatten() string {
	var rows = make([]string, 0, 20)
	switch elm := node.Element.(type) {
	case string:
		rows = append(rows, elm)
	case fmt.Stringer:
		rows = append(rows, elm.String())
	case nil: // do nothing
	default: // do nothing
	}
	for _, nd := range node.Children {
		rows = append(rows, nd.Flatten())
	}
	return strings.Join(rows, " ")
}

func (node *Node) String() string {
	return node.Flatten()
}

func isNotWhiteString(sentence string) bool {
	return sentence != "" && !whiteSpaces.MatchString(sentence)
}

func normalizeString(sentence string) string {
	trimmed := strings.TrimSpace(sentence)
	return strings.Replace(trimmed, "\n", " ", -1)
}

func (node *Node) handleCharData(charData xml.CharData) {
	data := normalizeString(string(charData))
	if isNotWhiteString(data) {
		node.Children.Append(&Node{"chardata", data, nil})
	}
}

func (node *Node) handleNode(d *xml.Decoder, start xml.StartElement) (err error) {
	childNode := &Node{start.Name.Local, nil, Nodes{}}
	err = d.DecodeElement(childNode, &start)
	node.Children.Append(childNode)
	return err
}

func (node *Node) decodeTag(d *xml.Decoder, start xml.StartElement) (err error) {
	switch start.Name.Local {
	case "citerefentry":
		manEntry := &ManEntry{}
		err = d.DecodeElement(manEntry, &start)
		node.Children.Append(&Node{start.Name.Local, manEntry, nil})
	default:
		err = node.handleNode(d, start)
	}
	return err
}

func (node *Node) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var token xml.Token
	node.Type = start.Name.Local
	token, _ = d.Token()
	for token != nil {
		switch ttype := token.(type) {
		case xml.StartElement:
			err = node.decodeTag(d, ttype)
			if err != nil {
				break
			}
		case xml.CharData:
			node.handleCharData(ttype)
		}
		token, _ = d.Token()
	}
	return err
}
