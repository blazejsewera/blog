package molecule

type DraftProps struct {
	DraftDescription string
}

func (p DraftProps) Description() string {
	if p.DraftDescription == "" {
		return "This is a draft article. It may be incomplete."
	}
	return p.DraftDescription
}
