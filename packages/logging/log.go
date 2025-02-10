package logging

var (
	LogSerive logger
)

type ILogger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

type logger struct {
	log ILogger
}

func (l logger) debug(msg string) {
	l.log.Debug(msg)
}

func (l logger) info(msg string) {
	l.log.Info(msg)
}

func (l logger) warn(msg string) {
	l.log.Warn(msg)
}
func (l logger) error(err string) {
	l.log.Error(err)
}

func Debug(msg string, log logger) {
	log.debug(msg)
}

func Info(msg string, log logger) {
	log.info(msg)
}

func Warn(msg string, log logger) {
	log.warn(msg)
}

func Error(err string, log logger) {
	log.error(err)
}
