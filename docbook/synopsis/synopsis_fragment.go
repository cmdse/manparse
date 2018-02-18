package synopsis

import "fmt"

// Arg | Group
// http://tdg.docbook.org/tdg/4.5/synopfragment.html
type SynopsisFragment interface {
	fmt.Stringer
}
