package domain

type SummaryPositive struct {
	Positive []string
}

func NewSummaryPositive(positive []string) *SummaryPositive {
	return &SummaryPositive{
		Positive: positive,
	}
}
