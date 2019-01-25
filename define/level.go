package define;

type Level uint8

const (
    ERROR Level = 0x01
    WARN  Level = 0x02
    INFO  Level = 0x03
    DEBUG Level = 0x04
)

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "debug"
	case INFO:
		return "info"
	case WARN:
		return "warn"
	case ERROR:
		return "eror"
	default:
		panic("bad level")
	}
}
