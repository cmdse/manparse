package guesses

type Guess struct {
	Name        string
	Description string
}

var (
	RemoteOptSynopsis = &Guess{
		"REMOTE_OPT_SYNOPSIS",
		"An option synopsis has more than two coma separated values and a reference to a sub-command appears in the description.",
	}
	SuggestedPosixImplicitAssignment = &Guess{
		"SUGGESTED_POSIX_IMPLICIT_ASSIGNMENT",
		"An option synopsis has two option descriptions with the latest having a value assignment and the foremost being a short POSIX switch.",
	}
	OptionalExplicitAssignment = &Guess{
		"OPTIONAL_EXPLICIT_ASSIGNMENT",
		"An option synopsis holds an optional explicit assignment.",
	}
	OptionalImplicitAssignment = &Guess{
		"OPTIONAL_IMPLICIT_ASSIGNMENT",
		"An option synopsis holds an optional implicit assignment.",
	}
)

func (guess *Guess) String() string {
	return guess.Name
}
