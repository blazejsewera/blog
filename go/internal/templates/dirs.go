package templates

const relativeComponentDir = "component/"

func Pages(pageNames ...string) []string {
	return transform(pageNames, func(name string) string {
		return templateName(name)
	})
}

func Components(componentNames ...string) []string {
	return transform(componentNames, func(name string) string {
		return relativeComponentDir + templateName(name)
	})
}

func templateName(name string) string {
	return name + ".html.tmpl"
}

func transform(in []string, f func(string) string) []string {
	out := make([]string, len(in))
	for i, s := range in {
		out[i] = f(s)
	}
	return out
}
