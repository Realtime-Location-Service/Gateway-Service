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
