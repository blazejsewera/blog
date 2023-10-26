package templates

func require(s string) string {
	if s == "" {
		panic("require: value was empty")
	}
	return s
}
