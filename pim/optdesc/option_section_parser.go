package optdesc

import (
	"regexp"

	"github.com/cmdse/manparse/docbook/section"
	"github.com/cmdse/manparse/pim/optdesc/extractor"
)

var regexIsSection = regexp.MustCompile(`^refsect[123]|refsection$`)

var regexIsOptionSynopsis = regexp.MustCompile(`^(-|--).*$`)

const (
	DocbookNodeVariableList   = "variablelist"
	DocbookNodeParagraph      = "para"
	DocbookNodeLiteralLayout  = "literallayout"
	DocbookNodeProgramListing = "programlisting"
	DocbookNodeBlockQuote     = "blockquote"
)

func isSection(node *section.Node) bool {
	return regexIsSection.MatchString(node.Type)
}

func isVarList(node *section.Node) bool {
	return node.Type == DocbookNodeVariableList
}

// This function returns a slice of nodes which are all comprised of "non-predicate" children.
// When a "predicated" child is encountered, it is replaced with its non-predicate children recursively
func bubbleNodes(sections section.Nodes, predicate func(*section.Node) bool) []*section.Node {
	nodes := make(section.Nodes, 0, 10)
	for _, child := range sections {
		if predicate(child) {
			nodes = append(nodes, bubbleNodes(child.Children, predicate)...)
		} else {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func filterVarLists(children section.Nodes) section.Nodes {
	varLists := make(section.Nodes, 0, 3)
	for _, child := range children {
		if isVarList(child) {
			varLists = append(varLists, child)
		}
	}
	return varLists
}

func varListToExtracts(varList *section.Node) extractor.RawOptExtracts {
	entries := varList.Children
	rawOptExtracts := make(extractor.RawOptExtracts, 0, 15)
	for _, entry := range entries {
		extract := extractor.NewRawOptExtract(entry.Children[0], entry.Children[1])
		rawOptExtracts = append(rawOptExtracts, extract)
	}
	return rawOptExtracts
}

func aggregateRelevantNode(children section.Nodes) extractor.RawOptExtracts {
	varLists := filterVarLists(children)
	switch len(varLists) {
	case 0:
		return aggregateSiblingsToExtract(children)
	case 1:
		return varListToExtracts(varLists[0])
	default:
		// create a unique variable list by concatenating all variable lists
		nodes := bubbleNodes(varLists, isVarList)
		varList := &section.Node{
			Type:     DocbookNodeVariableList,
			Element:  nil,
			Children: nodes,
		}
		return varListToExtracts(varList)
	}
}

// * SIBLINGS :
//   (*any node) >
//   	para                                        : optSynopsis
//      literallayout | programlisting | blockquote : optDescription
//
func matchSiblingPattern(current *section.Node, next *section.Node) bool {
	return current.Type == DocbookNodeParagraph &&
		regexIsOptionSynopsis.MatchString(current.Flatten()) &&
		(next.Type == DocbookNodeLiteralLayout || next.Type == DocbookNodeProgramListing || next.Type == DocbookNodeBlockQuote)
}

func aggregateSiblingsToExtract(nodes section.Nodes) extractor.RawOptExtracts {
	rawOptExtracts := make(extractor.RawOptExtracts, 0, 15)
	var skip = new(int)
	*skip = 1
	index := 0
	for index < len(nodes)-1 {
		node := nodes[index]
		sibling := nodes[index+1]
		if matchSiblingPattern(node, sibling) {
			extract := extractor.NewRawOptExtract(node, sibling)
			rawOptExtracts = append(rawOptExtracts, extract)
			*skip = 2
		}
		index = index + *skip
	}
	*skip = 1
	return rawOptExtracts
}

var OptionSectionParser = SectionParser{
	TargetSection: optionSectionName,
	aggregateExtracts: func(parser *SectionParser, section *section.Section) extractor.RawOptExtracts {
		// flatten sections, if any sub-sections
		flattenChildren := bubbleNodes(section.Children, isSection)
		return aggregateRelevantNode(flattenChildren)
	},
}
