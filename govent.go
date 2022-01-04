package govent

type Bus struct {
	Events map[string]*Event
}

func New() *Bus {
	return &Bus{
		Events: make(map[string]*Event),
	}
}

func (b *Bus) On(name string, handler EventHandler) {
	if _, ok := b.Events[name]; !ok {
		b.Events[name] = NewEvent(handler)
		return
	}

	b.Events[name].AddHandler(handler)
}

func (b Bus) Emit(name string, args ...interface{}) {
	event, ok := b.Events[name]
	if !ok {
		panic("no handler for event " + name)
	}

	event.ExecHandlers(args...)
}
