package constants

type VerbosityLevel int

const (
	Silent VerbosityLevel = iota
	Error
	Warn
	Info
	Debug
)
