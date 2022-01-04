package govent

type EventHandlerCallback func(Event)
type EventHandler struct {
	Callback EventHandlerCallback
}

type Event struct {
	Name     string
	Handlers []EventHandler
	Args     []interface{}
}

func NewEvent(name string, handlers ...EventHandlerCallback) *Event {
	eventHandlers := []EventHandler{}
	for _, callback := range handlers {
		eventHandlers = append(eventHandlers, EventHandler{Callback: callback})
	}

	return &Event{
		Handlers: eventHandlers,
	}
}

func (e *Event) AddHandler(handler EventHandlerCallback) {
	e.Handlers = append(e.Handlers, EventHandler{Callback: handler})
}

func (e Event) ExecHandlers(args ...interface{}) {
	for _, handler := range e.Handlers {
		e.Args = args
		handler.Callback(e)
	}
}
