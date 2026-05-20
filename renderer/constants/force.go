package constants

type ForceLevel int

const (
	NoForce ForceLevel = iota
	ReRender
	RemoveAndReRender
	ReDownload
)

func (l ForceLevel) String() string {
	switch l {
	case NoForce:
		return "no force"
	case ReRender:
		return "re-render"
	case RemoveAndReRender:
		return "remove and re-render"
	case ReDownload:
		return "re-download"
	default:
		return "unknown"
	}
}
