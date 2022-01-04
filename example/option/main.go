package main

import (
	"fmt"

	"github.com/hadihammurabi/govent"
	"github.com/hadihammurabi/govent/option"
)

func main() {
	bus := govent.New(option.SingleListener)
	bus.On("click", func(e govent.Event) {
		fmt.Println("clicked 1", e.Args)
	})

	bus.On("click", func(e govent.Event) {
		fmt.Println("clicked 2", e.Args)
	})

	bus.Emit("click", "mantap", "asiap")
}
