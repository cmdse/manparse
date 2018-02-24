package extractor

type DryOptExtract struct {
	optSynopsis    string
	optDescription string
}

type DryOptExtracts []*DryOptExtract

func dryUpOptExtract(rawOptExtracts RawOptExtracts) DryOptExtracts {
	dried := make(DryOptExtracts, len(rawOptExtracts))
	for i, raw := range rawOptExtracts {
		dried[i] = &DryOptExtract{
			optSynopsis:    raw.optSynopsis.Flatten(),
			optDescription: raw.optDescription.Flatten(),
		}
	}
	return dried
}
