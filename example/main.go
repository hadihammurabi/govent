package main

import (
	"fmt"

	"github.com/hadihammurabi/govent"
)

func main() {
	bus := govent.New()
	bus.On("click", func(args ...interface{}) {
		fmt.Println("clicked 1", args)
	})

	bus.On("click", func(args ...interface{}) {
		fmt.Println("clicked 2", args)
	})

	bus.Emit("click", "mantap", "asiap")
}
