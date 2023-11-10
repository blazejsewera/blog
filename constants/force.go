package constants

type ForceLevel int

const (
	NoForce ForceLevel = iota
	ReRender
	ReDownload
)
