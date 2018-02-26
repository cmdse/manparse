package extractor

import (
	"regexp"

	"github.com/cmdse/core/schema"
	"github.com/cmdse/manparse/pim/optdesc/normalize"
	"github.com/cmdse/manparse/reporter"
	"github.com/cmdse/manparse/reporter/guesses"
)

type optionSynopsis struct {
	*reporter.ParseReporter
	raw         string
	expressions []string
	description string
}

var whitespacesRegex = regexp.MustCompile(`\s+`)

func newOptionSynopsis(reporter *reporter.ParseReporter, raw string, expressions []string, description string) *optionSynopsis {
	return &optionSynopsis{
		reporter,
		raw,
		expressions,
		description,
	}
}

func (synopsis *optionSynopsis) handleOptionAssignment(synopses *optionSynopses, guess *guesses.Guess, finder func(expression string) ([]string, bool)) bool {
	expr := synopsis.expressions[0]
	split, ok := finder(expr)
	if ok {
		synopsis.ReportGuessf(
			guess,
			"found an option expression witch matches the optional option assignment pattern, so I split it to two synopsis, '%v' and '%v'",
			split[0], split[1])
		flagSynopsis := newOptionSynopsis(
			synopsis.ParseReporter,
			split[0],
			[]string{split[0]},
			synopsis.description,
		)
		assignmentSynopsis := newOptionSynopsis(
			synopsis.ParseReporter,
			split[1],
			[]string{split[1]},
			synopsis.description,
		)
		synopses.append(flagSynopsis, assignmentSynopsis)
	}
	return ok
}

func (synopsis *optionSynopsis) splitSynopsisIfOptionalAssignment() optionSynopses {
	var newSynopses = make(optionSynopses, 0, 2)
	context := synopsis.SetContextf(synopsis.raw)
	defer synopsis.ReleaseContext(context)
	if len(synopsis.expressions) == 1 {
		foundExplicitAssignment := synopsis.handleOptionAssignment(&newSynopses, guesses.OptionalExplicitAssignment, findOptionalExplicitAssignment)
		if foundExplicitAssignment {
			return newSynopses
		}
		foundImplicitAssignment := synopsis.handleOptionAssignment(&newSynopses, guesses.OptionalImplicitAssignment, findOptionalImplicitAssignment)
		if foundImplicitAssignment {
			return newSynopses
		}
	}
	return append(newSynopses, synopsis)
}

func (synopsis *optionSynopsis) toOptDescription() *schema.OptDescription {
	context := synopsis.SetContextf(synopsis.raw)
	optDescription := synopsis.extractOptDescription()
	if optDescription != nil {
		variantNames := formatVariantNames(optDescription.Variants())
		synopsis.ReportSuccessf("found '%v'", variantNames)
	}
	synopsis.ReleaseContext(context)
	return optDescription
}

func (synopsis *optionSynopsis) extractOptDescription() *schema.OptDescription {
	matchModels := synopsis.extractMatchModels()
	if len(matchModels) > 0 {
		return &schema.OptDescription{
			Description: synopsis.description,
			MatchModels: matchModels,
		}
	}
	return nil
}

func (synopsis *optionSynopsis) extractMatchModels() schema.MatchModels {
	models := make(schema.MatchModels, 0, 2)
	expressions := normalize.NormalizeOptDescriptions(synopsis.expressions)
	argsList := make([][]string, len(expressions))
	for i, expr := range expressions {
		model, args := synopsis.extractModelFromOptionExpression(expr)
		argsList[i] = args
		if model != nil {
			models = append(models, model)
		} else {
			return nil
		}
	}
	return synopsis.normalizeModelsAssignments(models, argsList)
}

func (synopsis *optionSynopsis) extractModelFromOptionExpression(expr string) (*schema.MatchModel, []string) {
	args := whitespacesRegex.Split(expr, -1)
	tokens := ParseOptSynopsis(args)
	matchModel := synopsis.optionPartsToToMatchModel(tokens)
	return matchModel, args
}

func (synopsis *optionSynopsis) optionPartsToToMatchModel(optParts schema.TokenList) *schema.MatchModel {
	var semanticTypes = make([]*schema.SemanticTokenType, len(optParts))
	for i, token := range optParts {
		switch ttype := token.Ttype.(type) {
		case *schema.SemanticTokenType:
			semanticTypes[i] = ttype
		case *schema.ContextFreeTokenType:
			synopsis.ReportSkipf("I could not recognize the option expression\n'%v' token is ambiguous since it can match the following candidates: '%v' %", token.Value, token.Ttype.Name(), token.SemanticCandidates)
			return nil
		}
	}
	switch len(optParts) {
	case 1, 2:
		optExpression, err := semanticTypes[0].Variant().Assemble(optParts)
		if err != nil {
			synopsis.ReportFailuref("failure to extract MatchModel: %v", err.Error())
			return nil
		}
		definition := optExpression.Options()[0]
		return schema.NewMatchModelFromDefinition(definition)
	default:
		synopsis.ReportFailuref("failure to extract MatchModel: optionSynopsis has %v option parts instead of 1 or 2 expected", len(optParts))
		return nil
	}
}

func (synopsis *optionSynopsis) normalizeModelsAssignments(matchModels schema.MatchModels, argsList [][]string) schema.MatchModels {
	assignmentParamName, atLeastOneButNotAllExpressionsAreAssignments := synopsis.findAssignmentInExpressions(matchModels)
	if atLeastOneButNotAllExpressionsAreAssignments {
		synopsis.convertPOSIXFlagsToAssignments(matchModels, argsList, assignmentParamName)
	}
	return matchModels
}

func (synopsis *optionSynopsis) findAssignmentInExpressions(matchModels schema.MatchModels) (assignmentParamName string, atLeastOneButNotAllExpressionsAreAssignments bool) {
	foundAssignmentExpression := false
	allAreAssignmentExpressions := true
	assignmentParamName = ""
	for _, model := range matchModels {
		if model.ParamName() != "" {
			foundAssignmentExpression = true
			assignmentParamName = model.ParamName()
		} else {
			allAreAssignmentExpressions = false
		}
	}
	return assignmentParamName, foundAssignmentExpression && !allAreAssignmentExpressions
}

func (synopsis *optionSynopsis) convertPOSIXFlagsToAssignments(matchModels schema.MatchModels, argsList [][]string, assignmentParamName string) {
	for i, model := range matchModels {
		// this MatchModel is not assignment
		context := synopsis.SetContextf("expression '%v' with expression model '%v'", argsList[i], model.Variant().Name())
		if model.ParamName() == "" {
			args := append(argsList[i], assignmentParamName)
			tokens := ParseOptSynopsis(args)
			matchModel := synopsis.optionPartsToToMatchModel(tokens)
			if matchModel != nil {
				matchModels[i] = matchModel
				synopsis.ReportGuessf(
					guesses.SuggestedPosixImplicitAssignment,
					"the latest option expression has an implicit assignment of param '%v' while foremost has no option assignment, so I guessed it should have an implicit option assignment",
					assignmentParamName)
				synopsis.ReleaseContext(context)
				break
			} else {
				synopsis.ReportFailuref("didn't expect match model extraction to fail for optionSynopsis : %v", args)
			}
		}
		synopsis.ReleaseContext(context)
	}
}
