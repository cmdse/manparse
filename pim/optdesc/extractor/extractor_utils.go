package extractor

import (
	"regexp"

	"github.com/cmdse/core/schema"
	"github.com/cmdse/manparse/pim/optdesc/normalize"
	"github.com/cmdse/manparse/reporter/guesses"
)

var whitespacesRegex = regexp.MustCompile(`\s+`)

func (extractor *Extractor) extractToOptDescription(synopsis *optionSynopsis) *schema.OptDescription {
	matchModels := extractor.matchModelsFromSynopsis(synopsis)
	if len(matchModels) > 0 {
		return &schema.OptDescription{
			Description: synopsis.description,
			MatchModels: matchModels,
		}
	} else {
		return nil
	}
}

func (extractor *Extractor) matchModelsFromSynopsis(synopsis *optionSynopsis) schema.MatchModels {
	models := make(schema.MatchModels, 0, 2)
	expressions := normalize.NormalizeOptDescriptions(synopsis.expressions)
	argsList := make([][]string, len(expressions))
	for i, expr := range expressions {
		extractor.SetContextf("with option expression '%v'", synopsis.expressions)
		args := whitespacesRegex.Split(expr, -1)
		argsList[i] = args
		tokens := ParseOptSynopsis(args)
		matchModel := extractor.optionPartsToToMatchModel(tokens)
		if matchModel == nil {
			return nil
		} else {
			models = append(models, matchModel)
		}
		extractor.RedeemContext()
	}
	return extractor.normalizeModelsAssignments(models, argsList)
}

func (extractor *Extractor) optionPartsToToMatchModel(optParts schema.TokenList) *schema.MatchModel {
	var semanticTypes = make([]*schema.SemanticTokenType, len(optParts))
	for i, token := range optParts {
		switch ttype := token.Ttype.(type) {
		case *schema.SemanticTokenType:
			semanticTypes[i] = ttype
		case *schema.ContextFreeTokenType:
			extractor.ReportFailuref("failure to extract MatchModel: '%v' token could not be converted to semantic type ; found instead '%v' with candidate '%v' ", token.Value, token.Ttype.Name(), token.SemanticCandidates)
			return nil
		}
	}
	switch len(optParts) {
	case 1, 2:
		optExpression, err := semanticTypes[0].Variant().Assemble(optParts)
		if err != nil {
			extractor.ReportFailuref("failure to extract MatchModel: %v", err.Error())
			return nil
		}
		definition := optExpression.Options()[0]
		return schema.NewMatchModelFromDefinition(definition)
	default:
		extractor.ReportFailuref("failure to extract MatchModel: optionSynopsis has %v option parts instead of 1 or 2 expected", len(optParts))
		return nil
	}
}

func (extractor *Extractor) normalizeModelsAssignments(synopsis schema.MatchModels, argsList [][]string) schema.MatchModels {
	assignmentParamName, atLeastOneButNotAllExpressionsAreAssignments := extractor.findAssignmentInExpressions(synopsis)
	if atLeastOneButNotAllExpressionsAreAssignments {
		extractor.convertPOSIXFlagsToAssignments(synopsis, argsList, assignmentParamName)
	}
	return synopsis
}

func (extractor *Extractor) findAssignmentInExpressions(synopsis schema.MatchModels) (assignmentParamName string, atLeastOneButNotAllExpressionsAreAssignments bool) {
	foundAssignmentExpression := false
	allAreAssignmentExpressions := true
	assignmentParamName = ""
	for _, model := range synopsis {
		if model.ParamName() != "" {
			foundAssignmentExpression = true
			assignmentParamName = model.ParamName()
		} else {
			allAreAssignmentExpressions = false
		}
	}
	return assignmentParamName, foundAssignmentExpression && !allAreAssignmentExpressions
}

func (extractor *Extractor) convertPOSIXFlagsToAssignments(synopsis schema.MatchModels, argsList [][]string, assignmentParamName string) {
	for i, model := range synopsis {
		// this MatchModel is not assignment
		extractor.SetContextf("expression '%v' with expression model '%v'", argsList[i], model.Variant().Name())
		if model.ParamName() == "" {
			args := append(argsList[i], assignmentParamName)
			tokens := ParseOptSynopsis(args)
			matchModel := extractor.optionPartsToToMatchModel(tokens)
			if matchModel != nil {
				synopsis[i] = matchModel
				extractor.ReportGuessf(
					guesses.SuggestedPosixImplicitAssignment,
					"a synopsis had the latest option expression with implicit assignment of param '%v' and foremost has no option assignment, so I guessed it should have an implicit option assignment",
					assignmentParamName)
				break
			} else {
				extractor.ReportFailuref("didn't expect match model extraction to fail for optionSynopsis : %v", args)
			}
		}
		extractor.RedeemContext()
	}
}
