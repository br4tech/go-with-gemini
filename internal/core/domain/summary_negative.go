package domain

type SummaryNegative struct {
	NegativeAspects []string `json:"negative_aspects"`
}

func NewSummaryNegative(negative_aspects []string) *SummaryNegative {
	return &SummaryNegative{
		NegativeAspects: negative_aspects,
	}
}
