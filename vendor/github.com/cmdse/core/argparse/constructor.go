package argparse

import (
	"github.com/cmdse/core/schema"
)

// NewParser create a new parser from a token list, a program interface model and a behavior.
//
// See also
//
// * ArgParseBehavior for an example of implementing a behavior
// * InitTokens to convert arguments into tokens
// * ProgramInterfaceModel
func NewParser(tokens schema.TokenList, pim *schema.ProgramInterfaceModel, behaviour *Behavior) *Parser {
	return &Parser{
		behaviour,
		pim,
		tokens,
		1,
		1,
	}
}
