package domain

type Summary struct {
	Positive []string
	Negative []string
	Opinions []Opinion
}

func NewSummary(positive []string, negative []string, opinions []Opinion) *Summary {
	return &Summary{
		Positive: positive,
		Negative: negative,
		Opinions: opinions,
	}
}
