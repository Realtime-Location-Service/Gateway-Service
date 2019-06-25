package consts

// LoggerType ...
type LoggerType string

const (
	// KitLogger ...
	KitLogger LoggerType = "kitlogger"
	// Logrus ...
	Logrus LoggerType = "logrus"
)

// SupportedLogger ...
var SupportedLogger = map[LoggerType]bool{
	KitLogger: true,
}

// content types
const (
	ContentTypeJSON = "application/json; utf8"
	ContentType     = "Content-Type"
)

// Request headers
const (
	RLSReferrer    = "RLS-Referrer"
	AppKey         = "appKey"
	UserID         = "userId"
	SubordinateIDs = "subordinateIds"
)
