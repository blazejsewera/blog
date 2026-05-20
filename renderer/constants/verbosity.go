package constants

type VerbosityLevel int

const (
	Silent VerbosityLevel = iota
	Error
	Warn
	Info
	Debug
)

func (l VerbosityLevel) String() string {
	switch l {
	case Silent:
		return "silent"
	case Error:
		return "error"
	case Warn:
		return "warn"
	case Info:
		return "info"
	case Debug:
		return "debug"
	default:
		return "unknown"
	}
}
