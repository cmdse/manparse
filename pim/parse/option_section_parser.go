package parse

import (
	"regexp"

	"github.com/cmdse/manparse/docbook/section"
)

var regexIsSection = regexp.MustCompile(`^refsect[123]|refsection$`)

// -L, --dereference
// -p
// -P
// -h
// -ldflags 'flag list'
// -file-line-error-style
// -interaction mode
// -output-directory directory
// --parents
// -t, --target-directory=DIRECTORY
// --context[=CTX]
// --backup[=CONTROL]
// -a, --archive
// --warnings[=warnings]
// -C file, --config-file=file
// -L locale, --locale=locale
// -e sub-extension, --extension=sub-extension
// -m system[,...], --systems=system[,...]
var regexIsOptionSynopsis = regexp.MustCompile(`^(-|--).*$`)

func isSection(node *section.Node) bool {
	return regexIsSection.MatchString(node.Type)
}

func isVarList(node *section.Node) bool {
	return node.Type == "variablelist"
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

func varListToExtracts(varList *section.Node) rawOptExtracts {
	entries := varList.Children
	rawOptExtracts := make(rawOptExtracts, 0, 15)
	for _, entry := range entries {
		extract := &rawOptExtract{
			optSynopsis:    entry.Children[0],
			optDescription: entry.Children[1],
		}
		rawOptExtracts = append(rawOptExtracts, extract)
	}
	return rawOptExtracts
}

func aggregateRelevantNode(children section.Nodes) rawOptExtracts {
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
			Type:     "variablelist",
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
	return current.Type == "par" &&
		regexIsOptionSynopsis.MatchString(current.Flatten()) &&
		(next.Type == "literallayout" || next.Type == "programlisting" || next.Type == "blockquote")
}

func aggregateSiblingsToExtract(nodes section.Nodes) rawOptExtracts {
	rawOptExtracts := make(rawOptExtracts, 0, 15)
	var skip = new(int)
	*skip = 1
	for index := 0; index < len(nodes); index = index + *skip {
		node := nodes[index]
		sibling := nodes[index+1]
		if matchSiblingPattern(node, sibling) {
			extract := &rawOptExtract{
				optSynopsis:    node,
				optDescription: sibling,
			}
			rawOptExtracts = append(rawOptExtracts, extract)
			*skip = 2
			break
		}
	}
	*skip = 1
	return rawOptExtracts
}

var OptionSectionParser = SectionParser{
	TargetSection: "OPTION",
	aggregateExtracts: func(parser *SectionParser, section *section.Section) rawOptExtracts {
		// flatten sections, if any sub-sections
		flattenChildren := bubbleNodes(section.Children, isSection)
		return aggregateRelevantNode(flattenChildren)
	},
}
