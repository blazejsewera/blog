package draft

import "github.com/blazejsewera/blog/internal/templates"

var templateNames = templates.Components("draft/draft-info", "draft/draft-badge")

func Draft(t *templates.Template) {
	t.ParseTFS(templateNames)
}

type Props struct {
	DraftDescription string
}

func (p Props) Description() string {
	if p.DraftDescription == "" {
		return "This is a draft article. It may be incomplete."
	}
	return p.DraftDescription
}
