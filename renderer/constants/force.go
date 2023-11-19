package constants

type ForceLevel int

const (
	NoForce ForceLevel = iota
	ReRender
	RemoveAndReRender
	ReDownload
)
