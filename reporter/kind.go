package reporter

import "github.com/fatih/color"

type Kind struct {
	name  string
	char  string
	color *color.Color
}

func (kind *Kind) Sprint(message string) string {
	return kind.color.Sprintf("%v %v", kind.char, message)
}

var (
	KindSuccess = &Kind{
		"SUCCESS",
		"✔",
		color.New(color.FgGreen),
	}
	KindSkip = &Kind{
		"SKIP",
		"?",
		color.New(color.FgYellow),
	}
	KindFailure = &Kind{
		"FAILURE",
		"✘",
		color.New(color.FgRed),
	}
	KindGuess = &Kind{
		"GUESS",
		"💡",
		color.New(color.FgBlue),
	}
)
