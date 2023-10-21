package whitespace

import (
	"bytes"
)

func Collapse(bb []byte) []byte {
	buf := &bytes.Buffer{}
	type Char byte
	const (
		TabOrSpace Char = iota
		Newline
		Other
	)
	state := Other

	for _, b := range bb {
		switch b {
		case '\r', '\n':
			switch state {
			case Newline:
				continue
			default:
				state = Newline
				buf.WriteByte('\n')
			}
		case '\t', ' ':
			switch state {
			case TabOrSpace:
				continue
			default:
				state = TabOrSpace
				buf.WriteByte(' ')
			}
		default:
			state = Other
			buf.WriteByte(b)
		}
	}
	return buf.Bytes()
}
