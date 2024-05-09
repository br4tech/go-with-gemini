package domain

type SummaryPositive struct {
	PositiveAspects []string `json:"positive_aspects"`
}

func NewSummaryPositive(positive_aspects []string) *SummaryPositive {
	return &SummaryPositive{
		PositiveAspects: positive_aspects,
	}
}
