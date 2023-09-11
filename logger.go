package httpclient

import "log"

type logger interface {
	Println(string)
}

type loggerFunc func(string)

func (fn loggerFunc) Println(msg string) {
	fn(msg)
}

var _defaultLogger loggerFunc = func(s string) {
	log.Println(s)
}

func DefaultLogger() logger {
	return _defaultLogger
}
