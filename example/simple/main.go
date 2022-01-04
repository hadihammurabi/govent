package main

import (
	"fmt"

	"github.com/hadihammurabi/govent"
)

func main() {
	bus := govent.New()
	bus.On("click", func(e govent.Event) error {
		fmt.Println("clicked 1", e.Args)
		return nil
	})

	bus.On("click", func(e govent.Event) error {
		fmt.Println("clicked 2", e.Args)
		return nil
	})

	bus.Emit("click", "mantap", "asiap")
}
