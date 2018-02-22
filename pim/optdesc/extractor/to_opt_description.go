package extractor

import (
	"regexp"

	"github.com/cmdse/core/schema"
	"github.com/cmdse/manparse/pim/optdesc/normalize"
)

var whitespacesRegex = regexp.MustCompile(`\s+`)

func (extractor *Extractor) tokensToMatchModel(optParts schema.TokenList) *schema.MatchModel {
	extractor.SetContext("tokensToMatchModel") // TODO handle context better
	var semanticTypes = make([]*schema.SemanticTokenType, len(optParts))
	for i, token := range optParts {
		switch ttype := token.Ttype.(type) {
		case *schema.SemanticTokenType:
			semanticTypes[i] = ttype
		case *schema.ContextFreeTokenType:
			extractor.Reportf("failure to extract MatchModel: '%v' token could not be converted to semantic type ; found instead '%v' with candidate '%v' ", token.Value, token.Ttype.Name(), token.SemanticCandidates)
			return nil
		}
	}
	switch len(optParts) {
	case 1, 2:
		optExpression, err := semanticTypes[0].Variant().Assemble(optParts)
		if err != nil {
			extractor.Errorf("failure to extract MatchModel: %v", err.Error())
			return nil
		}
		definition := optExpression.Options()[0]
		return schema.NewMatchModelFromDefinition(definition)
	default:
		extractor.Errorf("failure to extract MatchModel: optionSynopsis has %v option parts instead of 1 or 2 expected", len(optParts))
		return nil
	}
}

func (extractor *Extractor) checkModelsAreAssignments(synopsis schema.MatchModels, argsList [][]string) schema.MatchModels {
	isAssignmentSynopsis := false
	assignmentParamName := ""
	for _, model := range synopsis {
		if model.ParamName() != "" {
			isAssignmentSynopsis = true
			assignmentParamName = model.ParamName()
		}
	}
	for i, model := range synopsis {
		// this MatchModel is not assignment
		if isAssignmentSynopsis && model.ParamName() == "" {
			args := append(argsList[i], assignmentParamName)
			tokens := ParseOptSynopsis(args)
			matchModel := extractor.tokensToMatchModel(tokens)
			if matchModel != nil {
				synopsis[i] = matchModel
				break
			} else {
				extractor.Errorf("didn't expect match model extraction to fail for optionSynopsis : %v", args)
			}
		}
	}
	return synopsis
}

func (extractor *Extractor) matchModelsFromSynopsis(synopsis *optionSynopsis) schema.MatchModels {
	models := make(schema.MatchModels, 0, 2)
	expressions := normalize.NormalizeOptDescriptions(synopsis.expressions)
	argsList := make([][]string, len(expressions))
	for i, expr := range expressions {
		args := whitespacesRegex.Split(expr, -1)
		argsList[i] = args
		tokens := ParseOptSynopsis(args)
		matchModel := extractor.tokensToMatchModel(tokens)
		if matchModel == nil {
			return nil
		} else {
			models = append(models, matchModel)
		}
	}
	return extractor.checkModelsAreAssignments(models, argsList)
}
