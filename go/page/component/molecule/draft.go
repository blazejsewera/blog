package molecule

import "github.com/blazejsewera/blog/internal/templates"

func Draft(t *templates.Template) {
	templateNames := templates.Components("molecule/draft-info", "molecule/draft-badge")
	t.ParseTFS(templateNames)
}

type DraftProps struct {
	DraftDescription string
}

func (p DraftProps) Description() string {
	if p.DraftDescription == "" {
		return "This is a draft article. It may be incomplete."
	}
	return p.DraftDescription
}
