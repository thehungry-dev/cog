package log

type MultiLogger []Logger

func (loggers MultiLogger) Put(msg Message) interface{} {
	var result interface{}

	for _, logger := range loggers {
		result = logger.Put(msg)
	}

	return result
}
