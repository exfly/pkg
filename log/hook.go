package log

const LoggerNameField = "logger"

type LoggerNameHook struct {
	tag string
}

func (*LoggerNameHook) Levels() []Level {
	return AllLevels
}

func (h *LoggerNameHook) Fire(entry *Entry) error {
	entry.Data[LoggerNameField] = h.tag
	return nil
}

func NewLoggerNameHook(tag string) *LoggerNameHook {
	return &LoggerNameHook{tag: tag}
}
