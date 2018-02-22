package extractor

import (
	"github.com/cmdse/core/argparse"
	"github.com/cmdse/core/schema"
)

func dismissOperandAndStackedCandidates(tokens schema.TokenList) {
	for _, token := range tokens.WhenContextFree() {
		token.ReduceCandidates(func(tokenType *schema.SemanticTokenType) bool {
			return !tokenType.Equal(schema.SemOperand) &&
				!tokenType.Equal(schema.SemPOSIXStackedShortSwitches) &&
				!tokenType.Equal(schema.SemHeadlessOption)
		})
	}
}

// OptSynopsisBehavior is a behavior specifically designed for option optionSynopsis parsing.
var OptSynopsisBehavior = &argparse.Behavior{
	RunInferences: func(p *argparse.Parser, token *schema.Token) {
		token.InferRight()
		token.InferLeft()
	},
	RunStaticChecks: func(p *argparse.Parser) {
		dismissOperandAndStackedCandidates(p.Tokens())
	},
}

// ParseOptSynopsis turn given arguments into a collection of tokens.
// All tokens are guaranteed to be of semantic type.
func ParseOptSynopsis(args []string) schema.TokenList {
	tokens := argparse.InitTokens(args)
	parser := argparse.NewParser(tokens, nil, OptSynopsisBehavior)
	return parser.ParseTokens()
}
