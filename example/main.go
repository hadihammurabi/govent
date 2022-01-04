package main

import (
	"fmt"

	"github.com/hadihammurabi/govent"
)

func main() {
	bus := govent.New()
	bus.On("click", func(e govent.Event) {
		fmt.Println("clicked 1", e.Args)
	})

	bus.On("click", func(e govent.Event) {
		fmt.Println("clicked 1", e.Args)
	})

	bus.Emit("click", "mantap", "asiap")
}
