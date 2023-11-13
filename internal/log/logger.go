package log

import (
	"github.com/blazejsewera/blog/constants"
	ll "log"
)

const (
	errorPrefix = "error: "
	warnPrefix  = "warn: "
	infoPrefix  = "info: "
	debugPrefix = "debug: "
)

var verbosity = constants.Silent

func SetVerbosity(level constants.VerbosityLevel) {
	verbosity = level
}

func Error(format string, v ...any) {
	if verbosity >= constants.Error {
		ll.Printf(errorPrefix+format, v...)
	}
}

func Warn(format string, v ...any) {
	if verbosity >= constants.Warn {
		ll.Printf(warnPrefix+format, v...)
	}
}

func Info(format string, v ...any) {
	if verbosity >= constants.Info {
		ll.Printf(infoPrefix+format, v...)
	}
}

func Debug(format string, v ...any) {
	if verbosity >= constants.Debug {
		ll.Printf(debugPrefix+format, v...)
	}
}
