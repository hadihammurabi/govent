package govent

type EventHandler func(args ...interface{})

type Event struct {
	Handlers []EventHandler
}

func NewEvent(handers ...EventHandler) *Event {
	return &Event{
		Handlers: handers,
	}
}

func (e *Event) AddHandler(handler EventHandler) {
	e.Handlers = append(e.Handlers, handler)
}

func (e Event) ExecHandlers(args ...interface{}) {
	for _, handler := range e.Handlers {
		handler(args...)
	}
}
