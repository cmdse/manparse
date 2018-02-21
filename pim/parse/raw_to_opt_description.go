package parse

import (
	"fmt"
	"regexp"

	"github.com/cmdse/core/schema"
	"github.com/cmdse/manparse/docbook/section"
)

var whitespacesRegex = regexp.MustCompile(`\s+`)

func tokensToMatchModel(optParts schema.TokenList) (*schema.MatchModel, error) {
	var semanticTypes = make([]*schema.SemanticTokenType, len(optParts))
	for i, token := range optParts {
		switch ttype := token.Ttype.(type) {
		case *schema.SemanticTokenType:
			semanticTypes[i] = ttype
		case *schema.ContextFreeTokenType:
			return nil, fmt.Errorf("failure to extract MatchModel: '%v' token could not be converted to semantic type ; found instead '%v' with candidate '%v' ", token.Value, token.Ttype.Name(), token.SemanticCandidates)
		}
	}
	switch len(optParts) {
	case 1, 2:
		optExpression, err := semanticTypes[0].Variant().Assemble(optParts)
		if err != nil {
			return nil, fmt.Errorf("failure to extract MatchModel: %v", err.Error())
		}
		definition := optExpression.Options()[0]
		return schema.NewMatchModelFromDefinition(definition), nil
	default:
		return nil, fmt.Errorf("failure to extract MatchModel: synopsis has %v option parts instead of 1 or 2 expected", len(optParts))
	}
}

func matchModelsFromSynopsisString(synopsis string) (schema.MatchModels, error) {
	models := make(schema.MatchModels, 0, 2)
	expressions := normalizeOptSynopsises(splitSynopsis(synopsis))
	for _, expr := range expressions {
		args := whitespacesRegex.Split(expr, -1)
		tokens := ParseOptSynopsis(args)
		matchModel, err := tokensToMatchModel(tokens)
		if err != nil {
			return nil, err
		} else {
			models = append(models, matchModel)
		}
	}
	return models, nil
}

func matchModelsFromSection(synopsis *section.Node) (schema.MatchModels, error) {
	return matchModelsFromSynopsisString(synopsis.Flatten())
}

func extractToOptDescription(extract *rawOptExtract) (*schema.OptDescription, error) {
	matchModels, err := matchModelsFromSection(extract.optSynopsis)
	if err != nil {
		return nil, err
	}
	return &schema.OptDescription{
		Description: extract.optDescription.Flatten(),
		MatchModels: matchModels,
	}, nil
}
