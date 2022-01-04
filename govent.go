package govent

import "github.com/hadihammurabi/govent/option"

type Bus struct {
	Events map[string]*Event

	singleListener bool
}

func New(options ...*option.Option) *Bus {
	bus := &Bus{
		Events: make(map[string]*Event),
	}
	bus.parseOptions(options...)
	return bus
}

func (b *Bus) parseOptions(options ...*option.Option) {
	for _, opt := range options {
		switch opt {
		case option.SingleListener:
			b.singleListener = true
		}
	}
}

func (b *Bus) On(name string, handler EventHandlerCallback) {
	if _, ok := b.Events[name]; !ok {
		b.Events[name] = NewEvent(name, handler)
		return
	}

	b.Events[name].AddHandler(handler, b.singleListener)
}

func (b Bus) Emit(name string, args ...interface{}) {
	event, ok := b.Events[name]
	if !ok {
		panic("no handler for event " + name)
	}

	event.ExecHandlers(args...)
}
